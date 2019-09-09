// package: standard
// file: standard.proto

/* tslint:disable */

import * as grpc from "grpc";
import * as standard_pb from "./standard_pb";

interface IAccountService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createUser: IAccountService_ICreateUser;
    queryUserByID: IAccountService_IQueryUserByID;
    updateUserByID: IAccountService_IUpdateUserByID;
    deleteUserByID: IAccountService_IDeleteUserByID;
    queryUserByUsername: IAccountService_IQueryUserByUsername;
    updateUserPasswordByID: IAccountService_IUpdateUserPasswordByID;
    verifyUserPasswordByID: IAccountService_IVerifyUserPasswordByID;
    verifyUserPasswordByUsername: IAccountService_IVerifyUserPasswordByUsername;
    queryLabelByID: IAccountService_IQueryLabelByID;
    updateLabelByID: IAccountService_IUpdateLabelByID;
    deleteLabelByID: IAccountService_IDeleteLabelByID;
    queryLabelByOwner: IAccountService_IQueryLabelByOwner;
    createLabelByOwner: IAccountService_ICreateLabelByOwner;
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
interface IAccountService_IUpdateUserByID extends grpc.MethodDefinition<standard_pb.UpdateUserByIDRequest, standard_pb.UpdateUserByIDResponse> {
    path: string; // "/standard.Account/UpdateUserByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateUserByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateUserByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateUserByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateUserByIDResponse>;
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
interface IAccountService_IQueryLabelByID extends grpc.MethodDefinition<standard_pb.QueryLabelByIDRequest, standard_pb.QueryLabelByIDResponse> {
    path: string; // "/standard.Account/QueryLabelByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryLabelByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryLabelByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryLabelByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryLabelByIDResponse>;
}
interface IAccountService_IUpdateLabelByID extends grpc.MethodDefinition<standard_pb.UpdateLabelByIDRequest, standard_pb.UpdateLabelByIDResponse> {
    path: string; // "/standard.Account/UpdateLabelByID"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.UpdateLabelByIDRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.UpdateLabelByIDRequest>;
    responseSerialize: grpc.serialize<standard_pb.UpdateLabelByIDResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.UpdateLabelByIDResponse>;
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
interface IAccountService_IQueryLabelByOwner extends grpc.MethodDefinition<standard_pb.QueryLabelByOwnerRequest, standard_pb.QueryLabelByOwnerResponse> {
    path: string; // "/standard.Account/QueryLabelByOwner"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.QueryLabelByOwnerRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.QueryLabelByOwnerRequest>;
    responseSerialize: grpc.serialize<standard_pb.QueryLabelByOwnerResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.QueryLabelByOwnerResponse>;
}
interface IAccountService_ICreateLabelByOwner extends grpc.MethodDefinition<standard_pb.CreateLabelByOwnerRequest, standard_pb.CreateLabelByOwnerResponse> {
    path: string; // "/standard.Account/CreateLabelByOwner"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<standard_pb.CreateLabelByOwnerRequest>;
    requestDeserialize: grpc.deserialize<standard_pb.CreateLabelByOwnerRequest>;
    responseSerialize: grpc.serialize<standard_pb.CreateLabelByOwnerResponse>;
    responseDeserialize: grpc.deserialize<standard_pb.CreateLabelByOwnerResponse>;
}

export const AccountService: IAccountService;

export interface IAccountServer {
    createUser: grpc.handleUnaryCall<standard_pb.CreateUserRequest, standard_pb.CreateUserResponse>;
    queryUserByID: grpc.handleUnaryCall<standard_pb.QueryUserByIDRequest, standard_pb.QueryUserByIDResponse>;
    updateUserByID: grpc.handleUnaryCall<standard_pb.UpdateUserByIDRequest, standard_pb.UpdateUserByIDResponse>;
    deleteUserByID: grpc.handleUnaryCall<standard_pb.DeleteUserByIDRequest, standard_pb.DeleteUserByIDResponse>;
    queryUserByUsername: grpc.handleUnaryCall<standard_pb.QueryUserByUsernameRequest, standard_pb.QueryUserByUsernameResponse>;
    updateUserPasswordByID: grpc.handleUnaryCall<standard_pb.UpdateUserPasswordByIDRequest, standard_pb.UpdateUserPasswordByIDResponse>;
    verifyUserPasswordByID: grpc.handleUnaryCall<standard_pb.VerifyUserPasswordByIDRequest, standard_pb.VerifyUserPasswordByIDResponse>;
    verifyUserPasswordByUsername: grpc.handleUnaryCall<standard_pb.VerifyUserPasswordByUsernameRequest, standard_pb.VerifyUserPasswordByUsernameResponse>;
    queryLabelByID: grpc.handleUnaryCall<standard_pb.QueryLabelByIDRequest, standard_pb.QueryLabelByIDResponse>;
    updateLabelByID: grpc.handleUnaryCall<standard_pb.UpdateLabelByIDRequest, standard_pb.UpdateLabelByIDResponse>;
    deleteLabelByID: grpc.handleUnaryCall<standard_pb.DeleteLabelByIDRequest, standard_pb.DeleteLabelByIDResponse>;
    queryLabelByOwner: grpc.handleUnaryCall<standard_pb.QueryLabelByOwnerRequest, standard_pb.QueryLabelByOwnerResponse>;
    createLabelByOwner: grpc.handleUnaryCall<standard_pb.CreateLabelByOwnerRequest, standard_pb.CreateLabelByOwnerResponse>;
}

export interface IAccountClient {
    createUser(request: standard_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    updateUserByID(request: standard_pb.UpdateUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
    updateUserByID(request: standard_pb.UpdateUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
    updateUserByID(request: standard_pb.UpdateUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
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
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
}

export class AccountClient extends grpc.Client implements IAccountClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createUser(request: standard_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: standard_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public queryUserByID(request: standard_pb.QueryUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryUserByIDResponse) => void): grpc.ClientUnaryCall;
    public updateUserByID(request: standard_pb.UpdateUserByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
    public updateUserByID(request: standard_pb.UpdateUserByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
    public updateUserByID(request: standard_pb.UpdateUserByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateUserByIDResponse) => void): grpc.ClientUnaryCall;
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
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByID(request: standard_pb.QueryLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public updateLabelByID(request: standard_pb.UpdateLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.UpdateLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public deleteLabelByID(request: standard_pb.DeleteLabelByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.DeleteLabelByIDResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    public queryLabelByOwner(request: standard_pb.QueryLabelByOwnerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.QueryLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    public createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    public createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
    public createLabelByOwner(request: standard_pb.CreateLabelByOwnerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: standard_pb.CreateLabelByOwnerResponse) => void): grpc.ClientUnaryCall;
}
