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
  updateLabelByID: grpc.MethodDefinition<standard_pb.UpdateLabelByIDRequest, standard_pb.UpdateLabelByIDResponse>;
  deleteLabelByID: grpc.MethodDefinition<standard_pb.DeleteLabelByIDRequest, standard_pb.DeleteLabelByIDResponse>;
  queryLabelByOwner: grpc.MethodDefinition<standard_pb.QueryLabelByOwnerRequest, standard_pb.QueryLabelByOwnerResponse>;
  createLabelByOwner: grpc.MethodDefinition<standard_pb.CreateLabelByOwnerRequest, standard_pb.CreateLabelByOwnerResponse>;
  removeLabelByOwner: grpc.MethodDefinition<standard_pb.RemoveLabelByOwnerRequest, standard_pb.RemoveLabelByOwnerResponse>;
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
  updateLabelByID(argument: standard_pb.UpdateLabelByIDRequest, callback: grpc.requestCallback<standard_pb.UpdateLabelByIDResponse>): grpc.ClientUnaryCall;
  updateLabelByID(argument: standard_pb.UpdateLabelByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelByIDResponse>): grpc.ClientUnaryCall;
  updateLabelByID(argument: standard_pb.UpdateLabelByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  deleteLabelByID(argument: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DeleteLabelByIDResponse>): grpc.ClientUnaryCall;
  queryLabelByOwner(argument: standard_pb.QueryLabelByOwnerRequest, callback: grpc.requestCallback<standard_pb.QueryLabelByOwnerResponse>): grpc.ClientUnaryCall;
  queryLabelByOwner(argument: standard_pb.QueryLabelByOwnerRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryLabelByOwnerResponse>): grpc.ClientUnaryCall;
  queryLabelByOwner(argument: standard_pb.QueryLabelByOwnerRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.QueryLabelByOwnerResponse>): grpc.ClientUnaryCall;
  createLabelByOwner(argument: standard_pb.CreateLabelByOwnerRequest, callback: grpc.requestCallback<standard_pb.CreateLabelByOwnerResponse>): grpc.ClientUnaryCall;
  createLabelByOwner(argument: standard_pb.CreateLabelByOwnerRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateLabelByOwnerResponse>): grpc.ClientUnaryCall;
  createLabelByOwner(argument: standard_pb.CreateLabelByOwnerRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateLabelByOwnerResponse>): grpc.ClientUnaryCall;
  removeLabelByOwner(argument: standard_pb.RemoveLabelByOwnerRequest, callback: grpc.requestCallback<standard_pb.RemoveLabelByOwnerResponse>): grpc.ClientUnaryCall;
  removeLabelByOwner(argument: standard_pb.RemoveLabelByOwnerRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveLabelByOwnerResponse>): grpc.ClientUnaryCall;
  removeLabelByOwner(argument: standard_pb.RemoveLabelByOwnerRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.RemoveLabelByOwnerResponse>): grpc.ClientUnaryCall;
}
