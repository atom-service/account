package provider

import (
	"context"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
	validators "github.com/grpcbrick/account/validators"
)

// CreateGroup 创建分组
func (srv *Service) CreateGroup(ctx context.Context, req *standard.CreateGroupRequest) (resp *standard.CreateGroupResponse, err error) {
	resp = new(standard.CreateGroupResponse)

	if ok, msg := validators.GroupName(req.Name); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.GroupClass(req.Class); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.GroupState(req.State); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.GroupDescription(req.Description); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	// 查询 用户名是否已经存在
	count, err := dao.CountGroupByName(req.Name)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if count > 0 {
		resp.State = standard.State_GROUP_ALREADY_EXISTS
		resp.Message = "该分组已存在"
		return resp, nil
	}

	err = dao.CreateGroup(req.Name, req.Class, req.State, req.Description)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	return resp, nil
}

// TODO: 待实现
func (srv *Service) QueryGroups(ctx context.Context, req *standard.QueryGroupsRequest) (resp *standard.QueryGroupsResponse, err error) {
	resp = new(standard.QueryGroupsResponse)
	return nil, nil
}

// QueryGroupByID 通过 ID 查询组信息
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

// DeleteGroupByID 通过 ID 删除分支
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

	if ok, msg := validators.GroupName(req.Name); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

	if ok, msg := validators.GroupClass(req.Class); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

	if ok, msg := validators.GroupState(req.State); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

	if ok, msg := validators.GroupDescription(req.Description); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
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

// AddUserToGroupByID 添加用户进组
func (srv *Service) AddUserToGroupByID(ctx context.Context, req *standard.AddUserToGroupByIDRequest) (resp *standard.AddUserToGroupByIDResponse, err error) {
	resp = new(standard.AddUserToGroupByIDResponse)
	if req.ID == 0 || req.GroupID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的分组 ID"
		return resp, nil
	}

	groupCount, err := dao.CountGroupByID(req.GroupID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if groupCount <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	userCount, err := dao.CountUserByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if userCount <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	already, err := dao.IsAlreadyInGroup(req.GroupID, req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	// 已经在组里了
	if already == true {
		resp.State = standard.State_SUCCESS
		resp.Message = "添加成功"
		return resp, nil
	}

	// 是否已在该组
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

// RemoveUserFromGroupByID 将用户移出组
func (srv *Service) RemoveUserFromGroupByID(ctx context.Context, req *standard.RemoveUserFromGroupByIDRequest) (resp *standard.RemoveUserFromGroupByIDResponse, err error) {
	resp = new(standard.RemoveUserFromGroupByIDResponse)
	if req.ID == 0 || req.UserID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	groupCount, err := dao.CountGroupByID(req.ID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if groupCount <= 0 { // 没有找到
		resp.State = standard.State_GROUP_NOT_EXIST
		resp.Message = "该分组不存在"
		return resp, nil
	}

	userCount, err := dao.CountUserByID(req.UserID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	if userCount <= 0 { // 没有找到用户
		resp.State = standard.State_USER_NOT_EXIST
		resp.Message = "该用户不存在"
		return resp, nil
	}

	already, err := dao.IsAlreadyInGroup(req.ID, req.UserID)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	// 不存在
	if already == false {
		resp.State = standard.State_SUCCESS
		resp.Message = "移除成功"
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
