# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [standard.proto](#standard.proto)
    - [CreateLabelByOwnerRequest](#standard.CreateLabelByOwnerRequest)
    - [CreateLabelByOwnerResponse](#standard.CreateLabelByOwnerResponse)
    - [CreateUserRequest](#standard.CreateUserRequest)
    - [CreateUserResponse](#standard.CreateUserResponse)
    - [DeleteLabelByIDRequest](#standard.DeleteLabelByIDRequest)
    - [DeleteLabelByIDResponse](#standard.DeleteLabelByIDResponse)
    - [DeleteUserByIDRequest](#standard.DeleteUserByIDRequest)
    - [DeleteUserByIDResponse](#standard.DeleteUserByIDResponse)
    - [Label](#standard.Label)
    - [QueryLabelByIDRequest](#standard.QueryLabelByIDRequest)
    - [QueryLabelByIDResponse](#standard.QueryLabelByIDResponse)
    - [QueryLabelByOwnerRequest](#standard.QueryLabelByOwnerRequest)
    - [QueryLabelByOwnerResponse](#standard.QueryLabelByOwnerResponse)
    - [QueryUserByIDRequest](#standard.QueryUserByIDRequest)
    - [QueryUserByIDResponse](#standard.QueryUserByIDResponse)
    - [QueryUserByUsernameRequest](#standard.QueryUserByUsernameRequest)
    - [QueryUserByUsernameResponse](#standard.QueryUserByUsernameResponse)
    - [UpdateLabelByIDRequest](#standard.UpdateLabelByIDRequest)
    - [UpdateLabelByIDResponse](#standard.UpdateLabelByIDResponse)
    - [UpdateUserByIDRequest](#standard.UpdateUserByIDRequest)
    - [UpdateUserByIDResponse](#standard.UpdateUserByIDResponse)
    - [UpdateUserPasswordByIDRequest](#standard.UpdateUserPasswordByIDRequest)
    - [UpdateUserPasswordByIDResponse](#standard.UpdateUserPasswordByIDResponse)
    - [User](#standard.User)
    - [VerifyUserPasswordByIDRequest](#standard.VerifyUserPasswordByIDRequest)
    - [VerifyUserPasswordByIDResponse](#standard.VerifyUserPasswordByIDResponse)
    - [VerifyUserPasswordByUsernameRequest](#standard.VerifyUserPasswordByUsernameRequest)
    - [VerifyUserPasswordByUsernameResponse](#standard.VerifyUserPasswordByUsernameResponse)
  
    - [State](#standard.State)
  
  
    - [Account](#standard.Account)
  

- [Scalar Value Types](#scalar-value-types)



<a name="standard.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## standard.proto



<a name="standard.CreateLabelByOwnerRequest"></a>

### CreateLabelByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Owner | [uint64](#uint64) |  |  |
| Label | [Label](#standard.Label) |  |  |






<a name="standard.CreateLabelByOwnerResponse"></a>

### CreateLabelByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Type | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Inviter | [uint64](#uint64) |  |  |
| Nickname | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.DeleteLabelByIDRequest"></a>

### DeleteLabelByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.DeleteLabelByIDResponse"></a>

### DeleteLabelByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.DeleteUserByIDRequest"></a>

### DeleteUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.DeleteUserByIDResponse"></a>

### DeleteUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.Label"></a>

### Label
Label 标签


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Type | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Value | [string](#string) |  |  |
| Owner | [uint64](#uint64) |  |  |
| CreateTime | [string](#string) |  |  |
| UpdateTime | [string](#string) |  |  |






<a name="standard.QueryLabelByIDRequest"></a>

### QueryLabelByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.QueryLabelByIDResponse"></a>

### QueryLabelByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [Label](#standard.Label) |  |  |






<a name="standard.QueryLabelByOwnerRequest"></a>

### QueryLabelByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Owner | [uint64](#uint64) |  |  |
| Limit | [uint64](#uint64) |  |  |
| Offset | [uint64](#uint64) |  |  |






<a name="standard.QueryLabelByOwnerResponse"></a>

### QueryLabelByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Total | [uint64](#uint64) |  |  |
| Data | [Label](#standard.Label) | repeated |  |






<a name="standard.QueryUserByIDRequest"></a>

### QueryUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.QueryUserByIDResponse"></a>

### QueryUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [User](#standard.User) |  |  |






<a name="standard.QueryUserByUsernameRequest"></a>

### QueryUserByUsernameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |






<a name="standard.QueryUserByUsernameResponse"></a>

### QueryUserByUsernameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [User](#standard.User) |  |  |






<a name="standard.UpdateLabelByIDRequest"></a>

### UpdateLabelByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Data | [Label](#standard.Label) |  |  |






<a name="standard.UpdateLabelByIDResponse"></a>

### UpdateLabelByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserByIDRequest"></a>

### UpdateUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Data | [User](#standard.User) |  |  |






<a name="standard.UpdateUserByIDResponse"></a>

### UpdateUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserPasswordByIDRequest"></a>

### UpdateUserPasswordByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.UpdateUserPasswordByIDResponse"></a>

### UpdateUserPasswordByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.User"></a>

### User
User 用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Type | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Inviter | [uint64](#uint64) |  |  |
| Nickname | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| CreateTime | [string](#string) |  |  |
| UpdateTime | [string](#string) |  |  |






<a name="standard.VerifyUserPasswordByIDRequest"></a>

### VerifyUserPasswordByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.VerifyUserPasswordByIDResponse"></a>

### VerifyUserPasswordByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [bool](#bool) |  |  |






<a name="standard.VerifyUserPasswordByUsernameRequest"></a>

### VerifyUserPasswordByUsernameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.VerifyUserPasswordByUsernameResponse"></a>

### VerifyUserPasswordByUsernameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [bool](#bool) |  |  |





 


<a name="standard.State"></a>

### State
状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知 |
| SUCCESS | 1 | 成功 |
| FAILURE | 2 | 失败 |
| SERVICE_ERROR | 3 | 服务错误 |
| PARAMS_INVALID | 5 | 参数不合法 |
| ILLEGAL_REQUEST | 6 | 非法请求 |
| LABEL_NOT_EXIST | 7 | 标签不存在 |
| USER_NOT_EXIST | 8 | 用户不存在 |
| USER_ALREADY_EXISTS | 9 | 用户已经存在 |
| USER_VERIFY_FAILURE | 10 | 用户验证失败 |
| LABEL_ALREADY_EXISTS | 11 | 标签已经存在 |
| DB_OPERATION_FATLURE | 12 | 数据库操作失败 |


 

 


<a name="standard.Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#standard.CreateUserRequest) | [CreateUserResponse](#standard.CreateUserResponse) | 用户操作 |
| QueryUserByID | [QueryUserByIDRequest](#standard.QueryUserByIDRequest) | [QueryUserByIDResponse](#standard.QueryUserByIDResponse) |  |
| UpdateUserByID | [UpdateUserByIDRequest](#standard.UpdateUserByIDRequest) | [UpdateUserByIDResponse](#standard.UpdateUserByIDResponse) |  |
| DeleteUserByID | [DeleteUserByIDRequest](#standard.DeleteUserByIDRequest) | [DeleteUserByIDResponse](#standard.DeleteUserByIDResponse) |  |
| QueryUserByUsername | [QueryUserByUsernameRequest](#standard.QueryUserByUsernameRequest) | [QueryUserByUsernameResponse](#standard.QueryUserByUsernameResponse) |  |
| UpdateUserPasswordByID | [UpdateUserPasswordByIDRequest](#standard.UpdateUserPasswordByIDRequest) | [UpdateUserPasswordByIDResponse](#standard.UpdateUserPasswordByIDResponse) |  |
| VerifyUserPasswordByID | [VerifyUserPasswordByIDRequest](#standard.VerifyUserPasswordByIDRequest) | [VerifyUserPasswordByIDResponse](#standard.VerifyUserPasswordByIDResponse) |  |
| VerifyUserPasswordByUsername | [VerifyUserPasswordByUsernameRequest](#standard.VerifyUserPasswordByUsernameRequest) | [VerifyUserPasswordByUsernameResponse](#standard.VerifyUserPasswordByUsernameResponse) |  |
| QueryLabelByID | [QueryLabelByIDRequest](#standard.QueryLabelByIDRequest) | [QueryLabelByIDResponse](#standard.QueryLabelByIDResponse) | 标签操作 |
| UpdateLabelByID | [UpdateLabelByIDRequest](#standard.UpdateLabelByIDRequest) | [UpdateLabelByIDResponse](#standard.UpdateLabelByIDResponse) |  |
| DeleteLabelByID | [DeleteLabelByIDRequest](#standard.DeleteLabelByIDRequest) | [DeleteLabelByIDResponse](#standard.DeleteLabelByIDResponse) |  |
| QueryLabelByOwner | [QueryLabelByOwnerRequest](#standard.QueryLabelByOwnerRequest) | [QueryLabelByOwnerResponse](#standard.QueryLabelByOwnerResponse) |  |
| CreateLabelByOwner | [CreateLabelByOwnerRequest](#standard.CreateLabelByOwnerRequest) | [CreateLabelByOwnerResponse](#standard.CreateLabelByOwnerResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
