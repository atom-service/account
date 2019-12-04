package provider

import (
	"context"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
	validators "github.com/grpcbrick/account/validators"
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
	resp = new(standard.CreateUserResponse)
	if ok, msg := validators.Username(req.Username); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.Nickname(req.Nickname); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.Password(req.Password); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

	id, err := dao.CreateUser(req.Class, req.Nickname, req.Username, req.Password, req.Inviter)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	// 查询数据
	queryResult, err := srv.QueryUserByID(ctx, &standard.QueryUserByIDRequest{ID: uint64(id)})
	if err != nil {
		resp.State = standard.State_SERVICE_ERROR
		resp.Message = err.Error()
		return resp, nil
	}

	// 查询失败了
	if queryResult.State != standard.State_SUCCESS {
		resp.State = queryResult.State
		resp.Message = queryResult.Message
		return resp, nil
	}

	resp.State = queryResult.State
	resp.Data = queryResult.Data
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

// TODO:
func (srv *Service) QueryUsers(ctx context.Context, req *standard.QueryUsersRequest) (resp *standard.QueryUsersResponse, err error) {
	resp = new(standard.QueryUsersResponse)
	return nil, nil
}

// TODO:
func (srv *Service) QueryUsersByInviter(ctx context.Context, req *standard.QueryUsersByInviterRequest) (resp *standard.QueryUsersByInviterResponse, err error) {
	resp = new(standard.QueryUsersByInviterResponse)
	return nil, nil
}

// QueryUserByUsername 通过 用户名 查询用户
func (srv *Service) QueryUserByUsername(ctx context.Context, req *standard.QueryUserByUsernameRequest) (resp *standard.QueryUserByUsernameResponse, err error) {
	resp = new(standard.QueryUserByUsernameResponse)

	if ok, msg := validators.Username(req.Username); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

// DeleteUserByID 通过 ID 删除用户（逻辑删除）
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

// UpdateUserPasswordByID 通过 ID 更新用户密码
func (srv *Service) UpdateUserPasswordByID(ctx context.Context, req *standard.UpdateUserPasswordByIDRequest) (resp *standard.UpdateUserPasswordByIDResponse, err error) {
	resp = new(standard.UpdateUserPasswordByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if ok, msg := validators.Password(req.Password); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

// VerifyUserPasswordByID 通过 ID 验证用户密码
func (srv *Service) VerifyUserPasswordByID(ctx context.Context, req *standard.VerifyUserPasswordByIDRequest) (resp *standard.VerifyUserPasswordByIDResponse, err error) {
	resp = new(standard.VerifyUserPasswordByIDResponse)
	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	if req.Password == "" {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的密码"
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
		resp.State = standard.State_FAILURE
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

	if req.Password == "" {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的密码"
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
		resp.State = standard.State_FAILURE
		resp.Message = "账户或密码错误"
		return
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "验证成功"
	return resp, nil
}
