package provider

import (
	"context"

	"github.com/yinxulai/goutils/restful"

	"github.com/grpcbrick/account/database"
	"github.com/grpcbrick/account/models"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/pattern"
)

// NewService NewService
func NewService() *Service {
	service := new(Service)
	return service
}

// Service Service
type Service struct {
}

// CreateUser 创建用户
func (srv *Service) CreateUser(ctx context.Context, req *standard.CreateUserRequest) (resp *standard.CreateUserResponse, err error) {
	var count uint64
	var user models.User
	user.SetPassword(req.Password)
	resp = new(standard.CreateUserResponse)

	if !pattern.Username.MatchString(req.Username) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	if !pattern.Password.MatchString(req.Password) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	// 查询 用户名是否已经存在
	err = database.CountUserByUsernameNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count > 0 {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = "该用户已存在"
		return resp, nil
	}

	// 执行插入
	req.Password = user.Password // 重新赋值加密过后的密码
	_, err = database.InsertUserNamedStmt.ExecContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "创建成功"

	return resp, nil
}

// QueryUserByID 通过ID查询用户
func (srv *Service) QueryUserByID(ctx context.Context, req *standard.QueryUserByIDRequest) (resp *standard.QueryUserByIDResponse, err error) {
	users := []*models.User{}
	resp = new(standard.QueryUserByIDResponse)

	if req.ID == 0 {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "无效的 ID"
		return resp, nil
	}

	rows, err := database.QueryUserByIDNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localUser models.User
		err = rows.StructScan(&localUser)
		if err == nil {
			localUser.Password = "secret field"
			users = append(users, &localUser)
		}
	}

	if len(users) <= 0 { // 没有找到用户
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Data = users[0].OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

// QueryUserByUsername 通过ID查询用户
func (srv *Service) QueryUserByUsername(ctx context.Context, req *standard.QueryUserByUsernameRequest) (resp *standard.QueryUserByUsernameResponse, err error) {
	users := []*models.User{}
	resp = new(standard.QueryUserByUsernameResponse)

	if !pattern.Username.MatchString(req.Username) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	rows, err := database.QueryUserByUsernameNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
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
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Data = users[0].OutProtoStruct()
	resp.Message = "查询成功"

	return resp, nil
}

// UpdateUserByID 通过ID更新用户
func (srv *Service) UpdateUserByID(ctx context.Context, req *standard.UpdateUserByIDRequest) (resp *standard.UpdateUserByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	user := new(models.User)
	user.LoadProtoStruct(req.Data)
	resp = new(standard.UpdateUserByIDResponse)

	if req.ID == 0 {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !pattern.Username.MatchString(req.Data.Username) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	if !pattern.Nickname.MatchString(req.Data.Nickname) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查昵称格式"
		return resp, nil
	}

	err = database.CountUserByIDNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 用户不存在
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	req.Data.ID = req.ID
	_, err = database.UpdateUserByIDNamedStmt.ExecContext(ctx, req.Data)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "更新成功"
	return resp, nil
}

// DeleteUserByID 通过ID删除用户
func (srv *Service) DeleteUserByID(ctx context.Context, req *standard.DeleteUserByIDRequest) (resp *standard.DeleteUserByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	resp = new(standard.DeleteUserByIDResponse)

	if req.ID == 0 {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = database.CountUserByIDNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	_, err = database.DeleteUserByIDNamedStmt.ExecContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "删除成功"

	return resp, nil
}

// UpdateUserPasswordByID 更新用户密码
func (srv *Service) UpdateUserPasswordByID(ctx context.Context, req *standard.UpdateUserPasswordByIDRequest) (resp *standard.UpdateUserPasswordByIDResponse, err error) {
	// 检查是否存在该记录
	var count uint64
	var user models.User
	user.SetPassword(req.Password)
	resp = new(standard.UpdateUserPasswordByIDResponse)

	if req.ID == 0 {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !pattern.Password.MatchString(req.Password) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	err = database.CountUserByIDNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	user.ID = req.ID
	_, err = database.UpdateUserPasswordByIDNamedStmt.ExecContext(ctx, user)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "更新成功"

	return resp, nil
}

// VerifyUserPasswordByID 验证密码
func (srv *Service) VerifyUserPasswordByID(ctx context.Context, req *standard.VerifyUserPasswordByIDRequest) (resp *standard.VerifyUserPasswordByIDResponse, err error) {
	user := new(models.User)
	users := []*models.User{}
	user.SetPassword(req.Password)
	resp = new(standard.VerifyUserPasswordByIDResponse)

	if req.ID == 0 {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if !pattern.Password.MatchString(req.Password) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	rows, err := database.QueryUserByIDNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
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
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	if users[0].Password != user.Password {
		resp.State = uint64(restful.OK)
		resp.Message = "查询成功"
		resp.Data = false
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "查询成功"
	resp.Data = true

	return resp, nil
}

// VerifyUserPasswordByUsername 验证密码
func (srv *Service) VerifyUserPasswordByUsername(ctx context.Context, req *standard.VerifyUserPasswordByUsernameRequest) (resp *standard.VerifyUserPasswordByUsernameResponse, err error) {
	user := new(models.User)
	users := []*models.User{}
	user.SetPassword(req.Password)
	resp = new(standard.VerifyUserPasswordByUsernameResponse)

	if !pattern.Password.MatchString(req.Password) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	if !pattern.Username.MatchString(req.Username) {
		resp.State = uint64(restful.BADREQUEST)
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	rows, err := database.QueryUserByUsernameNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
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
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该用户不存在"
		return resp, nil
	}

	if users[0].Password != user.Password {
		resp.State = uint64(restful.OK)
		resp.Message = "查询成功"
		resp.Data = false
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "查询成功"
	resp.Data = true

	return resp, nil
}

// CreateLabelByOwner 给指定用户创建标签
func (srv *Service) CreateLabelByOwner(ctx context.Context, req *standard.CreateLabelByOwnerRequest) (resp *standard.CreateLabelByOwnerResponse, err error) {
	var count uint64
	resp = new(standard.CreateLabelByOwnerResponse)

	err = database.CountUserByIDNamedStmt.GetContext(ctx, &count, map[string]interface{}{"ID": req.Owner})
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "用户不存在"
		return resp, nil
	}

	req.Label.Owner = req.Owner
	_, err = database.InsertLabelByOwnerNamedStmt.ExecContext(ctx, req.Label)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "创建成功"
	return resp, nil
}

// QueryLabelByID 查询指定 Label 的信息
func (srv *Service) QueryLabelByID(ctx context.Context, req *standard.QueryLabelByIDRequest) (resp *standard.QueryLabelByIDResponse, err error) {
	labels := []*models.Label{}
	resp = new(standard.QueryLabelByIDResponse)

	rows, err := database.QueryLabelByIDNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
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
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该标签不存在"
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Data = labels[0].OutProtoStruct()
	resp.Message = "查询成功"

	return resp, nil
}

// UpdateLabelByID 更新指定 Label 的数据
func (srv *Service) UpdateLabelByID(ctx context.Context, req *standard.UpdateLabelByIDRequest) (resp *standard.UpdateLabelByIDResponse, err error) {
	var count uint64
	resp = new(standard.UpdateLabelByIDResponse)

	err = database.CountLabelByIDNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该标签不存在"
		return resp, nil
	}

	req.Data.ID = req.ID
	_, err = database.UpdateLabelByIDNamedStmt.ExecContext(ctx, req.Data)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
	resp.Message = "更新成功"
	return resp, nil
}

// DeleteLabelByID 删除指定 ID 的 Label
func (srv *Service) DeleteLabelByID(ctx context.Context, req *standard.DeleteLabelByIDRequest) (resp *standard.DeleteLabelByIDResponse, err error) {
	var count uint64
	resp = new(standard.DeleteLabelByIDResponse)

	err = database.CountLabelByIDNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 {
		resp.State = uint64(restful.NOTFOUND)
		resp.Message = "该标签不存在"
		return resp, nil
	}

	_, err = database.DeleteLabelByIDNamedStmt.ExecContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = uint64(restful.OK)
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
	err = database.CountLabelByOwnerNamedStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
		resp.Message = err.Error()
		return resp, nil
	}

	// 查当前页
	rows, err := database.QueryLabelByOwnerNamedStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = uint64(restful.INTERNALSERVERERROR)
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

	resp.State = uint64(restful.OK)
	resp.Message = "查询成功"
	resp.Data = stdlabels
	resp.Total = count
	return resp, nil
}