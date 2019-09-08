// package: standard
// file: standard.proto

/* tslint:disable */

import * as jspb from "google-protobuf";

export class Label extends jspb.Message { 
    getId(): number;
    setId(value: number): void;

    getType(): string;
    setType(value: string): void;

    getState(): string;
    setState(value: string): void;

    getValue(): string;
    setValue(value: string): void;

    getOwner(): number;
    setOwner(value: number): void;

    getCreatetime(): string;
    setCreatetime(value: string): void;

    getUpdatetime(): string;
    setUpdatetime(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Label.AsObject;
    static toObject(includeInstance: boolean, msg: Label): Label.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Label, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Label;
    static deserializeBinaryFromReader(message: Label, reader: jspb.BinaryReader): Label;
}

export namespace Label {
    export type AsObject = {
        id: number,
        type: string,
        state: string,
        value: string,
        owner: number,
        createtime: string,
        updatetime: string,
    }
}

export class User extends jspb.Message { 
    getId(): number;
    setId(value: number): void;

    getType(): string;
    setType(value: string): void;

    getAvatar(): string;
    setAvatar(value: string): void;

    getInviter(): number;
    setInviter(value: number): void;

    getNickname(): string;
    setNickname(value: string): void;

    getUsername(): string;
    setUsername(value: string): void;

    getPassword(): string;
    setPassword(value: string): void;

    getCreatetime(): string;
    setCreatetime(value: string): void;

    getUpdatetime(): string;
    setUpdatetime(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): User.AsObject;
    static toObject(includeInstance: boolean, msg: User): User.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): User;
    static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
    export type AsObject = {
        id: number,
        type: string,
        avatar: string,
        inviter: number,
        nickname: string,
        username: string,
        password: string,
        createtime: string,
        updatetime: string,
    }
}

export class Group extends jspb.Message { 
    getId(): number;
    setId(value: number): void;

    getType(): string;
    setType(value: string): void;

    getName(): string;
    setName(value: string): void;

    getCreatetime(): string;
    setCreatetime(value: string): void;

    getUpdatetime(): string;
    setUpdatetime(value: string): void;

    getDescription(): string;
    setDescription(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Group.AsObject;
    static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Group;
    static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
    export type AsObject = {
        id: number,
        type: string,
        name: string,
        createtime: string,
        updatetime: string,
        description: string,
    }
}

export class CreateUserRequest extends jspb.Message { 
    getType(): string;
    setType(value: string): void;

    getAvatar(): string;
    setAvatar(value: string): void;

    getInviter(): number;
    setInviter(value: number): void;

    getNickname(): string;
    setNickname(value: string): void;

    getUsername(): string;
    setUsername(value: string): void;

    getPassword(): string;
    setPassword(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
    static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
    export type AsObject = {
        type: string,
        avatar: string,
        inviter: number,
        nickname: string,
        username: string,
        password: string,
    }
}

export class CreateUserResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
    static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class QueryUserByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryUserByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: QueryUserByIDRequest): QueryUserByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryUserByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryUserByIDRequest;
    static deserializeBinaryFromReader(message: QueryUserByIDRequest, reader: jspb.BinaryReader): QueryUserByIDRequest;
}

export namespace QueryUserByIDRequest {
    export type AsObject = {
        id: number,
    }
}

export class QueryUserByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    hasData(): boolean;
    clearData(): void;
    getData(): User | undefined;
    setData(value?: User): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryUserByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QueryUserByIDResponse): QueryUserByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryUserByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryUserByIDResponse;
    static deserializeBinaryFromReader(message: QueryUserByIDResponse, reader: jspb.BinaryReader): QueryUserByIDResponse;
}

export namespace QueryUserByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
        data?: User.AsObject,
    }
}

export class UpdateUserByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    hasData(): boolean;
    clearData(): void;
    getData(): User | undefined;
    setData(value?: User): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserByIDRequest): UpdateUserByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserByIDRequest;
    static deserializeBinaryFromReader(message: UpdateUserByIDRequest, reader: jspb.BinaryReader): UpdateUserByIDRequest;
}

export namespace UpdateUserByIDRequest {
    export type AsObject = {
        id: number,
        data?: User.AsObject,
    }
}

export class QueryUserByUsernameRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryUserByUsernameRequest.AsObject;
    static toObject(includeInstance: boolean, msg: QueryUserByUsernameRequest): QueryUserByUsernameRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryUserByUsernameRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryUserByUsernameRequest;
    static deserializeBinaryFromReader(message: QueryUserByUsernameRequest, reader: jspb.BinaryReader): QueryUserByUsernameRequest;
}

export namespace QueryUserByUsernameRequest {
    export type AsObject = {
        username: string,
    }
}

export class QueryUserByUsernameResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    hasData(): boolean;
    clearData(): void;
    getData(): User | undefined;
    setData(value?: User): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryUserByUsernameResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QueryUserByUsernameResponse): QueryUserByUsernameResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryUserByUsernameResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryUserByUsernameResponse;
    static deserializeBinaryFromReader(message: QueryUserByUsernameResponse, reader: jspb.BinaryReader): QueryUserByUsernameResponse;
}

export namespace QueryUserByUsernameResponse {
    export type AsObject = {
        state: number,
        message: string,
        data?: User.AsObject,
    }
}

export class UpdateUserByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserByIDResponse): UpdateUserByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserByIDResponse;
    static deserializeBinaryFromReader(message: UpdateUserByIDResponse, reader: jspb.BinaryReader): UpdateUserByIDResponse;
}

export namespace UpdateUserByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class DeleteUserByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteUserByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteUserByIDRequest): DeleteUserByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteUserByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteUserByIDRequest;
    static deserializeBinaryFromReader(message: DeleteUserByIDRequest, reader: jspb.BinaryReader): DeleteUserByIDRequest;
}

export namespace DeleteUserByIDRequest {
    export type AsObject = {
        id: number,
    }
}

export class DeleteUserByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteUserByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteUserByIDResponse): DeleteUserByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteUserByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteUserByIDResponse;
    static deserializeBinaryFromReader(message: DeleteUserByIDResponse, reader: jspb.BinaryReader): DeleteUserByIDResponse;
}

export namespace DeleteUserByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class UpdateUserPasswordByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;

    getPassword(): string;
    setPassword(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserPasswordByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserPasswordByIDRequest): UpdateUserPasswordByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserPasswordByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserPasswordByIDRequest;
    static deserializeBinaryFromReader(message: UpdateUserPasswordByIDRequest, reader: jspb.BinaryReader): UpdateUserPasswordByIDRequest;
}

export namespace UpdateUserPasswordByIDRequest {
    export type AsObject = {
        id: number,
        password: string,
    }
}

export class UpdateUserPasswordByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserPasswordByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserPasswordByIDResponse): UpdateUserPasswordByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserPasswordByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserPasswordByIDResponse;
    static deserializeBinaryFromReader(message: UpdateUserPasswordByIDResponse, reader: jspb.BinaryReader): UpdateUserPasswordByIDResponse;
}

export namespace UpdateUserPasswordByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class VerifyUserPasswordByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;

    getPassword(): string;
    setPassword(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): VerifyUserPasswordByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: VerifyUserPasswordByIDRequest): VerifyUserPasswordByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: VerifyUserPasswordByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): VerifyUserPasswordByIDRequest;
    static deserializeBinaryFromReader(message: VerifyUserPasswordByIDRequest, reader: jspb.BinaryReader): VerifyUserPasswordByIDRequest;
}

export namespace VerifyUserPasswordByIDRequest {
    export type AsObject = {
        id: number,
        password: string,
    }
}

export class VerifyUserPasswordByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;

    getData(): boolean;
    setData(value: boolean): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): VerifyUserPasswordByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: VerifyUserPasswordByIDResponse): VerifyUserPasswordByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: VerifyUserPasswordByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): VerifyUserPasswordByIDResponse;
    static deserializeBinaryFromReader(message: VerifyUserPasswordByIDResponse, reader: jspb.BinaryReader): VerifyUserPasswordByIDResponse;
}

export namespace VerifyUserPasswordByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
        data: boolean,
    }
}

export class VerifyUserPasswordByUsernameRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): void;

    getPassword(): string;
    setPassword(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): VerifyUserPasswordByUsernameRequest.AsObject;
    static toObject(includeInstance: boolean, msg: VerifyUserPasswordByUsernameRequest): VerifyUserPasswordByUsernameRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: VerifyUserPasswordByUsernameRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): VerifyUserPasswordByUsernameRequest;
    static deserializeBinaryFromReader(message: VerifyUserPasswordByUsernameRequest, reader: jspb.BinaryReader): VerifyUserPasswordByUsernameRequest;
}

export namespace VerifyUserPasswordByUsernameRequest {
    export type AsObject = {
        username: string,
        password: string,
    }
}

export class VerifyUserPasswordByUsernameResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;

    getData(): boolean;
    setData(value: boolean): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): VerifyUserPasswordByUsernameResponse.AsObject;
    static toObject(includeInstance: boolean, msg: VerifyUserPasswordByUsernameResponse): VerifyUserPasswordByUsernameResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: VerifyUserPasswordByUsernameResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): VerifyUserPasswordByUsernameResponse;
    static deserializeBinaryFromReader(message: VerifyUserPasswordByUsernameResponse, reader: jspb.BinaryReader): VerifyUserPasswordByUsernameResponse;
}

export namespace VerifyUserPasswordByUsernameResponse {
    export type AsObject = {
        state: number,
        message: string,
        data: boolean,
    }
}

export class CreateLabelByOwnerRequest extends jspb.Message { 
    getOwner(): number;
    setOwner(value: number): void;


    hasLabel(): boolean;
    clearLabel(): void;
    getLabel(): Label | undefined;
    setLabel(value?: Label): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateLabelByOwnerRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateLabelByOwnerRequest): CreateLabelByOwnerRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateLabelByOwnerRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateLabelByOwnerRequest;
    static deserializeBinaryFromReader(message: CreateLabelByOwnerRequest, reader: jspb.BinaryReader): CreateLabelByOwnerRequest;
}

export namespace CreateLabelByOwnerRequest {
    export type AsObject = {
        owner: number,
        label?: Label.AsObject,
    }
}

export class CreateLabelByOwnerResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateLabelByOwnerResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateLabelByOwnerResponse): CreateLabelByOwnerResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateLabelByOwnerResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateLabelByOwnerResponse;
    static deserializeBinaryFromReader(message: CreateLabelByOwnerResponse, reader: jspb.BinaryReader): CreateLabelByOwnerResponse;
}

export namespace CreateLabelByOwnerResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class QueryLabelByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryLabelByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: QueryLabelByIDRequest): QueryLabelByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryLabelByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryLabelByIDRequest;
    static deserializeBinaryFromReader(message: QueryLabelByIDRequest, reader: jspb.BinaryReader): QueryLabelByIDRequest;
}

export namespace QueryLabelByIDRequest {
    export type AsObject = {
        id: number,
    }
}

export class QueryLabelByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    hasData(): boolean;
    clearData(): void;
    getData(): Label | undefined;
    setData(value?: Label): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryLabelByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QueryLabelByIDResponse): QueryLabelByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryLabelByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryLabelByIDResponse;
    static deserializeBinaryFromReader(message: QueryLabelByIDResponse, reader: jspb.BinaryReader): QueryLabelByIDResponse;
}

export namespace QueryLabelByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
        data?: Label.AsObject,
    }
}

export class UpdateLabelByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    hasData(): boolean;
    clearData(): void;
    getData(): Label | undefined;
    setData(value?: Label): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateLabelByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateLabelByIDRequest): UpdateLabelByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateLabelByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateLabelByIDRequest;
    static deserializeBinaryFromReader(message: UpdateLabelByIDRequest, reader: jspb.BinaryReader): UpdateLabelByIDRequest;
}

export namespace UpdateLabelByIDRequest {
    export type AsObject = {
        id: number,
        data?: Label.AsObject,
    }
}

export class UpdateLabelByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateLabelByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateLabelByIDResponse): UpdateLabelByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateLabelByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateLabelByIDResponse;
    static deserializeBinaryFromReader(message: UpdateLabelByIDResponse, reader: jspb.BinaryReader): UpdateLabelByIDResponse;
}

export namespace UpdateLabelByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class DeleteLabelByIDRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteLabelByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteLabelByIDRequest): DeleteLabelByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteLabelByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteLabelByIDRequest;
    static deserializeBinaryFromReader(message: DeleteLabelByIDRequest, reader: jspb.BinaryReader): DeleteLabelByIDRequest;
}

export namespace DeleteLabelByIDRequest {
    export type AsObject = {
        id: number,
    }
}

export class DeleteLabelByIDResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteLabelByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteLabelByIDResponse): DeleteLabelByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteLabelByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteLabelByIDResponse;
    static deserializeBinaryFromReader(message: DeleteLabelByIDResponse, reader: jspb.BinaryReader): DeleteLabelByIDResponse;
}

export namespace DeleteLabelByIDResponse {
    export type AsObject = {
        state: number,
        message: string,
    }
}

export class QueryLabelByOwnerRequest extends jspb.Message { 
    getOwner(): number;
    setOwner(value: number): void;

    getLimit(): number;
    setLimit(value: number): void;

    getOffset(): number;
    setOffset(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryLabelByOwnerRequest.AsObject;
    static toObject(includeInstance: boolean, msg: QueryLabelByOwnerRequest): QueryLabelByOwnerRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryLabelByOwnerRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryLabelByOwnerRequest;
    static deserializeBinaryFromReader(message: QueryLabelByOwnerRequest, reader: jspb.BinaryReader): QueryLabelByOwnerRequest;
}

export namespace QueryLabelByOwnerRequest {
    export type AsObject = {
        owner: number,
        limit: number,
        offset: number,
    }
}

export class QueryLabelByOwnerResponse extends jspb.Message { 
    getState(): number;
    setState(value: number): void;

    getMessage(): string;
    setMessage(value: string): void;

    getTotal(): number;
    setTotal(value: number): void;

    clearDataList(): void;
    getDataList(): Array<Label>;
    setDataList(value: Array<Label>): void;
    addData(value?: Label, index?: number): Label;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryLabelByOwnerResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QueryLabelByOwnerResponse): QueryLabelByOwnerResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryLabelByOwnerResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryLabelByOwnerResponse;
    static deserializeBinaryFromReader(message: QueryLabelByOwnerResponse, reader: jspb.BinaryReader): QueryLabelByOwnerResponse;
}

export namespace QueryLabelByOwnerResponse {
    export type AsObject = {
        state: number,
        message: string,
        total: number,
        dataList: Array<Label.AsObject>,
    }
}
