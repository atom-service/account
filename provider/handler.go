package provider

import (
	"context"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
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
	if !usernamePattern.MatchString(req.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	if !nicknamePattern.MatchString(req.Nickname) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查昵称格式"
		return resp, nil
	}

	if !passwordPattern.MatchString(req.Password) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查密码格式"
		return resp, nil
	}

	// 查询 用户名是否已经存在
	count, err := dao.CountUserByUsername(req.Username)
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

	err = dao.CreateUser(req.Class, req.Nickname, req.Nickname, req.Password, req.Inviter)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	return resp, nil
}

// QueryUserByID 通过ID查询用户
func (srv *Service) QueryUserByID(ctx context.Context, req *standard.QueryUserByIDRequest) (resp *standard.QueryUserByIDResponse, err error) {
	resp = new(standard.QueryUserByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	user, err := dao.QueryUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = user.OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

// QueryUserByUsername 通过 用户名 查询用户
func (srv *Service) QueryUserByUsername(ctx context.Context, req *standard.QueryUserByUsernameRequest) (resp *standard.QueryUserByUsernameResponse, err error) {
	resp = new(standard.QueryUserByUsernameResponse)

	if !usernamePattern.MatchString(req.Username) {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "请检查用户名格式"
		return resp, nil
	}

	count, err := dao.CountUserByUsername(req.Username)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	user, err := dao.QueryUserByUsername(req.Username)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = user.OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

func (srv *Service) DeleteUserByID(ctx context.Context, req *standard.DeleteUserByIDRequest) (resp *standard.DeleteUserByIDResponse, err error) {
	resp = new(standard.DeleteUserByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	err = dao.DeleteUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "删除成功"
	return resp, nil
}

func (srv *Service) UpdateUserPasswordByID(ctx context.Context, req *standard.UpdateUserPasswordByIDRequest) (resp *standard.UpdateUserPasswordByIDResponse, err error) {
	resp = new(standard.UpdateUserPasswordByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	err = dao.UpdateUserPasswordByID(req.ID, req.Password)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

func (srv *Service) VerifyUserPasswordByID(ctx context.Context, req *standard.VerifyUserPasswordByIDRequest) (resp *standard.VerifyUserPasswordByIDResponse, err error) {
	resp = new(standard.VerifyUserPasswordByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	pass, err := dao.VerifyUserPasswordByID(req.ID, req.Password)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if pass == false {
		resp.State = standard.State_USER_VERIFY_FAILURE
		resp.Message = "账户或密码错误"
		return
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "验证成功"
	return resp, nil
}

// VerifyUserPasswordByUsername VerifyUserPasswordByUsername
func (srv *Service) VerifyUserPasswordByUsername(ctx context.Context, req *standard.VerifyUserPasswordByUsernameRequest) (resp *standard.VerifyUserPasswordByUsernameResponse, err error) {
	resp = new(standard.VerifyUserPasswordByUsernameResponse)
	if req.Username == "" {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 Username"
		return resp, nil
	}

	count, err := dao.CountUserByUsername(req.Username)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	user, err := dao.QueryUserByUsername(req.Username)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	pass, err := dao.VerifyUserPasswordByID(user.ID, req.Password)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if pass == false {
		resp.State = standard.State_USER_VERIFY_FAILURE
		resp.Message = "账户或密码错误"
		return
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "验证成功"
	return resp, nil
}

func (srv *Service) QueryLabelByID(ctx context.Context, req *standard.QueryLabelByIDRequest) (resp *standard.QueryLabelByIDResponse, err error) {
	resp = new(standard.QueryLabelByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	label, err := dao.QueryLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = label.OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

func (srv *Service) DeleteLabelByID(ctx context.Context, req *standard.DeleteLabelByIDRequest) (resp *standard.DeleteLabelByIDResponse, err error) {
	resp = new(standard.DeleteLabelByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	err = dao.DeleteLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "删除成功"
	return resp, nil
}

func (srv *Service) UpdateLabelClassByID(ctx context.Context, req *standard.UpdateLabelClassByIDRequest) (resp *standard.UpdateLabelClassByIDResponse, err error) {
	resp = new(standard.UpdateLabelClassByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	err = dao.UpdateLabelClassByID(req.ID, req.Class)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

func (srv *Service) UpdateLabelStateByID(ctx context.Context, req *standard.UpdateLabelStateByIDRequest) (resp *standard.UpdateLabelStateByIDResponse, err error) {
	resp = new(standard.UpdateLabelStateByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	err = dao.UpdateLabelStateByID(req.ID, req.State)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

func (srv *Service) UpdateLabelValueByID(ctx context.Context, req *standard.UpdateLabelValueByIDRequest) (resp *standard.UpdateLabelValueByIDResponse, err error) {
	resp = new(standard.UpdateLabelValueByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountLabelByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_LABEL_NOT_EXIST
		resp.Message = "该标签不存在"
		return resp, nil
	}

	err = dao.UpdateLabelValueByID(req.ID, req.Value)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// 标签关系操作
func (srv *Service) AddLabelToUserByID(ctx context.Context, req *standard.AddLabelToUserByIDRequest) (resp *standard.AddLabelToUserByIDResponse, err error) {
	resp = new(standard.AddLabelToUserByIDResponse)
	if req.ID == 0 || req.UserID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = dao.AddLabelToUserByID(req.ID, req.UserID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "添加成功"
	return resp, nil
}

func (srv *Service) RemoveLabelFromUserByID(ctx context.Context, req *standard.RemoveLabelFromUserByIDRequest) (resp *standard.RemoveLabelFromUserByIDResponse, err error) {
	resp = new(standard.RemoveLabelFromUserByIDResponse)
	if req.ID == 0 || req.LabelID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = dao.RemoveLabelFromUserByID(req.LabelID, req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "移除成功"
	return resp, nil
}

// 组操作
func (srv *Service) QueryGroupByID(ctx context.Context, req *standard.QueryGroupByIDRequest) (resp *standard.QueryGroupByIDResponse, err error) {
	resp = new(standard.QueryGroupByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	group, err := dao.QueryGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}
	resp.State = standard.State_SUCCESS
	resp.Data = group.OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}
func (srv *Service) DeleteGroupByID(ctx context.Context, req *standard.DeleteGroupByIDRequest) (resp *standard.DeleteGroupByIDResponse, err error) {
	resp = new(standard.DeleteGroupByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	err = dao.DeleteGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "删除成功"
	return resp, nil
}

// UpdateGroupNameByID 更新分组名称
func (srv *Service) UpdateGroupNameByID(ctx context.Context, req *standard.UpdateGroupNameByIDRequest) (resp *standard.UpdateGroupNameByIDResponse, err error) {
	resp = new(standard.UpdateGroupNameByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	err = dao.UpdateGroupNameByID(req.ID, req.Name)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// UpdateGroupClassByID 更新分组的 Class 信息
func (srv *Service) UpdateGroupClassByID(ctx context.Context, req *standard.UpdateGroupClassByIDRequest) (resp *standard.UpdateGroupClassByIDResponse, err error) {
	resp = new(standard.UpdateGroupClassByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	err = dao.UpdateGroupClassByID(req.ID, req.Class)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// UpdateGroupStateByID 更新分组的状态
func (srv *Service) UpdateGroupStateByID(ctx context.Context, req *standard.UpdateGroupStateByIDRequest) (resp *standard.UpdateGroupStateByIDResponse, err error) {
	resp = new(standard.UpdateGroupStateByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	err = dao.UpdateGroupStateByID(req.ID, req.State)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// UpdateGroupDescriptionByID 更新分组的介绍信息
func (srv *Service) UpdateGroupDescriptionByID(ctx context.Context, req *standard.UpdateGroupDescriptionByIDRequest) (resp *standard.UpdateGroupDescriptionByIDResponse, err error) {
	resp = new(standard.UpdateGroupDescriptionByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	count, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	err = dao.UpdateGroupDescriptionByID(req.ID, req.Description)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, nil
}

// 组关系操作
func (srv *Service) AddUserToGroupByID(ctx context.Context, req *standard.AddUserToGroupByIDRequest) (resp *standard.AddUserToGroupByIDResponse, err error) {
	resp = new(standard.AddUserToGroupByIDResponse)
	if req.ID == 0 || req.GroupID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = dao.AddUserToGroupByID(req.GroupID, req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "添加成功"
	return resp, nil
}
func (srv *Service) RemoveUserFromGroupByID(ctx context.Context, req *standard.RemoveUserFromGroupByIDRequest) (resp *standard.RemoveUserFromGroupByIDResponse, err error) {
	resp = new(standard.RemoveUserFromGroupByIDResponse)
	if req.ID == 0 || req.UserID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	err = dao.AddUserToGroupByID(req.ID, req.UserID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "移除成功"
	return resp, nil
}
