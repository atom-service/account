// package: standard
// file: standard.proto

/* tslint:disable */

import * as grpc from "grpc";
import * as standard_pb from "./standard_pb";

interface IAccountService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createUser: IAccountService_ICreateUser;
    queryUserByID: IAccountService_IQueryUserByID;
    deleteUserByID: IAccountService_IDeleteUserByID;
    queryUserByUsername: IAccountService_IQueryUserByUsername;
    updateUserPasswordByID: IAccountService_IUpdateUserPasswordByID;
    verifyUserPasswordByID: IAccountService_IVerifyUserPasswordByID;
    verifyUserPasswordByUsername: IAccountService_IVerifyUserPasswordByUsername;
    createLabel: IAccountService_ICreateLabel;
    queryLabelByID: IAccountService_IQueryLabelByID;
    deleteLabelByID: IAccountService_IDeleteLabelByID;
    updateLabelClassByID: IAccountService_IUpdateLabelClassByID;
    updateLabelStateByID: IAccountService_IUpdateLabelStateByID;
    updateLabelValueByID: IAccountService_IUpdateLabelValueByID;
    addLabelToUserByID: IAccountService_IAddLabelToUserByID;
    removeLabelFromUserByID: IAccountService_IRemoveLabelFromUserByID;
    createGroup: IAccountService_ICreateGroup;
    queryGroupByID: IAccountService_IQueryGroupByID;
    deleteGroupByID: IAccountService_IDeleteGroupByID;
    updateGroupNameByID: IAccountService_IUpdateGroupNameByID;
    updateGroupClassByID: IAccountService_IUpdateGroupClassByID;
    updateGroupStateByID: IAccountService_IUpdateGroupStateByID;
    updateGroupDescriptionByID: IAccountService_IUpdateGroupDescriptionByID;
    addUserToGroupByID: IAccountService_IAddUserToGroupByID;
    removeUserFromGroupByID: IAccountService_IRemoveUserFromGroupByID;
}

interface IAccountService_ICreateUser extends grpc.MethodDefinition<standard_pb.CreateUserRequest, standard_pb.CreateUserResponse> {
    path: string; // "/standard.Account/CreateUser"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.CreateUserRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.CreateUserRequest>;
    responseSerialize: grpc.serialize<standard_pb.CreateUserResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.CreateUserResponse>;
}
interface IAccountService_IQueryUserByID extends grpc.MethodDefinition<standard_pb.QueryUserByIDRequest, standard_pb.QueryUserByIDResponse> {
    path: string; // "/standard.Account/QueryUserByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryUserByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryUserByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryUserByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryUserByIDResponse>;
}
interface IAccountService_IDeleteUserByID extends grpc.MethodDefinition<standard_pb.DeleteUserByIDRequest, standard_pb.DeleteUserByIDResponse> {
    path: string; // "/standard.Account/DeleteUserByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.DeleteUserByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.DeleteUserByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.DeleteUserByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.DeleteUserByIDResponse>;
}
interface IAccountService_IQueryUserByUsername extends grpc.MethodDefinition<standard_pb.QueryUserByUsernameRequest, standard_pb.QueryUserByUsernameResponse> {
    path: string; // "/standard.Account/QueryUserByUsername"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryUserByUsernameRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryUserByUsernameRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryUserByUsernameResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryUserByUsernameResponse>;
}
interface IAccountService_IUpdateUserPasswordByID extends grpc.MethodDefinition<standard_pb.UpdateUserPasswordByIDRequest, standard_pb.UpdateUserPasswordByIDResponse> {
    path: string; // "/standard.Account/UpdateUserPasswordByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateUserPasswordByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateUserPasswordByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateUserPasswordByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateUserPasswordByIDResponse>;
}
interface IAccountService_IVerifyUserPasswordByID extends grpc.MethodDefinition<standard_pb.VerifyUserPasswordByIDRequest, standard_pb.VerifyUserPasswordByIDResponse> {
    path: string; // "/standard.Account/VerifyUserPasswordByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.VerifyUserPasswordByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.VerifyUserPasswordByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.VerifyUserPasswordByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.VerifyUserPasswordByIDResponse>;
}
interface IAccountService_IVerifyUserPasswordByUsername extends grpc.MethodDefinition<standard_pb.VerifyUserPasswordByUsernameRequest, standard_pb.VerifyUserPasswordByUsernameResponse> {
    path: string; // "/standard.Account/VerifyUserPasswordByUsername"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.VerifyUserPasswordByUsernameRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.VerifyUserPasswordByUsernameRequest>;
    responseSerialize: grpc.serialize<standard_pb.VerifyUserPasswordByUsernameResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.VerifyUserPasswordByUsernameResponse>;
}
interface IAccountService_ICreateLabel extends grpc.MethodDefinition<standard_pb.CreateLabelRequest, standard_pb.CreateLabelResponse> {
    path: string; // "/standard.Account/CreateLabel"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.CreateLabelRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.CreateLabelRequest>;
    responseSerialize: grpc.serialize<standard_pb.CreateLabelResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.CreateLabelResponse>;
}
interface IAccountService_IQueryLabelByID extends grpc.MethodDefinition<standard_pb.QueryLabelByIDRequest, standard_pb.QueryLabelByIDResponse> {
    path: string; // "/standard.Account/QueryLabelByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryLabelByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryLabelByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryLabelByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryLabelByIDResponse>;
}
interface IAccountService_IDeleteLabelByID extends grpc.MethodDefinition<standard_pb.DeleteLabelByIDRequest, standard_pb.DeleteLabelByIDResponse> {
    path: string; // "/standard.Account/DeleteLabelByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.DeleteLabelByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.DeleteLabelByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.DeleteLabelByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.DeleteLabelByIDResponse>;
}
interface IAccountService_IUpdateLabelClassByID extends grpc.MethodDefinition<standard_pb.UpdateLabelClassByIDRequest, standard_pb.UpdateLabelClassByIDResponse> {
    path: string; // "/standard.Account/UpdateLabelClassByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateLabelClassByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateLabelClassByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateLabelClassByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateLabelClassByIDResponse>;
}
interface IAccountService_IUpdateLabelStateByID extends grpc.MethodDefinition<standard_pb.UpdateLabelStateByIDRequest, standard_pb.UpdateLabelStateByIDResponse> {
    path: string; // "/standard.Account/UpdateLabelStateByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateLabelStateByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateLabelStateByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateLabelStateByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateLabelStateByIDResponse>;
}
interface IAccountService_IUpdateLabelValueByID extends grpc.MethodDefinition<standard_pb.UpdateLabelValueByIDRequest, standard_pb.UpdateLabelValueByIDResponse> {
    path: string; // "/standard.Account/UpdateLabelValueByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateLabelValueByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateLabelValueByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateLabelValueByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateLabelValueByIDResponse>;
}
interface IAccountService_IAddLabelToUserByID extends grpc.MethodDefinition<standard_pb.AddLabelToUserByIDRequest, standard_pb.AddLabelToUserByIDResponse> {
    path: string; // "/standard.Account/AddLabelToUserByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.AddLabelToUserByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.AddLabelToUserByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.AddLabelToUserByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.AddLabelToUserByIDResponse>;
}
interface IAccountService_IRemoveLabelFromUserByID extends grpc.MethodDefinition<standard_pb.RemoveLabelFromUserByIDRequest, standard_pb.RemoveLabelFromUserByIDResponse> {
    path: string; // "/standard.Account/RemoveLabelFromUserByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.RemoveLabelFromUserByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.RemoveLabelFromUserByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.RemoveLabelFromUserByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.RemoveLabelFromUserByIDResponse>;
}
interface IAccountService_ICreateGroup extends grpc.MethodDefinition<standard_pb.CreateGroupRequest, standard_pb.CreateGroupResponse> {
    path: string; // "/standard.Account/CreateGroup"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.CreateGroupRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.CreateGroupRequest>;
    responseSerialize: grpc.serialize<standard_pb.CreateGroupResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.CreateGroupResponse>;
}
interface IAccountService_IQueryGroupByID extends grpc.MethodDefinition<standard_pb.QueryGroupByIDRequest, standard_pb.QueryGroupByIDResponse> {
    path: string; // "/standard.Account/QueryGroupByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryGroupByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryGroupByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryGroupByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryGroupByIDResponse>;
}
interface IAccountService_IDeleteGroupByID extends grpc.MethodDefinition<standard_pb.DeleteGroupByIDRequest, standard_pb.DeleteGroupByIDResponse> {
    path: string; // "/standard.Account/DeleteGroupByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.DeleteGroupByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.DeleteGroupByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.DeleteGroupByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.DeleteGroupByIDResponse>;
}
interface IAccountService_IUpdateGroupNameByID extends grpc.MethodDefinition<standard_pb.UpdateGroupNameByIDRequest, standard_pb.UpdateGroupNameByIDResponse> {
    path: string; // "/standard.Account/UpdateGroupNameByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateGroupNameByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateGroupNameByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateGroupNameByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateGroupNameByIDResponse>;
}
interface IAccountService_IUpdateGroupClassByID extends grpc.MethodDefinition<standard_pb.UpdateGroupClassByIDRequest, standard_pb.UpdateGroupClassByIDResponse> {
    path: string; // "/standard.Account/UpdateGroupClassByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateGroupClassByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateGroupClassByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateGroupClassByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateGroupClassByIDResponse>;
}
interface IAccountService_IUpdateGroupStateByID extends grpc.MethodDefinition<standard_pb.UpdateGroupStateByIDRequest, standard_pb.UpdateGroupStateByIDResponse> {
    path: string; // "/standard.Account/UpdateGroupStateByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateGroupStateByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateGroupStateByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateGroupStateByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateGroupStateByIDResponse>;
}
interface IAccountService_IUpdateGroupDescriptionByID extends grpc.MethodDefinition<standard_pb.UpdateGroupDescriptionByIDRequest, standard_pb.UpdateGroupDescriptionByIDResponse> {
    path: string; // "/standard.Account/UpdateGroupDescriptionByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateGroupDescriptionByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateGroupDescriptionByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateGroupDescriptionByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateGroupDescriptionByIDResponse>;
}
interface IAccountService_IAddUserToGroupByID extends grpc.MethodDefinition<standard_pb.AddUserToGroupByIDRequest, standard_pb.AddUserToGroupByIDResponse> {
    path: string; // "/standard.Account/AddUserToGroupByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.AddUserToGroupByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.AddUserToGroupByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.AddUserToGroupByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.AddUserToGroupByIDResponse>;
}
interface IAccountService_IRemoveUserFromGroupByID extends grpc.MethodDefinition<standard_pb.RemoveUserFromGroupByIDRequest, standard_pb.RemoveUserFromGroupByIDResponse> {
    path: string; // "/standard.Account/RemoveUserFromGroupByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.RemoveUserFromGroupByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.RemoveUserFromGroupByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.RemoveUserFromGroupByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.RemoveUserFromGroupByIDResponse>;
}

export const AccountService: IAccountService;

export interface IAccountServer {
    createUser: grpc.handleUnaryCall<standard_pb.CreateUserRequest, standard_pb.CreateUserResponse>;
    queryUserByID: grpc.handleUnaryCall<standard_pb.QueryUserByIDRequest, standard_pb.QueryUserByIDResponse>;
    deleteUserByID: grpc.handleUnaryCall<standard_pb.DeleteUserByIDRequest, standard_pb.DeleteUserByIDResponse>;
    queryUserByUsername: grpc.handleUnaryCall<standard_pb.QueryUserByUsernameRequest, standard_pb.QueryUserByUsernameResponse>;
    updateUserPasswordByID: grpc.handleUnaryCall<standard_pb.UpdateUserPasswordByIDRequest, standard_pb.UpdateUserPasswordByIDResponse>;
    verifyUserPasswordByID: grpc.handleUnaryCall<standard_pb.VerifyUserPasswordByIDRequest, standard_pb.VerifyUserPasswordByIDResponse>;
    verifyUserPasswordByUsername: grpc.handleUnaryCall<standard_pb.VerifyUserPasswordByUsernameRequest, standard_pb.VerifyUserPasswordByUsernameResponse>;
    createLabel: grpc.handleUnaryCall<standard_pb.CreateLabelRequest, standard_pb.CreateLabelResponse>;
    queryLabelByID: grpc.handleUnaryCall<standard_pb.QueryLabelByIDRequest, standard_pb.QueryLabelByIDResponse>;
    deleteLabelByID: grpc.handleUnaryCall<standard_pb.DeleteLabelByIDRequest, standard_pb.DeleteLabelByIDResponse>;
    updateLabelClassByID: grpc.handleUnaryCall<standard_pb.UpdateLabelClassByIDRequest, standard_pb.UpdateLabelClassByIDResponse>;
    updateLabelStateByID: grpc.handleUnaryCall<standard_pb.UpdateLabelStateByIDRequest, standard_pb.UpdateLabelStateByIDResponse>;
    updateLabelValueByID: grpc.handleUnaryCall<standard_pb.UpdateLabelValueByIDRequest, standard_pb.UpdateLabelValueByIDResponse>;
    addLabelToUserByID: grpc.handleUnaryCall<standard_pb.AddLabelToUserByIDRequest, standard_pb.AddLabelToUserByIDResponse>;
    removeLabelFromUserByID: grpc.handleUnaryCall<standard_pb.RemoveLabelFromUserByIDRequest, standard_pb.RemoveLabelFromUserByIDResponse>;
    createGroup: grpc.handleUnaryCall<standard_pb.CreateGroupRequest, standard_pb.CreateGroupResponse>;
    queryGroupByID: grpc.handleUnaryCall<standard_pb.QueryGroupByIDRequest, standard_pb.QueryGroupByIDResponse>;
    deleteGroupByID: grpc.handleUnaryCall<standard_pb.DeleteGroupByIDRequest, standard_pb.DeleteGroupByIDResponse>;
    updateGroupNameByID: grpc.handleUnaryCall<standard_pb.UpdateGroupNameByIDRequest, standard_pb.UpdateGroupNameByIDResponse>;
    updateGroupClassByID: grpc.handleUnaryCall<standard_pb.UpdateGroupClassByIDRequest, standard_pb.UpdateGroupClassByIDResponse>;
    updateGroupStateByID: grpc.handleUnaryCall<standard_pb.UpdateGroupStateByIDRequest, standard_pb.UpdateGroupStateByIDResponse>;
    updateGroupDescriptionByID: grpc.handleUnaryCall<standard_pb.UpdateGroupDescriptionByIDRequest, standard_pb.UpdateGroupDescriptionByIDResponse>;
    addUserToGroupByID: grpc.handleUnaryCall<standard_pb.AddUserToGroupByIDRequest, standard_pb.AddUserToGroupByIDResponse>;
    removeUserFromGroupByID: grpc.handleUnaryCall<standard_pb.RemoveUserFromGroupByIDRequest, standard_pb.RemoveUserFromGroupByIDResponse>;
}

export interface IAccountClient {
    createUser(request: standard_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    deleteUserByID(request: standard_pb.DeleteUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    deleteUserByID(request: standard_pb.DeleteUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    deleteUserByID(request: standard_pb.DeleteUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    createLabel(request: standard_pb.CreateLabelRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    createLabel(request: standard_pb.CreateLabelRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    createLabel(request: standard_pb.CreateLabelRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    createGroup(request: standard_pb.CreateGroupRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    createGroup(request: standard_pb.CreateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    createGroup(request: standard_pb.CreateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    queryGroupByID(request: standard_pb.QueryGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    queryGroupByID(request: standard_pb.QueryGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    queryGroupByID(request: standard_pb.QueryGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
    removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
    removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
}

export class AccountClient extends grpc.Client implements IAccountClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createUser(request: standard_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteUserByID(request: standard_pb.DeleteUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteUserByID(request: standard_pb.DeleteUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteUserByID(request: standard_pb.DeleteUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteUserByIDResponse) => void): grpc.ClientUnaryCall;
    public queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    public queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    public queryUserByUsername(request: standard_pb.QueryUserByUsernameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByUsernameResponse) => void): grpc.ClientUnaryCall;
    public updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public updateUserPasswordByID(request: standard_pb.UpdateUserPasswordByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByID(request: standard_pb.VerifyUserPasswordByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByIDResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    public verifyUserPasswordByUsername(request: standard_pb.VerifyUserPasswordByUsernameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.VerifyUserPasswordByUsernameResponse) => void): grpc.ClientUnaryCall;
    public createLabel(request: standard_pb.CreateLabelRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    public createLabel(request: standard_pb.CreateLabelRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    public createLabel(request: standard_pb.CreateLabelRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelClassByID(request: standard_pb.UpdateLabelClassByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelStateByID(request: standard_pb.UpdateLabelStateByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelValueByID(request: standard_pb.UpdateLabelValueByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelValueByIDResponse) => void): grpc.ClientUnaryCall;
    public addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    public addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    public addLabelToUserByID(request: standard_pb.AddLabelToUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.AddLabelToUserByIDResponse) => void): grpc.ClientUnaryCall;
    public removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    public removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    public removeLabelFromUserByID(request: standard_pb.RemoveLabelFromUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveLabelFromUserByIDResponse) => void): grpc.ClientUnaryCall;
    public createGroup(request: standard_pb.CreateGroupRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public createGroup(request: standard_pb.CreateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public createGroup(request: standard_pb.CreateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public queryGroupByID(request: standard_pb.QueryGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public queryGroupByID(request: standard_pb.QueryGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public queryGroupByID(request: standard_pb.QueryGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteGroupByID(request: standard_pb.DeleteGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupNameByID(request: standard_pb.UpdateGroupNameByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupNameByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupClassByID(request: standard_pb.UpdateGroupClassByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupClassByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupStateByID(request: standard_pb.UpdateGroupStateByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupStateByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    public updateGroupDescriptionByID(request: standard_pb.UpdateGroupDescriptionByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateGroupDescriptionByIDResponse) => void): grpc.ClientUnaryCall;
    public addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public addUserToGroupByID(request: standard_pb.AddUserToGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.AddUserToGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
    public removeUserFromGroupByID(request: standard_pb.RemoveUserFromGroupByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.RemoveUserFromGroupByIDResponse) => void): grpc.ClientUnaryCall;
}
