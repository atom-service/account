package provider

import (
	"context"

	"github.com/joho/godotenv"

	"github.com/yinxulai/grpc-services/account/models"
	"github.com/yinxulai/grpc-services/account/standard"
)

// NewService NewService
func NewService() *Service {
	godotenv.Load()
	service := new(Service)
	return service
}

// Service Service
type Service struct {
}

// Create 创建用户 TODO: 检查 Inviter 是否存在
func (srv *Service) Create(ctx context.Context, req *standard.CreateRequest) (resp *standard.CreateResponse, err error) {
	var count uint64
	var user models.User
	user.LoadProtoStruct(req.User)
	user.SetPassword(req.User.GetPassword())
	resp = new(standard.CreateResponse)

	if !usernamePattern.MatchString(req.User.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	if !passwordPattern.MatchString(req.User.Password) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	// 查询 用户名是否已经存在
	err = countUserByUsernameStmt.GetContext(ctx, &count, user)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count > 0 {
		resp.State = standard.State_USER_ALREADY_EXISTS
		resp.Message = "该用户已存在"
		return resp, nil
	}

	// 执行插入
	_, err = insertUserStmt.ExecContext(ctx, user)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"

	return resp, nil
}

// QueryByID 通过ID查询用户
func (srv *Service) QueryByID(ctx context.Context, req *standard.QueryByIDRequest) (resp *standard.QueryByIDResponse, err error) {
	users := []*models.User{}
	resp = new(standard.QueryByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	rows, err := queryUserByIDStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localUser models.User
		err = rows.StructScan(&localUser)
		if err == nil {
			localUser.Password = "加密字段"
			users = append(users, &localUser)
		}
	}

	if len(users) <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = users[0].OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

// QueryByUsername 通过ID查询用户
func (srv *Service) QueryByUsername(ctx context.Context, req *standard.QueryByUsernameRequest) (resp *standard.QueryByUsernameResponse, err error) {
	users := []*models.User{}
	resp = new(standard.QueryByUsernameResponse)

	if !usernamePattern.MatchString(req.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	rows, err := queryUserByUsernameStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localUser models.User
		err = rows.StructScan(&localUser)
		if err == nil {
			localUser.Password = "加密字段"
			users = append(users, &localUser)
		}
	}

	if len(users) <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = users[0].OutProtoStruct()
	resp.Message = "查询成功"

	return resp, nil
}

// UpdateByID 通过ID更新用户
func (srv *Service) UpdateByID(ctx context.Context, req *standard.UpdateByIDRequest) (resp *standard.UpdateByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	user := new(models.User)
	user.LoadProtoStruct(req.Data)
	resp = new(standard.UpdateByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !usernamePattern.MatchString(req.Data.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	if !nicknamePattern.MatchString(req.Data.Nickname) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查昵称格式"
		return resp, nil
	}

	err = countUserByIDStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 用户不存在
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	req.Data.ID = req.ID
	_, err = updateUserByIDStmt.ExecContext(ctx, req.Data)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// DeleteByID 通过ID删除用户
func (srv *Service) DeleteByID(ctx context.Context, req *standard.DeleteByIDRequest) (resp *standard.DeleteByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	resp = new(standard.DeleteByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = countUserByIDStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	_, err = deleteUserByIDStmt.ExecContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "删除成功"

	return resp, nil
}

// UpdatePasswordByID 更新用户密码
func (srv *Service) UpdatePasswordByID(ctx context.Context, req *standard.UpdatePasswordByIDRequest) (resp *standard.UpdatePasswordByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	var user models.User
	user.SetPassword(req.Password)
	resp = new(standard.UpdatePasswordByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !passwordPattern.MatchString(req.Password) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	err = countUserByIDStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	user.ID = req.ID
	_, err = updateUserPasswordByIDStmt.ExecContext(ctx, user)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"

	return resp, nil
}

// VerifyPasswordByID 验证密码
func (srv *Service) VerifyPasswordByID(ctx context.Context, req *standard.VerifyPasswordByIDRequest) (resp *standard.VerifyPasswordByIDResponse, err error) {
	user := new(models.User)
	users := []*models.User{}
	user.SetPassword(req.Password)
	resp = new(standard.VerifyPasswordByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !passwordPattern.MatchString(req.Password) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	rows, err := queryUserByIDStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localUser models.User
		err = rows.StructScan(&localUser)
		if err == nil {
			users = append(users, &localUser)
		}
	}

	if len(users) <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	if users[0].Password != user.Password {
		resp.State = standard.State_SUCCESS
		resp.Message = "查询成功"
		resp.Data = false
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "查询成功"
	resp.Data = true

	return resp, nil
}

// VerifyPasswordByUsername 验证密码
func (srv *Service) VerifyPasswordByUsername(ctx context.Context, req *standard.VerifyPasswordByUsernameRequest) (resp *standard.VerifyPasswordByUsernameResponse, err error) {
	user := new(models.User)
	users := []*models.User{}
	user.SetPassword(req.Password)
	resp = new(standard.VerifyPasswordByUsernameResponse)

	if !passwordPattern.MatchString(req.Password) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	if !usernamePattern.MatchString(req.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	rows, err := queryUserByUsernameStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localUser models.User
		err = rows.StructScan(&localUser)
		if err == nil {
			users = append(users, &localUser)
		}
	}

	if len(users) <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	if users[0].Password != user.Password {
		resp.State = standard.State_SUCCESS
		resp.Message = "查询成功"
		resp.Data = false
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "查询成功"
	resp.Data = true

	return resp, nil
}

// CreateLabelByOwner 给指定用户创建标签
func (srv *Service) CreateLabelByOwner(ctx context.Context, req *standard.CreateLabelByOwnerRequest) (resp *standard.CreateLabelByOwnerResponse, err error) {
	var count uint64
	resp = new(standard.CreateLabelByOwnerResponse)

	err = countUserByIDStmt.GetContext(ctx, &count, map[string]interface{}{"ID": req.Owner})
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "用户不存在"
		return resp, nil
	}

	req.Label.Owner = req.Owner
	_, err = insertLabelByOwnerStmt.ExecContext(ctx, req.Label)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	return resp, nil
}

// QueryLabelByID 查询指定 Label 的信息
func (srv *Service) QueryLabelByID(ctx context.Context, req *standard.QueryLabelByIDRequest) (resp *standard.QueryLabelByIDResponse, err error) {
	labels := []*models.Label{}
	resp = new(standard.QueryLabelByIDResponse)

	rows, err := queryLabelByIDStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localLabel models.Label
		err = rows.StructScan(&localLabel)
		if err == nil {
			labels = append(labels, &localLabel)
		}
	}

	if len(labels) <= 0 {
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = labels[0].OutProtoStruct()
	resp.Message = "查询成功"

	return resp, nil
}

// UpdateLabelByID 更新指定 Label 的数据
func (srv *Service) UpdateLabelByID(ctx context.Context, req *standard.UpdateLabelByIDRequest) (resp *standard.UpdateLabelByIDResponse, err error) {
	var count uint64
	resp = new(standard.UpdateLabelByIDResponse)

	err = countLabelByIDStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	req.Data.ID = req.ID
	_, err = updateLabelByIDStmt.ExecContext(ctx, req.Data)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// DeleteLabelByID 删除指定 ID 的 Label
func (srv *Service) DeleteLabelByID(ctx context.Context, req *standard.DeleteLabelByIDRequest) (resp *standard.DeleteLabelByIDResponse, err error) {
	var count uint64
	resp = new(standard.DeleteLabelByIDResponse)

	err = countLabelByIDStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	_, err = deleteLabelByIDStmt.ExecContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "删除成功"
	return resp, nil

}

// QueryLabelByOwner 查询指定 Owner 的所有标签
func (srv *Service) QueryLabelByOwner(ctx context.Context, req *standard.QueryLabelByOwnerRequest) (resp *standard.QueryLabelByOwnerResponse, err error) {
	var count uint64
	labels := []*models.Label{}
	stdlabels := []*standard.Label{}
	resp = new(standard.QueryLabelByOwnerResponse)

	// 插总数
	err = countLabelByOwnerStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	// 查当前页
	rows, err := queryLabelByOwnerStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localLabel models.Label
		err = rows.StructScan(&localLabel)
		if err == nil {
			labels = append(labels, &localLabel)
		}
	}

	for _, label := range labels {
		stdlabels = append(stdlabels, label.OutProtoStruct())
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "查询成功"
	resp.Data = stdlabels
	resp.Total = count
	return resp, nil
}
