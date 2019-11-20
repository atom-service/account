// package: standard
// file: standard.proto

import * as jspb from "google-protobuf";

export class Group extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getClass(): string;
  setClass(value: string): void;

  getState(): string;
  setState(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getCreatedtime(): string;
  setCreatedtime(value: string): void;

  getUpdatedtime(): string;
  setUpdatedtime(value: string): void;

  getDeletedtime(): string;
  setDeletedtime(value: string): void;

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
    name: string,
    pb_class: string,
    state: string,
    description: string,
    createdtime: string,
    updatedtime: string,
    deletedtime: string,
  }
}

export class Label extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getClass(): string;
  setClass(value: string): void;

  getState(): string;
  setState(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getCreatedtime(): string;
  setCreatedtime(value: string): void;

  getUpdatedtime(): string;
  setUpdatedtime(value: string): void;

  getDeletedtime(): string;
  setDeletedtime(value: string): void;

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
    pb_class: string,
    state: string,
    value: string,
    createdtime: string,
    updatedtime: string,
    deletedtime: string,
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getClass(): string;
  setClass(value: string): void;

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

  getCreatedtime(): string;
  setCreatedtime(value: string): void;

  getUpdatedtime(): string;
  setUpdatedtime(value: string): void;

  getDeletedtime(): string;
  setDeletedtime(value: string): void;

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
    pb_class: string,
    avatar: string,
    inviter: number,
    nickname: string,
    username: string,
    password: string,
    createdtime: string,
    updatedtime: string,
    deletedtime: string,
  }
}

export class CreateUserRequest extends jspb.Message {
  getClass(): string;
  setClass(value: string): void;

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
    pb_class: string,
    inviter: number,
    nickname: string,
    username: string,
    password: string,
  }
}

export class CreateUserResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
    message: string,
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
    message: string,
    data?: User.AsObject,
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

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
    state: StateMap[keyof StateMap],
    message: string,
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

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
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class CreateLabelRequest extends jspb.Message {
  getClass(): string;
  setClass(value: string): void;

  getState(): string;
  setState(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateLabelRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateLabelRequest): CreateLabelRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateLabelRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateLabelRequest;
  static deserializeBinaryFromReader(message: CreateLabelRequest, reader: jspb.BinaryReader): CreateLabelRequest;
}

export namespace CreateLabelRequest {
  export type AsObject = {
    pb_class: string,
    state: string,
    value: string,
  }
}

export class CreateLabelResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateLabelResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateLabelResponse): CreateLabelResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateLabelResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateLabelResponse;
  static deserializeBinaryFromReader(message: CreateLabelResponse, reader: jspb.BinaryReader): CreateLabelResponse;
}

export namespace CreateLabelResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
    message: string,
    data?: Label.AsObject,
  }
}

export class UpdateLabelClassByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getClass(): string;
  setClass(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelClassByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelClassByIDRequest): UpdateLabelClassByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelClassByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelClassByIDRequest;
  static deserializeBinaryFromReader(message: UpdateLabelClassByIDRequest, reader: jspb.BinaryReader): UpdateLabelClassByIDRequest;
}

export namespace UpdateLabelClassByIDRequest {
  export type AsObject = {
    id: number,
    pb_class: string,
  }
}

export class UpdateLabelClassByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelClassByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelClassByIDResponse): UpdateLabelClassByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelClassByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelClassByIDResponse;
  static deserializeBinaryFromReader(message: UpdateLabelClassByIDResponse, reader: jspb.BinaryReader): UpdateLabelClassByIDResponse;
}

export namespace UpdateLabelClassByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateLabelStateByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getState(): string;
  setState(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelStateByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelStateByIDRequest): UpdateLabelStateByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelStateByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelStateByIDRequest;
  static deserializeBinaryFromReader(message: UpdateLabelStateByIDRequest, reader: jspb.BinaryReader): UpdateLabelStateByIDRequest;
}

export namespace UpdateLabelStateByIDRequest {
  export type AsObject = {
    id: number,
    state: string,
  }
}

export class UpdateLabelStateByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelStateByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelStateByIDResponse): UpdateLabelStateByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelStateByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelStateByIDResponse;
  static deserializeBinaryFromReader(message: UpdateLabelStateByIDResponse, reader: jspb.BinaryReader): UpdateLabelStateByIDResponse;
}

export namespace UpdateLabelStateByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateLabelValueByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelValueByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelValueByIDRequest): UpdateLabelValueByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelValueByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelValueByIDRequest;
  static deserializeBinaryFromReader(message: UpdateLabelValueByIDRequest, reader: jspb.BinaryReader): UpdateLabelValueByIDRequest;
}

export namespace UpdateLabelValueByIDRequest {
  export type AsObject = {
    id: number,
    value: string,
  }
}

export class UpdateLabelValueByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateLabelValueByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateLabelValueByIDResponse): UpdateLabelValueByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateLabelValueByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateLabelValueByIDResponse;
  static deserializeBinaryFromReader(message: UpdateLabelValueByIDResponse, reader: jspb.BinaryReader): UpdateLabelValueByIDResponse;
}

export namespace UpdateLabelValueByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
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
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

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
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class AddLabelToUserByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getUserid(): number;
  setUserid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddLabelToUserByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddLabelToUserByIDRequest): AddLabelToUserByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddLabelToUserByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddLabelToUserByIDRequest;
  static deserializeBinaryFromReader(message: AddLabelToUserByIDRequest, reader: jspb.BinaryReader): AddLabelToUserByIDRequest;
}

export namespace AddLabelToUserByIDRequest {
  export type AsObject = {
    id: number,
    userid: number,
  }
}

export class AddLabelToUserByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddLabelToUserByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddLabelToUserByIDResponse): AddLabelToUserByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddLabelToUserByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddLabelToUserByIDResponse;
  static deserializeBinaryFromReader(message: AddLabelToUserByIDResponse, reader: jspb.BinaryReader): AddLabelToUserByIDResponse;
}

export namespace AddLabelToUserByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class RemoveLabelFromUserByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getLabelid(): number;
  setLabelid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveLabelFromUserByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveLabelFromUserByIDRequest): RemoveLabelFromUserByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoveLabelFromUserByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveLabelFromUserByIDRequest;
  static deserializeBinaryFromReader(message: RemoveLabelFromUserByIDRequest, reader: jspb.BinaryReader): RemoveLabelFromUserByIDRequest;
}

export namespace RemoveLabelFromUserByIDRequest {
  export type AsObject = {
    id: number,
    labelid: number,
  }
}

export class RemoveLabelFromUserByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveLabelFromUserByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveLabelFromUserByIDResponse): RemoveLabelFromUserByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoveLabelFromUserByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveLabelFromUserByIDResponse;
  static deserializeBinaryFromReader(message: RemoveLabelFromUserByIDResponse, reader: jspb.BinaryReader): RemoveLabelFromUserByIDResponse;
}

export namespace RemoveLabelFromUserByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class CreateGroupRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getClass(): string;
  setClass(value: string): void;

  getState(): string;
  setState(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateGroupRequest): CreateGroupRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateGroupRequest;
  static deserializeBinaryFromReader(message: CreateGroupRequest, reader: jspb.BinaryReader): CreateGroupRequest;
}

export namespace CreateGroupRequest {
  export type AsObject = {
    name: string,
    pb_class: string,
    state: string,
    description: string,
  }
}

export class CreateGroupResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateGroupResponse): CreateGroupResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateGroupResponse;
  static deserializeBinaryFromReader(message: CreateGroupResponse, reader: jspb.BinaryReader): CreateGroupResponse;
}

export namespace CreateGroupResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class QueryGroupByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryGroupByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: QueryGroupByIDRequest): QueryGroupByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryGroupByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryGroupByIDRequest;
  static deserializeBinaryFromReader(message: QueryGroupByIDRequest, reader: jspb.BinaryReader): QueryGroupByIDRequest;
}

export namespace QueryGroupByIDRequest {
  export type AsObject = {
    id: number,
  }
}

export class QueryGroupByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Group | undefined;
  setData(value?: Group): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryGroupByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: QueryGroupByIDResponse): QueryGroupByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryGroupByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryGroupByIDResponse;
  static deserializeBinaryFromReader(message: QueryGroupByIDResponse, reader: jspb.BinaryReader): QueryGroupByIDResponse;
}

export namespace QueryGroupByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
    data?: Group.AsObject,
  }
}

export class UpdateGroupNameByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupNameByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupNameByIDRequest): UpdateGroupNameByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupNameByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupNameByIDRequest;
  static deserializeBinaryFromReader(message: UpdateGroupNameByIDRequest, reader: jspb.BinaryReader): UpdateGroupNameByIDRequest;
}

export namespace UpdateGroupNameByIDRequest {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class UpdateGroupNameByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupNameByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupNameByIDResponse): UpdateGroupNameByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupNameByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupNameByIDResponse;
  static deserializeBinaryFromReader(message: UpdateGroupNameByIDResponse, reader: jspb.BinaryReader): UpdateGroupNameByIDResponse;
}

export namespace UpdateGroupNameByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateGroupClassByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getClass(): string;
  setClass(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupClassByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupClassByIDRequest): UpdateGroupClassByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupClassByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupClassByIDRequest;
  static deserializeBinaryFromReader(message: UpdateGroupClassByIDRequest, reader: jspb.BinaryReader): UpdateGroupClassByIDRequest;
}

export namespace UpdateGroupClassByIDRequest {
  export type AsObject = {
    id: number,
    pb_class: string,
  }
}

export class UpdateGroupClassByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupClassByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupClassByIDResponse): UpdateGroupClassByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupClassByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupClassByIDResponse;
  static deserializeBinaryFromReader(message: UpdateGroupClassByIDResponse, reader: jspb.BinaryReader): UpdateGroupClassByIDResponse;
}

export namespace UpdateGroupClassByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateGroupStateByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getState(): string;
  setState(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupStateByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupStateByIDRequest): UpdateGroupStateByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupStateByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupStateByIDRequest;
  static deserializeBinaryFromReader(message: UpdateGroupStateByIDRequest, reader: jspb.BinaryReader): UpdateGroupStateByIDRequest;
}

export namespace UpdateGroupStateByIDRequest {
  export type AsObject = {
    id: number,
    state: string,
  }
}

export class UpdateGroupStateByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupStateByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupStateByIDResponse): UpdateGroupStateByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupStateByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupStateByIDResponse;
  static deserializeBinaryFromReader(message: UpdateGroupStateByIDResponse, reader: jspb.BinaryReader): UpdateGroupStateByIDResponse;
}

export namespace UpdateGroupStateByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateGroupDescriptionByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupDescriptionByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupDescriptionByIDRequest): UpdateGroupDescriptionByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupDescriptionByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupDescriptionByIDRequest;
  static deserializeBinaryFromReader(message: UpdateGroupDescriptionByIDRequest, reader: jspb.BinaryReader): UpdateGroupDescriptionByIDRequest;
}

export namespace UpdateGroupDescriptionByIDRequest {
  export type AsObject = {
    id: number,
    description: string,
  }
}

export class UpdateGroupDescriptionByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupDescriptionByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupDescriptionByIDResponse): UpdateGroupDescriptionByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateGroupDescriptionByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupDescriptionByIDResponse;
  static deserializeBinaryFromReader(message: UpdateGroupDescriptionByIDResponse, reader: jspb.BinaryReader): UpdateGroupDescriptionByIDResponse;
}

export namespace UpdateGroupDescriptionByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class DeleteGroupByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteGroupByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteGroupByIDRequest): DeleteGroupByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteGroupByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteGroupByIDRequest;
  static deserializeBinaryFromReader(message: DeleteGroupByIDRequest, reader: jspb.BinaryReader): DeleteGroupByIDRequest;
}

export namespace DeleteGroupByIDRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteGroupByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteGroupByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteGroupByIDResponse): DeleteGroupByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteGroupByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteGroupByIDResponse;
  static deserializeBinaryFromReader(message: DeleteGroupByIDResponse, reader: jspb.BinaryReader): DeleteGroupByIDResponse;
}

export namespace DeleteGroupByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class AddUserToGroupByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getGroupid(): number;
  setGroupid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUserToGroupByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddUserToGroupByIDRequest): AddUserToGroupByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddUserToGroupByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUserToGroupByIDRequest;
  static deserializeBinaryFromReader(message: AddUserToGroupByIDRequest, reader: jspb.BinaryReader): AddUserToGroupByIDRequest;
}

export namespace AddUserToGroupByIDRequest {
  export type AsObject = {
    id: number,
    groupid: number,
  }
}

export class AddUserToGroupByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUserToGroupByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddUserToGroupByIDResponse): AddUserToGroupByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddUserToGroupByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUserToGroupByIDResponse;
  static deserializeBinaryFromReader(message: AddUserToGroupByIDResponse, reader: jspb.BinaryReader): AddUserToGroupByIDResponse;
}

export namespace AddUserToGroupByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class RemoveUserFromGroupByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getUserid(): number;
  setUserid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveUserFromGroupByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveUserFromGroupByIDRequest): RemoveUserFromGroupByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoveUserFromGroupByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveUserFromGroupByIDRequest;
  static deserializeBinaryFromReader(message: RemoveUserFromGroupByIDRequest, reader: jspb.BinaryReader): RemoveUserFromGroupByIDRequest;
}

export namespace RemoveUserFromGroupByIDRequest {
  export type AsObject = {
    id: number,
    userid: number,
  }
}

export class RemoveUserFromGroupByIDResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveUserFromGroupByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveUserFromGroupByIDResponse): RemoveUserFromGroupByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoveUserFromGroupByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveUserFromGroupByIDResponse;
  static deserializeBinaryFromReader(message: RemoveUserFromGroupByIDResponse, reader: jspb.BinaryReader): RemoveUserFromGroupByIDResponse;
}

export namespace RemoveUserFromGroupByIDResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export interface StateMap {
  UNKNOWN: 0;
  SUCCESS: 1;
  FAILURE: 2;
  SERVICE_ERROR: 3;
  PARAMS_INVALID: 4;
  ILLEGAL_REQUEST: 5;
  LABEL_NOT_EXIST: 6;
  GROUP_NOT_EXIST: 7;
  USER_NOT_EXIST: 8;
  USER_ALREADY_EXISTS: 9;
  USER_VERIFY_FAILURE: 10;
  LABEL_ALREADY_EXISTS: 11;
  DB_OPERATION_FATLURE: 12;
}

export const State: StateMap;

