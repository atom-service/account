// GENERATED CODE -- DO NOT EDIT!

// package: standard
// file: standard.proto

import * as standard_pb from "./standard_pb";
import * as grpc from "grpc";

interface IAccountService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createUser: grpc.MethodDefinition<standard_pb.CreateUserRequest, standard_pb.CreateUserResponse>;
  queryUserByID: grpc.MethodDefinition<standard_pb.QueryUserByIDRequest, standard_pb.QueryUserByIDResponse>;
  deleteUserByID: grpc.MethodDefinition<standard_pb.DeleteUserByIDRequest, standard_pb.DeleteUserByIDResponse>;
  queryUserByUsername: grpc.MethodDefinition<standard_pb.QueryUserByUsernameRequest, standard_pb.QueryUserByUsernameResponse>;
  updateUserPasswordByID: grpc.MethodDefinition<standard_pb.UpdateUserPasswordByIDRequest, standard_pb.UpdateUserPasswordByIDResponse>;
  verifyUserPasswordByID: grpc.MethodDefinition<standard_pb.VerifyUserPasswordByIDRequest, standard_pb.VerifyUserPasswordByIDResponse>;
  verifyUserPasswordByUsername: grpc.MethodDefinition<standard_pb.VerifyUserPasswordByUsernameRequest, standard_pb.VerifyUserPasswordByUsernameResponse>;
  queryLabelByID: grpc.MethodDefinition<standard_pb.QueryLabelByIDRequest, standard_pb.QueryLabelByIDResponse>;
  deleteLabelByID: grpc.MethodDefinition<standard_pb.DeleteLabelByIDRequest, standard_pb.DeleteLabelByIDResponse>;
  updateLabelClassByID: grpc.MethodDefinition<standard_pb.UpdateLabelClassByIDRequest, standard_pb.UpdateLabelClassByIDResponse>;
  updateLabelStateByID: grpc.MethodDefinition<standard_pb.UpdateLabelStateByIDRequest, standard_pb.UpdateLabelStateByIDResponse>;
  updateLabelValueByID: grpc.MethodDefinition<standard_pb.UpdateLabelValueByIDRequest, standard_pb.UpdateLabelValueByIDResponse>;
  addLabelToUserByID: grpc.MethodDefinition<standard_pb.AddLabelToUserByIDRequest, standard_pb.AddLabelToUserByIDResponse>;
  removeLabelFromUserByID: grpc.MethodDefinition<standard_pb.RemoveLabelFromUserByIDRequest, standard_pb.RemoveLabelFromUserByIDResponse>;
  queryGroupByID: grpc.MethodDefinition<standard_pb.QueryGroupByIDRequest, standard_pb.QueryGroupByIDResponse>;
  deleteGroupByID: grpc.MethodDefinition<standard_pb.DeleteGroupByIDRequest, standard_pb.DeleteGroupByIDResponse>;
  updateGroupNameByID: grpc.MethodDefinition<standard_pb.UpdateGroupNameByIDRequest, standard_pb.UpdateGroupNameByIDResponse>;
  updateGroupClassByID: grpc.MethodDefinition<standard_pb.UpdateGroupClassByIDRequest, standard_pb.UpdateGroupClassByIDResponse>;
  updateGroupStateByID: grpc.MethodDefinition<standard_pb.UpdateGroupStateByIDRequest, standard_pb.UpdateGroupStateByIDResponse>;
  updateGroupDescriptionByID: grpc.MethodDefinition<standard_pb.UpdateGroupDescriptionByIDRequest, standard_pb.UpdateGroupDescriptionByIDResponse>;
  addUserToGroupByID: grpc.MethodDefinition<standard_pb.AddUserToGroupByIDRequest, standard_pb.AddUserToGroupByIDResponse>;
  removeUserFromGroupByID: grpc.MethodDefinition<standard_pb.RemoveUserFromGroupByIDRequest, standard_pb.RemoveUserFromGroupByIDResponse>;
}

export const AccountService: IAccountService;

export class AccountClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createUser(argument: standard_pb.CreateUserRequest, callback: grpc.requestCallback<standard_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  createUser(argument: standard_pb.CreateUserRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  createUser(argument: standard_pb.CreateUserRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  queryUserByID(argument: standard_pb.QueryUserByIDRequest, callback: grpc.requestCallback<standard_pb.QueryUserByIDResponse>): grpc.ClientUnaryCall;
  queryUserByID(argument: standard_pb.QueryUserByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryUserByIDResponse>): grpc.ClientUnaryCall;
  queryUserByID(argument: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryUserByIDResponse>): grpc.ClientUnaryCall;
  deleteUserByID(argument: standard_pb.DeleteUserByIDRequest, callback: grpc.requestCallback<standard_pb.DeleteUserByIDResponse>): grpc.ClientUnaryCall;
  deleteUserByID(argument: standard_pb.DeleteUserByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteUserByIDResponse>): grpc.ClientUnaryCall;
  deleteUserByID(argument: standard_pb.DeleteUserByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteUserByIDResponse>): grpc.ClientUnaryCall;
  queryUserByUsername(argument: standard_pb.QueryUserByUsernameRequest, callback: grpc.requestCallback<standard_pb.QueryUserByUsernameResponse>): grpc.ClientUnaryCall;
  queryUserByUsername(argument: standard_pb.QueryUserByUsernameRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryUserByUsernameResponse>): grpc.ClientUnaryCall;
  queryUserByUsername(argument: standard_pb.QueryUserByUsernameRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryUserByUsernameResponse>): grpc.ClientUnaryCall;
  updateUserPasswordByID(argument: standard_pb.UpdateUserPasswordByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  updateUserPasswordByID(argument: standard_pb.UpdateUserPasswordByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  updateUserPasswordByID(argument: standard_pb.UpdateUserPasswordByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByID(argument: standard_pb.VerifyUserPasswordByIDRequest, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByID(argument: standard_pb.VerifyUserPasswordByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByID(argument: standard_pb.VerifyUserPasswordByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByIDResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByUsername(argument: standard_pb.VerifyUserPasswordByUsernameRequest, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByUsernameResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByUsername(argument: standard_pb.VerifyUserPasswordByUsernameRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByUsernameResponse>): grpc.ClientUnaryCall;
  verifyUserPasswordByUsername(argument: standard_pb.VerifyUserPasswordByUsernameRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.VerifyUserPasswordByUsernameResponse>): grpc.ClientUnaryCall;
  queryLabelByID(argument: standard_pb.QueryLabelByIDRequest, callback: grpc.requestCallback<standard_pb.QueryLabelByIDResponse>): grpc.ClientUnaryCall;
  queryLabelByID(argument: standard_pb.QueryLabelByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryLabelByIDResponse>): grpc.ClientUnaryCall;
  queryLabelByID(argument: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  updateLabelClassByID(argument: standard_pb.UpdateLabelClassByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateLabelClassByIDResponse>): grpc.ClientUnaryCall;
  updateLabelClassByID(argument: standard_pb.UpdateLabelClassByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelClassByIDResponse>): grpc.ClientUnaryCall;
  updateLabelClassByID(argument: standard_pb.UpdateLabelClassByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelClassByIDResponse>): grpc.ClientUnaryCall;
  updateLabelStateByID(argument: standard_pb.UpdateLabelStateByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateLabelStateByIDResponse>): grpc.ClientUnaryCall;
  updateLabelStateByID(argument: standard_pb.UpdateLabelStateByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelStateByIDResponse>): grpc.ClientUnaryCall;
  updateLabelStateByID(argument: standard_pb.UpdateLabelStateByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelStateByIDResponse>): grpc.ClientUnaryCall;
  updateLabelValueByID(argument: standard_pb.UpdateLabelValueByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateLabelValueByIDResponse>): grpc.ClientUnaryCall;
  updateLabelValueByID(argument: standard_pb.UpdateLabelValueByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelValueByIDResponse>): grpc.ClientUnaryCall;
  updateLabelValueByID(argument: standard_pb.UpdateLabelValueByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelValueByIDResponse>): grpc.ClientUnaryCall;
  addLabelToUserByID(argument: standard_pb.AddLabelToUserByIDRequest, callback: grpc.requestCallback<standard_pb.AddLabelToUserByIDResponse>): grpc.ClientUnaryCall;
  addLabelToUserByID(argument: standard_pb.AddLabelToUserByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.AddLabelToUserByIDResponse>): grpc.ClientUnaryCall;
  addLabelToUserByID(argument: standard_pb.AddLabelToUserByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.AddLabelToUserByIDResponse>): grpc.ClientUnaryCall;
  removeLabelFromUserByID(argument: standard_pb.RemoveLabelFromUserByIDRequest, callback: grpc.requestCallback<standard_pb.RemoveLabelFromUserByIDResponse>): grpc.ClientUnaryCall;
  removeLabelFromUserByID(argument: standard_pb.RemoveLabelFromUserByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveLabelFromUserByIDResponse>): grpc.ClientUnaryCall;
  removeLabelFromUserByID(argument: standard_pb.RemoveLabelFromUserByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveLabelFromUserByIDResponse>): grpc.ClientUnaryCall;
  queryGroupByID(argument: standard_pb.QueryGroupByIDRequest, callback: grpc.requestCallback<standard_pb.QueryGroupByIDResponse>): grpc.ClientUnaryCall;
  queryGroupByID(argument: standard_pb.QueryGroupByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryGroupByIDResponse>): grpc.ClientUnaryCall;
  queryGroupByID(argument: standard_pb.QueryGroupByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryGroupByIDResponse>): grpc.ClientUnaryCall;
  deleteGroupByID(argument: standard_pb.DeleteGroupByIDRequest, callback: grpc.requestCallback<standard_pb.DeleteGroupByIDResponse>): grpc.ClientUnaryCall;
  deleteGroupByID(argument: standard_pb.DeleteGroupByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteGroupByIDResponse>): grpc.ClientUnaryCall;
  deleteGroupByID(argument: standard_pb.DeleteGroupByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteGroupByIDResponse>): grpc.ClientUnaryCall;
  updateGroupNameByID(argument: standard_pb.UpdateGroupNameByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateGroupNameByIDResponse>): grpc.ClientUnaryCall;
  updateGroupNameByID(argument: standard_pb.UpdateGroupNameByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupNameByIDResponse>): grpc.ClientUnaryCall;
  updateGroupNameByID(argument: standard_pb.UpdateGroupNameByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupNameByIDResponse>): grpc.ClientUnaryCall;
  updateGroupClassByID(argument: standard_pb.UpdateGroupClassByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateGroupClassByIDResponse>): grpc.ClientUnaryCall;
  updateGroupClassByID(argument: standard_pb.UpdateGroupClassByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupClassByIDResponse>): grpc.ClientUnaryCall;
  updateGroupClassByID(argument: standard_pb.UpdateGroupClassByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupClassByIDResponse>): grpc.ClientUnaryCall;
  updateGroupStateByID(argument: standard_pb.UpdateGroupStateByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateGroupStateByIDResponse>): grpc.ClientUnaryCall;
  updateGroupStateByID(argument: standard_pb.UpdateGroupStateByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupStateByIDResponse>): grpc.ClientUnaryCall;
  updateGroupStateByID(argument: standard_pb.UpdateGroupStateByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupStateByIDResponse>): grpc.ClientUnaryCall;
  updateGroupDescriptionByID(argument: standard_pb.UpdateGroupDescriptionByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateGroupDescriptionByIDResponse>): grpc.ClientUnaryCall;
  updateGroupDescriptionByID(argument: standard_pb.UpdateGroupDescriptionByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupDescriptionByIDResponse>): grpc.ClientUnaryCall;
  updateGroupDescriptionByID(argument: standard_pb.UpdateGroupDescriptionByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateGroupDescriptionByIDResponse>): grpc.ClientUnaryCall;
  addUserToGroupByID(argument: standard_pb.AddUserToGroupByIDRequest, callback: grpc.requestCallback<standard_pb.AddUserToGroupByIDResponse>): grpc.ClientUnaryCall;
  addUserToGroupByID(argument: standard_pb.AddUserToGroupByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.AddUserToGroupByIDResponse>): grpc.ClientUnaryCall;
  addUserToGroupByID(argument: standard_pb.AddUserToGroupByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.AddUserToGroupByIDResponse>): grpc.ClientUnaryCall;
  removeUserFromGroupByID(argument: standard_pb.RemoveUserFromGroupByIDRequest, callback: grpc.requestCallback<standard_pb.RemoveUserFromGroupByIDResponse>): grpc.ClientUnaryCall;
  removeUserFromGroupByID(argument: standard_pb.RemoveUserFromGroupByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveUserFromGroupByIDResponse>): grpc.ClientUnaryCall;
  removeUserFromGroupByID(argument: standard_pb.RemoveUserFromGroupByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveUserFromGroupByIDResponse>): grpc.ClientUnaryCall;
}
