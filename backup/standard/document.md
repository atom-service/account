# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [standard.proto](#standard.proto)
    - [AddLabelToUserByIDRequest](#standard.AddLabelToUserByIDRequest)
    - [AddLabelToUserByIDResponse](#standard.AddLabelToUserByIDResponse)
    - [AddUserToGroupByIDRequest](#standard.AddUserToGroupByIDRequest)
    - [AddUserToGroupByIDResponse](#standard.AddUserToGroupByIDResponse)
    - [CreateGroupRequest](#standard.CreateGroupRequest)
    - [CreateGroupResponse](#standard.CreateGroupResponse)
    - [CreateLabelForUserRequest](#standard.CreateLabelForUserRequest)
    - [CreateLabelForUserResponse](#standard.CreateLabelForUserResponse)
    - [CreateLabelRequest](#standard.CreateLabelRequest)
    - [CreateLabelResponse](#standard.CreateLabelResponse)
    - [CreateUserRequest](#standard.CreateUserRequest)
    - [CreateUserResponse](#standard.CreateUserResponse)
    - [DeleteGroupByIDRequest](#standard.DeleteGroupByIDRequest)
    - [DeleteGroupByIDResponse](#standard.DeleteGroupByIDResponse)
    - [DeleteLabelByIDRequest](#standard.DeleteLabelByIDRequest)
    - [DeleteLabelByIDResponse](#standard.DeleteLabelByIDResponse)
    - [DeleteUserByIDRequest](#standard.DeleteUserByIDRequest)
    - [DeleteUserByIDResponse](#standard.DeleteUserByIDResponse)
    - [Group](#standard.Group)
    - [Label](#standard.Label)
    - [QueryGroupByIDRequest](#standard.QueryGroupByIDRequest)
    - [QueryGroupByIDResponse](#standard.QueryGroupByIDResponse)
    - [QueryGroupsRequest](#standard.QueryGroupsRequest)
    - [QueryGroupsResponse](#standard.QueryGroupsResponse)
    - [QueryLabelByIDRequest](#standard.QueryLabelByIDRequest)
    - [QueryLabelByIDResponse](#standard.QueryLabelByIDResponse)
    - [QueryLabelsRequest](#standard.QueryLabelsRequest)
    - [QueryLabelsResponse](#standard.QueryLabelsResponse)
    - [QueryUserByIDRequest](#standard.QueryUserByIDRequest)
    - [QueryUserByIDResponse](#standard.QueryUserByIDResponse)
    - [QueryUserByUsernameRequest](#standard.QueryUserByUsernameRequest)
    - [QueryUserByUsernameResponse](#standard.QueryUserByUsernameResponse)
    - [QueryUsersByInviterRequest](#standard.QueryUsersByInviterRequest)
    - [QueryUsersByInviterResponse](#standard.QueryUsersByInviterResponse)
    - [QueryUsersRequest](#standard.QueryUsersRequest)
    - [QueryUsersResponse](#standard.QueryUsersResponse)
    - [RemoveLabelFromUserByIDRequest](#standard.RemoveLabelFromUserByIDRequest)
    - [RemoveLabelFromUserByIDResponse](#standard.RemoveLabelFromUserByIDResponse)
    - [RemoveUserFromGroupByIDRequest](#standard.RemoveUserFromGroupByIDRequest)
    - [RemoveUserFromGroupByIDResponse](#standard.RemoveUserFromGroupByIDResponse)
    - [UpdateGroupCategoryByIDRequest](#standard.UpdateGroupCategoryByIDRequest)
    - [UpdateGroupCategoryByIDResponse](#standard.UpdateGroupCategoryByIDResponse)
    - [UpdateGroupDescriptionByIDRequest](#standard.UpdateGroupDescriptionByIDRequest)
    - [UpdateGroupDescriptionByIDResponse](#standard.UpdateGroupDescriptionByIDResponse)
    - [UpdateGroupNameByIDRequest](#standard.UpdateGroupNameByIDRequest)
    - [UpdateGroupNameByIDResponse](#standard.UpdateGroupNameByIDResponse)
    - [UpdateGroupStateByIDRequest](#standard.UpdateGroupStateByIDRequest)
    - [UpdateGroupStateByIDResponse](#standard.UpdateGroupStateByIDResponse)
    - [UpdateLabelCategoryByIDRequest](#standard.UpdateLabelCategoryByIDRequest)
    - [UpdateLabelCategoryByIDResponse](#standard.UpdateLabelCategoryByIDResponse)
    - [UpdateLabelNameByIDRequest](#standard.UpdateLabelNameByIDRequest)
    - [UpdateLabelNameByIDResponse](#standard.UpdateLabelNameByIDResponse)
    - [UpdateLabelStateByIDRequest](#standard.UpdateLabelStateByIDRequest)
    - [UpdateLabelStateByIDResponse](#standard.UpdateLabelStateByIDResponse)
    - [UpdateLabelValueByIDRequest](#standard.UpdateLabelValueByIDRequest)
    - [UpdateLabelValueByIDResponse](#standard.UpdateLabelValueByIDResponse)
    - [UpdateUserAvatarByIDRequest](#standard.UpdateUserAvatarByIDRequest)
    - [UpdateUserAvatarByIDResponse](#standard.UpdateUserAvatarByIDResponse)
    - [UpdateUserCategoryByIDRequest](#standard.UpdateUserCategoryByIDRequest)
    - [UpdateUserCategoryByIDResponse](#standard.UpdateUserCategoryByIDResponse)
    - [UpdateUserInviterByIDRequest](#standard.UpdateUserInviterByIDRequest)
    - [UpdateUserInviterByIDResponse](#standard.UpdateUserInviterByIDResponse)
    - [UpdateUserNicknameByIDRequest](#standard.UpdateUserNicknameByIDRequest)
    - [UpdateUserNicknameByIDResponse](#standard.UpdateUserNicknameByIDResponse)
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



<a name="standard.AddLabelToUserByIDRequest"></a>

### AddLabelToUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| UserID | [int64](#int64) |  | 添加标签给用户 谓语是用户 |






<a name="standard.AddLabelToUserByIDResponse"></a>

### AddLabelToUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.AddUserToGroupByIDRequest"></a>

### AddUserToGroupByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| GroupID | [int64](#int64) |  | 添加用户到组 组是谓语 |






<a name="standard.AddUserToGroupByIDResponse"></a>

### AddUserToGroupByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.CreateGroupRequest"></a>

### CreateGroupRequest
组操作


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| Category | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Description | [string](#string) |  |  |






<a name="standard.CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [Group](#standard.Group) |  |  |






<a name="standard.CreateLabelForUserRequest"></a>

### CreateLabelForUserRequest
给指定用户创建标签


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Category | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Value | [string](#string) |  |  |






<a name="standard.CreateLabelForUserResponse"></a>

### CreateLabelForUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [Label](#standard.Label) |  |  |






<a name="standard.CreateLabelRequest"></a>

### CreateLabelRequest
标签操作


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| Category | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Value | [string](#string) |  |  |






<a name="standard.CreateLabelResponse"></a>

### CreateLabelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [Label](#standard.Label) |  |  |






<a name="standard.CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Category | [string](#string) |  |  |
| Inviter | [int64](#int64) |  |  |
| Nickname | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [User](#standard.User) |  |  |






<a name="standard.DeleteGroupByIDRequest"></a>

### DeleteGroupByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="standard.DeleteGroupByIDResponse"></a>

### DeleteGroupByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.DeleteLabelByIDRequest"></a>

### DeleteLabelByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






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
| ID | [int64](#int64) |  |  |






<a name="standard.DeleteUserByIDResponse"></a>

### DeleteUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.Group"></a>

### Group
Group 组


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Category | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| CreatedTime | [string](#string) |  |  |
| UpdatedTime | [string](#string) |  |  |
| DeletedTime | [string](#string) |  |  |






<a name="standard.Label"></a>

### Label
Label 标签


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Category | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Value | [string](#string) |  |  |
| CreatedTime | [string](#string) |  |  |
| UpdatedTime | [string](#string) |  |  |
| DeletedTime | [string](#string) |  |  |






<a name="standard.QueryGroupByIDRequest"></a>

### QueryGroupByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="standard.QueryGroupByIDResponse"></a>

### QueryGroupByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [Group](#standard.Group) |  |  |






<a name="standard.QueryGroupsRequest"></a>

### QueryGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Page | [int64](#int64) |  |  |
| Limit | [int64](#int64) |  |  |






<a name="standard.QueryGroupsResponse"></a>

### QueryGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| totalPage | [int64](#int64) |  |  |
| currentPage | [int64](#int64) |  |  |
| Data | [Group](#standard.Group) | repeated |  |






<a name="standard.QueryLabelByIDRequest"></a>

### QueryLabelByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="standard.QueryLabelByIDResponse"></a>

### QueryLabelByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  | 状态 |
| Message | [string](#string) |  |  |
| Data | [Label](#standard.Label) |  |  |






<a name="standard.QueryLabelsRequest"></a>

### QueryLabelsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Page | [int64](#int64) |  |  |
| Limit | [int64](#int64) |  |  |






<a name="standard.QueryLabelsResponse"></a>

### QueryLabelsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| totalPage | [int64](#int64) |  |  |
| currentPage | [int64](#int64) |  |  |
| Data | [Label](#standard.Label) | repeated |  |






<a name="standard.QueryUserByIDRequest"></a>

### QueryUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






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






<a name="standard.QueryUsersByInviterRequest"></a>

### QueryUsersByInviterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Inviter | [int64](#int64) |  |  |
| Page | [int64](#int64) |  |  |
| Limit | [int64](#int64) |  |  |






<a name="standard.QueryUsersByInviterResponse"></a>

### QueryUsersByInviterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| totalPage | [int64](#int64) |  |  |
| currentPage | [int64](#int64) |  |  |
| Data | [User](#standard.User) | repeated |  |






<a name="standard.QueryUsersRequest"></a>

### QueryUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Page | [int64](#int64) |  |  |
| Limit | [int64](#int64) |  |  |






<a name="standard.QueryUsersResponse"></a>

### QueryUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| totalPage | [int64](#int64) |  |  |
| currentPage | [int64](#int64) |  |  |
| Data | [User](#standard.User) | repeated |  |






<a name="standard.RemoveLabelFromUserByIDRequest"></a>

### RemoveLabelFromUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| LabelID | [int64](#int64) |  | 从用户身上撕下标签 所以谓语是标签 |






<a name="standard.RemoveLabelFromUserByIDResponse"></a>

### RemoveLabelFromUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.RemoveUserFromGroupByIDRequest"></a>

### RemoveUserFromGroupByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| UserID | [int64](#int64) |  | 从组里移除用户 用户是谓语 |






<a name="standard.RemoveUserFromGroupByIDResponse"></a>

### RemoveUserFromGroupByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateGroupCategoryByIDRequest"></a>

### UpdateGroupCategoryByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Category | [string](#string) |  |  |






<a name="standard.UpdateGroupCategoryByIDResponse"></a>

### UpdateGroupCategoryByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateGroupDescriptionByIDRequest"></a>

### UpdateGroupDescriptionByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Description | [string](#string) |  |  |






<a name="standard.UpdateGroupDescriptionByIDResponse"></a>

### UpdateGroupDescriptionByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateGroupNameByIDRequest"></a>

### UpdateGroupNameByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |






<a name="standard.UpdateGroupNameByIDResponse"></a>

### UpdateGroupNameByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateGroupStateByIDRequest"></a>

### UpdateGroupStateByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| State | [string](#string) |  |  |






<a name="standard.UpdateGroupStateByIDResponse"></a>

### UpdateGroupStateByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateLabelCategoryByIDRequest"></a>

### UpdateLabelCategoryByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Category | [string](#string) |  |  |






<a name="standard.UpdateLabelCategoryByIDResponse"></a>

### UpdateLabelCategoryByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateLabelNameByIDRequest"></a>

### UpdateLabelNameByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |






<a name="standard.UpdateLabelNameByIDResponse"></a>

### UpdateLabelNameByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateLabelStateByIDRequest"></a>

### UpdateLabelStateByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| State | [string](#string) |  |  |






<a name="standard.UpdateLabelStateByIDResponse"></a>

### UpdateLabelStateByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateLabelValueByIDRequest"></a>

### UpdateLabelValueByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Value | [string](#string) |  |  |






<a name="standard.UpdateLabelValueByIDResponse"></a>

### UpdateLabelValueByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserAvatarByIDRequest"></a>

### UpdateUserAvatarByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Avatar | [string](#string) |  |  |






<a name="standard.UpdateUserAvatarByIDResponse"></a>

### UpdateUserAvatarByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserCategoryByIDRequest"></a>

### UpdateUserCategoryByIDRequest
===== //


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Category | [string](#string) |  |  |






<a name="standard.UpdateUserCategoryByIDResponse"></a>

### UpdateUserCategoryByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserInviterByIDRequest"></a>

### UpdateUserInviterByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Inviter | [int64](#int64) |  |  |






<a name="standard.UpdateUserInviterByIDResponse"></a>

### UpdateUserInviterByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserNicknameByIDRequest"></a>

### UpdateUserNicknameByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Nickname | [string](#string) |  |  |






<a name="standard.UpdateUserNicknameByIDResponse"></a>

### UpdateUserNicknameByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUserPasswordByIDRequest"></a>

### UpdateUserPasswordByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
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
| ID | [int64](#int64) |  |  |
| Category | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Inviter | [int64](#int64) |  |  |
| Nickname | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| CreatedTime | [string](#string) |  |  |
| UpdatedTime | [string](#string) |  |  |
| DeletedTime | [string](#string) |  |  |






<a name="standard.VerifyUserPasswordByIDRequest"></a>

### VerifyUserPasswordByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| Password | [string](#string) |  |  |






<a name="standard.VerifyUserPasswordByIDResponse"></a>

### VerifyUserPasswordByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






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





 


<a name="standard.State"></a>

### State
状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知 |
| SUCCESS | 1 | 成功 |
| FAILURE | 2 | 失败 |
| SERVICE_ERROR | 3 | 服务错误 |
| PARAMS_INVALID | 4 | 参数不合法 |
| ILLEGAL_REQUEST | 5 | 非法请求 |
| USER_NOT_EXIST | 8 | 用户不存在 |
| LABEL_NOT_EXIST | 6 | 标签不存在 |
| GROUP_NOT_EXIST | 7 | 分组不存在 |
| USER_ALREADY_EXISTS | 9 | 用户已经存在 |
| LABEL_ALREADY_EXISTS | 11 | 标签已经存在 |
| GROUP_ALREADY_EXISTS | 12 | 分组已经存在 |
| USER_ALREADY_DELETE | 14 | 用户已经删除 |
| LABEL_ALREADY_DELETE | 15 | 标签已经删除 |
| GROUP_ALREADY_DELETE | 16 | 分组已经删除 |
| DB_OPERATION_FATLURE | 13 | 数据库操作失败 |


 

 


<a name="standard.Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#standard.CreateUserRequest) | [CreateUserResponse](#standard.CreateUserResponse) | 用户操作 |
| QueryUsers | [QueryUsersRequest](#standard.QueryUsersRequest) | [QueryUsersResponse](#standard.QueryUsersResponse) |  |
| QueryUserByID | [QueryUserByIDRequest](#standard.QueryUserByIDRequest) | [QueryUserByIDResponse](#standard.QueryUserByIDResponse) |  |
| QueryUsersByInviter | [QueryUsersByInviterRequest](#standard.QueryUsersByInviterRequest) | [QueryUsersByInviterResponse](#standard.QueryUsersByInviterResponse) |  |
| DeleteUserByID | [DeleteUserByIDRequest](#standard.DeleteUserByIDRequest) | [DeleteUserByIDResponse](#standard.DeleteUserByIDResponse) |  |
| QueryUserByUsername | [QueryUserByUsernameRequest](#standard.QueryUserByUsernameRequest) | [QueryUserByUsernameResponse](#standard.QueryUserByUsernameResponse) |  |
| UpdateUserCategoryByID | [UpdateUserCategoryByIDRequest](#standard.UpdateUserCategoryByIDRequest) | [UpdateUserCategoryByIDResponse](#standard.UpdateUserCategoryByIDResponse) |  |
| UpdateUserAvatarByID | [UpdateUserAvatarByIDRequest](#standard.UpdateUserAvatarByIDRequest) | [UpdateUserAvatarByIDResponse](#standard.UpdateUserAvatarByIDResponse) |  |
| UpdateUserInviterByID | [UpdateUserInviterByIDRequest](#standard.UpdateUserInviterByIDRequest) | [UpdateUserInviterByIDResponse](#standard.UpdateUserInviterByIDResponse) |  |
| UpdateUserNicknameByID | [UpdateUserNicknameByIDRequest](#standard.UpdateUserNicknameByIDRequest) | [UpdateUserNicknameByIDResponse](#standard.UpdateUserNicknameByIDResponse) |  |
| UpdateUserPasswordByID | [UpdateUserPasswordByIDRequest](#standard.UpdateUserPasswordByIDRequest) | [UpdateUserPasswordByIDResponse](#standard.UpdateUserPasswordByIDResponse) |  |
| VerifyUserPasswordByID | [VerifyUserPasswordByIDRequest](#standard.VerifyUserPasswordByIDRequest) | [VerifyUserPasswordByIDResponse](#standard.VerifyUserPasswordByIDResponse) |  |
| VerifyUserPasswordByUsername | [VerifyUserPasswordByUsernameRequest](#standard.VerifyUserPasswordByUsernameRequest) | [VerifyUserPasswordByUsernameResponse](#standard.VerifyUserPasswordByUsernameResponse) |  |
| QueryLabels | [QueryLabelsRequest](#standard.QueryLabelsRequest) | [QueryLabelsResponse](#standard.QueryLabelsResponse) | 标签操作 标签用来处理其他额外的用户数据、例如一些地址 手机 邮箱等信息 创建一个标签 然后分配给一个用户 多个用户可以共享同一个标签（共有数据） |
| CreateLabel | [CreateLabelRequest](#standard.CreateLabelRequest) | [CreateLabelResponse](#standard.CreateLabelResponse) |  |
| CreateLabelForUser | [CreateLabelForUserRequest](#standard.CreateLabelForUserRequest) | [CreateLabelForUserResponse](#standard.CreateLabelForUserResponse) |  |
| QueryLabelByID | [QueryLabelByIDRequest](#standard.QueryLabelByIDRequest) | [QueryLabelByIDResponse](#standard.QueryLabelByIDResponse) |  |
| DeleteLabelByID | [DeleteLabelByIDRequest](#standard.DeleteLabelByIDRequest) | [DeleteLabelByIDResponse](#standard.DeleteLabelByIDResponse) |  |
| UpdateLabelNameByID | [UpdateLabelNameByIDRequest](#standard.UpdateLabelNameByIDRequest) | [UpdateLabelNameByIDResponse](#standard.UpdateLabelNameByIDResponse) |  |
| UpdateLabelCategoryByID | [UpdateLabelCategoryByIDRequest](#standard.UpdateLabelCategoryByIDRequest) | [UpdateLabelCategoryByIDResponse](#standard.UpdateLabelCategoryByIDResponse) |  |
| UpdateLabelStateByID | [UpdateLabelStateByIDRequest](#standard.UpdateLabelStateByIDRequest) | [UpdateLabelStateByIDResponse](#standard.UpdateLabelStateByIDResponse) |  |
| UpdateLabelValueByID | [UpdateLabelValueByIDRequest](#standard.UpdateLabelValueByIDRequest) | [UpdateLabelValueByIDResponse](#standard.UpdateLabelValueByIDResponse) |  |
| AddLabelToUserByID | [AddLabelToUserByIDRequest](#standard.AddLabelToUserByIDRequest) | [AddLabelToUserByIDResponse](#standard.AddLabelToUserByIDResponse) |  |
| RemoveLabelFromUserByID | [RemoveLabelFromUserByIDRequest](#standard.RemoveLabelFromUserByIDRequest) | [RemoveLabelFromUserByIDResponse](#standard.RemoveLabelFromUserByIDResponse) |  |
| CreateGroup | [CreateGroupRequest](#standard.CreateGroupRequest) | [CreateGroupResponse](#standard.CreateGroupResponse) | 组操作 同一个用户可以存在于多个组里 |
| QueryGroups | [QueryGroupsRequest](#standard.QueryGroupsRequest) | [QueryGroupsResponse](#standard.QueryGroupsResponse) |  |
| QueryGroupByID | [QueryGroupByIDRequest](#standard.QueryGroupByIDRequest) | [QueryGroupByIDResponse](#standard.QueryGroupByIDResponse) |  |
| DeleteGroupByID | [DeleteGroupByIDRequest](#standard.DeleteGroupByIDRequest) | [DeleteGroupByIDResponse](#standard.DeleteGroupByIDResponse) |  |
| UpdateGroupNameByID | [UpdateGroupNameByIDRequest](#standard.UpdateGroupNameByIDRequest) | [UpdateGroupNameByIDResponse](#standard.UpdateGroupNameByIDResponse) |  |
| UpdateGroupCategoryByID | [UpdateGroupCategoryByIDRequest](#standard.UpdateGroupCategoryByIDRequest) | [UpdateGroupCategoryByIDResponse](#standard.UpdateGroupCategoryByIDResponse) |  |
| UpdateGroupStateByID | [UpdateGroupStateByIDRequest](#standard.UpdateGroupStateByIDRequest) | [UpdateGroupStateByIDResponse](#standard.UpdateGroupStateByIDResponse) |  |
| UpdateGroupDescriptionByID | [UpdateGroupDescriptionByIDRequest](#standard.UpdateGroupDescriptionByIDRequest) | [UpdateGroupDescriptionByIDResponse](#standard.UpdateGroupDescriptionByIDResponse) |  |
| AddUserToGroupByID | [AddUserToGroupByIDRequest](#standard.AddUserToGroupByIDRequest) | [AddUserToGroupByIDResponse](#standard.AddUserToGroupByIDResponse) | 组关系操作 |
| RemoveUserFromGroupByID | [RemoveUserFromGroupByIDRequest](#standard.RemoveUserFromGroupByIDRequest) | [RemoveUserFromGroupByIDResponse](#standard.RemoveUserFromGroupByIDResponse) |  |

 



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

