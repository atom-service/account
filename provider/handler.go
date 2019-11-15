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
	return resp, nil
}

func (srv *Service) UpdateUserPasswordByID(ctx context.Context, req *standard.UpdateUserPasswordByIDRequest) (resp *standard.UpdateUserPasswordByIDResponse, err error) {
	resp = new(standard.UpdateUserPasswordByIDResponse)
	return resp, nil
}
func (srv *Service) VerifyUserPasswordByID(ctx context.Context, req *standard.VerifyUserPasswordByIDRequest) (resp *standard.VerifyUserPasswordByIDResponse, err error) {
	resp = new(standard.VerifyUserPasswordByIDResponse)
	return resp, nil
}
func (srv *Service) VerifyUserPasswordByUsername(ctx context.Context, req *standard.VerifyUserPasswordByUsernameRequest) (resp *standard.VerifyUserPasswordByUsernameResponse, err error) {
	resp = new(standard.VerifyUserPasswordByUsernameResponse)
	return resp, nil
}

// 标签操作
func (srv *Service) QueryLabelByID(ctx context.Context, req *standard.QueryLabelByIDRequest) (resp *standard.QueryLabelByIDResponse, err error) {
	resp = new(standard.QueryLabelByIDResponse)
	return resp, nil
}
func (srv *Service) DeleteLabelByID(ctx context.Context, req *standard.DeleteLabelByIDRequest) (resp *standard.DeleteLabelByIDResponse, err error) {
	resp = new(standard.DeleteLabelByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateLabelClassByID(ctx context.Context, req *standard.UpdateLabelClassByIDRequest) (resp *standard.UpdateLabelClassByIDResponse, err error) {
	resp = new(standard.UpdateLabelClassByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateLabelStateByID(ctx context.Context, req *standard.UpdateLabelStateByIDRequest) (resp *standard.UpdateLabelStateByIDResponse, err error) {
	resp = new(standard.UpdateLabelStateByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateLabelValueByID(ctx context.Context, req *standard.UpdateLabelValueByIDRequest) (resp *standard.UpdateLabelValueByIDResponse, err error) {
	resp = new(standard.UpdateLabelValueByIDResponse)
	return resp, nil
}

// 标签关系操作
func (srv *Service) AddLabelToUserByID(ctx context.Context, req *standard.AddLabelToUserByIDRequest) (resp *standard.AddLabelToUserByIDResponse, err error) {
	resp = new(standard.AddLabelToUserByIDResponse)
	return resp, nil
}
func (srv *Service) RemoveLabelFromUserByID(ctx context.Context, req *standard.RemoveLabelFromUserByIDRequest) (resp *standard.RemoveLabelFromUserByIDResponse, err error) {
	resp = new(standard.RemoveLabelFromUserByIDResponse)
	return resp, nil
}

// 组操作
func (srv *Service) QueryGroupByID(ctx context.Context, req *standard.QueryGroupByIDRequest) (resp *standard.QueryGroupByIDResponse, err error) {
	resp = new(standard.QueryGroupByIDResponse)
	return resp, nil
}
func (srv *Service) DeleteGroupByID(ctx context.Context, req *standard.DeleteGroupByIDRequest) (resp *standard.DeleteGroupByIDResponse, err error) {
	resp = new(standard.DeleteGroupByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateGroupNameByID(ctx context.Context, req *standard.UpdateGroupNameByIDRequest) (resp *standard.UpdateGroupNameByIDResponse, err error) {
	resp = new(standard.UpdateGroupNameByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateGroupClassByID(ctx context.Context, req *standard.UpdateGroupClassByIDRequest) (resp *standard.UpdateGroupClassByIDResponse, err error) {
	resp = new(standard.UpdateGroupClassByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateGroupStateByID(ctx context.Context, req *standard.UpdateGroupStateByIDRequest) (resp *standard.UpdateGroupStateByIDResponse, err error) {
	resp = new(standard.UpdateGroupStateByIDResponse)
	return resp, nil
}
func (srv *Service) UpdateGroupDescriptionByID(ctx context.Context, req *standard.UpdateGroupDescriptionByIDRequest) (resp *standard.UpdateGroupDescriptionByIDResponse, err error) {
	resp = new(standard.UpdateGroupDescriptionByIDResponse)
	return resp, nil
}

// 组关系操作
func (srv *Service) AddUserToGroupByID(ctx context.Context, req *standard.AddUserToGroupByIDRequest) (resp *standard.AddUserToGroupByIDResponse, err error) {
	resp = new(standard.AddUserToGroupByIDResponse)
	return resp, nil
}
func (srv *Service) RemoveUserFromGroupByID(ctx context.Context, req *standard.RemoveUserFromGroupByIDRequest) (resp *standard.RemoveUserFromGroupByIDResponse, err error) {
	resp = new(standard.RemoveUserFromGroupByIDResponse)
	return resp, nil
}
