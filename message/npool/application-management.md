# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/application-management.proto](#npool/application-management.proto)
    - [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest)
    - [AuthoriseUserRequest](#application.management.v1.AuthoriseUserRequest)
    - [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest)
    - [CreateGroupRequest](#application.management.v1.CreateGroupRequest)
    - [CreateResourceRequest](#application.management.v1.CreateResourceRequest)
    - [CreateRoleRequest](#application.management.v1.CreateRoleRequest)
    - [CreateUsersRequest](#application.management.v1.CreateUsersRequest)
    - [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest)
    - [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest)
    - [DeleteGroupUsersRequest](#application.management.v1.DeleteGroupUsersRequest)
    - [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest)
    - [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest)
    - [GetGroupsRequest](#application.management.v1.GetGroupsRequest)
    - [GetGroupsResponse](#application.management.v1.GetGroupsResponse)
    - [GetGroupsResponse.GroupResponse](#application.management.v1.GetGroupsResponse.GroupResponse)
    - [GetResourcesRequest](#application.management.v1.GetResourcesRequest)
    - [GetResourcesResponse](#application.management.v1.GetResourcesResponse)
    - [GetResourcesResponse.ResourceResponse](#application.management.v1.GetResourcesResponse.ResourceResponse)
    - [GetRolesRequest](#application.management.v1.GetRolesRequest)
    - [GetRolesResponse](#application.management.v1.GetRolesResponse)
    - [GetRolesResponse.Response](#application.management.v1.GetRolesResponse.Response)
    - [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest)
    - [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse)
    - [GetUserRoleResponse.Role](#application.management.v1.GetUserRoleResponse.Role)
    - [GetUserRoleResponse.RoleResponse](#application.management.v1.GetUserRoleResponse.RoleResponse)
    - [GetUsersRequest](#application.management.v1.GetUsersRequest)
    - [GetUsersResponse](#application.management.v1.GetUsersResponse)
    - [GetUsersResponse.UsersResponse](#application.management.v1.GetUsersResponse.UsersResponse)
    - [GrantUserLoginPolicyRequest](#application.management.v1.GrantUserLoginPolicyRequest)
    - [GroupInfo](#application.management.v1.GroupInfo)
    - [NoDataResponse](#application.management.v1.NoDataResponse)
    - [PageInfo](#application.management.v1.PageInfo)
    - [RemoveUserLoginPoliciesRequest](#application.management.v1.RemoveUserLoginPoliciesRequest)
    - [RemoveUserPoliciesRequest](#application.management.v1.RemoveUserPoliciesRequest)
    - [ResourceInfo](#application.management.v1.ResourceInfo)
    - [ResourcePolicy](#application.management.v1.ResourcePolicy)
    - [RolesInfo](#application.management.v1.RolesInfo)
    - [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest)
    - [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest)
    - [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest)
    - [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest)
    - [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest)
    - [UpdateRolePoliciesRequest](#application.management.v1.UpdateRolePoliciesRequest)
    - [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest)
    - [UserInfo](#application.management.v1.UserInfo)
    - [VersionResponse](#application.management.v1.VersionResponse)
  
    - [ApplicationManagement](#application.management.v1.ApplicationManagement)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/application-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/application-management.proto



<a name="application.management.v1.AddGroupUsersRequest"></a>

### AddGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| GroupId | [string](#string) |  |  |
| Users | [string](#string) | repeated |  |






<a name="application.management.v1.AuthoriseUserRequest"></a>

### AuthoriseUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| Policies | [ResourcePolicy](#application.management.v1.ResourcePolicy) | repeated |  |






<a name="application.management.v1.CreateApplicationRequest"></a>

### CreateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ApplicationName | [string](#string) |  |  |
| ApplicationOwner | [string](#string) |  |  |
| HomepageUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |
| ApplicationLogo | [string](#string) |  |  |






<a name="application.management.v1.CreateGroupRequest"></a>

### CreateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| Group | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| Resource | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| RoleName | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Policies | [ResourcePolicy](#application.management.v1.ResourcePolicy) | repeated | Optional. |






<a name="application.management.v1.CreateUsersRequest"></a>

### CreateUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Users | [string](#string) | repeated |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.DeleteApplicationRequest"></a>

### DeleteApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupRequest"></a>

### DeleteGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupUsersRequest"></a>

### DeleteGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| GroupId | [string](#string) |  |  |
| Users | [string](#string) | repeated |  |






<a name="application.management.v1.DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| ResourceId | [string](#string) |  |  |






<a name="application.management.v1.DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.GetGroupsRequest"></a>

### GetGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.GetGroupsResponse"></a>

### GetGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GetGroupsResponse.GroupResponse](#application.management.v1.GetGroupsResponse.GroupResponse) |  |  |






<a name="application.management.v1.GetGroupsResponse.GroupResponse"></a>

### GetGroupsResponse.GroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupInfos | [GroupInfo](#application.management.v1.GroupInfo) | repeated |  |
| PageInfo | [PageInfo](#application.management.v1.PageInfo) |  |  |






<a name="application.management.v1.GetResourcesRequest"></a>

### GetResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.GetResourcesResponse"></a>

### GetResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GetResourcesResponse.ResourceResponse](#application.management.v1.GetResourcesResponse.ResourceResponse) |  |  |






<a name="application.management.v1.GetResourcesResponse.ResourceResponse"></a>

### GetResourcesResponse.ResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceInfos | [ResourceInfo](#application.management.v1.ResourceInfo) | repeated |  |
| PageInfo | [PageInfo](#application.management.v1.PageInfo) |  |  |






<a name="application.management.v1.GetRolesRequest"></a>

### GetRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GetRolesResponse.Response](#application.management.v1.GetRolesResponse.Response) |  |  |






<a name="application.management.v1.GetRolesResponse.Response"></a>

### GetRolesResponse.Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#application.management.v1.PageInfo) |  |  |
| RolesInfo | [RolesInfo](#application.management.v1.RolesInfo) | repeated |  |






<a name="application.management.v1.GetUserRoleRequest"></a>

### GetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| UserId | [string](#string) |  |  |






<a name="application.management.v1.GetUserRoleResponse"></a>

### GetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GetUserRoleResponse.RoleResponse](#application.management.v1.GetUserRoleResponse.RoleResponse) |  |  |






<a name="application.management.v1.GetUserRoleResponse.Role"></a>

### GetUserRoleResponse.Role



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |






<a name="application.management.v1.GetUserRoleResponse.RoleResponse"></a>

### GetUserRoleResponse.RoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Roles | [GetUserRoleResponse.Role](#application.management.v1.GetUserRoleResponse.Role) | repeated |  |
| PageInfo | [PageInfo](#application.management.v1.PageInfo) |  |  |






<a name="application.management.v1.GetUsersRequest"></a>

### GetUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.GetUsersResponse"></a>

### GetUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GetUsersResponse.UsersResponse](#application.management.v1.GetUsersResponse.UsersResponse) |  |  |






<a name="application.management.v1.GetUsersResponse.UsersResponse"></a>

### GetUsersResponse.UsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfos | [UserInfo](#application.management.v1.UserInfo) | repeated |  |
| PageInfo | [PageInfo](#application.management.v1.PageInfo) |  |  |






<a name="application.management.v1.GrantUserLoginPolicyRequest"></a>

### GrantUserLoginPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| Users | [string](#string) | repeated |  |






<a name="application.management.v1.GroupInfo"></a>

### GroupInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupId | [string](#string) |  |  |
| GroupName | [string](#string) |  |  |
| GroupType | [string](#string) |  |  |
| GroupLogo | [string](#string) |  |  |
| GroupOwner | [string](#string) |  |  |






<a name="application.management.v1.NoDataResponse"></a>

### NoDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [string](#string) |  |  |






<a name="application.management.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageCount | [int32](#int32) |  |  |
| Total | [int32](#int32) |  |  |






<a name="application.management.v1.RemoveUserLoginPoliciesRequest"></a>

### RemoveUserLoginPoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| Users | [string](#string) | repeated |  |






<a name="application.management.v1.RemoveUserPoliciesRequest"></a>

### RemoveUserPoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| Policies | [ResourcePolicy](#application.management.v1.ResourcePolicy) | repeated |  |






<a name="application.management.v1.ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceId | [string](#string) |  |  |
| ResourceName | [string](#string) |  |  |
| ResourceDescription | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Creator | [string](#string) |  |  |






<a name="application.management.v1.ResourcePolicy"></a>

### ResourcePolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceId | [string](#string) |  |  |
| Action | [string](#string) |  |  |






<a name="application.management.v1.RolesInfo"></a>

### RolesInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| Policies | [ResourcePolicy](#application.management.v1.ResourcePolicy) | repeated |  |
| users | [UserInfo](#application.management.v1.UserInfo) | repeated |  |






<a name="application.management.v1.SetUserRoleRequest"></a>

### SetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) | repeated |  |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.UnSetUserRoleRequest"></a>

### UnSetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) | repeated |  |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.UpdateApplicationRequest"></a>

### UpdateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| ApplicationName | [string](#string) |  |  |
| ApplicationOwner | [string](#string) |  |  |
| HomepageUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |
| ApplicationLogo | [string](#string) |  |  |






<a name="application.management.v1.UpdateGroupRequest"></a>

### UpdateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| GroupId | [string](#string) |  |  |
| Group | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| ResourceId | [string](#string) |  |  |
| Resource | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.UpdateRolePoliciesRequest"></a>

### UpdateRolePoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| Policies | [ResourcePolicy](#application.management.v1.ResourcePolicy) | repeated |  |






<a name="application.management.v1.UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| RoleId | [string](#string) |  |  |
| RoleName | [string](#string) |  |  |
| Description | [string](#string) |  |  |






<a name="application.management.v1.UserInfo"></a>

### UserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Age | [int32](#int32) |  |  |
| Gender | [string](#string) |  |  |
| Region | [string](#string) |  |  |
| Birthday | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Province | [string](#string) |  |  |
| City | [string](#string) |  |  |
| PhoneNumber | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |






<a name="application.management.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="application.management.v1.ApplicationManagement"></a>

### ApplicationManagement
rpc Echo (StringMessage) returns (StringMessage){
    option (google.api.http) = {
        post: &#34;/v1/echo&#34;
        body: &#34;*&#34;
    };
}

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#application.management.v1.VersionResponse) | Method Version |
| CreateApplication | [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Create an application. |
| UpdateApplication | [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Update an application&#39;s basic info. |
| DeleteApplication | [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Delete an application. During Deleting, system will use apis to delete app group, delete app resources and app&#39;s policies in anubis. Need apis: |
| CreateRole | [CreateRoleRequest](#application.management.v1.CreateRoleRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Create a role in app. When create a role, we can authorise policies to it or not. Need APIs: https://anubis.npool.top/v1/role, https://anubis.npool.top/v1/authorise-role. |
| UpdateRolePolicies | [UpdateRolePoliciesRequest](#application.management.v1.UpdateRolePoliciesRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Update role&#39;s policies. Need APIs: https://anubis.npool.top/v1/authorise-role |
| UpdateRole | [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Update role&#39;s basic info. |
| DeleteRole | [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Delete role from app. Need APIs: https://anubis.npool.top/v1/cancel-role-policies, https://anubis.npool.top/v1/remove-role. |
| SetUserRole | [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Set role to user. When set role to a user, we will use api to authorise policies to user. Need APIs: https://anubis.npool.top/v1/user/role. |
| UnSetUserRole | [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Unset role of user. When unset role to user, we also need to use api to remove role from user. Need APIs: https://anubis.npool.top/v1/remove-user. |
| GetRoles | [GetRolesRequest](#application.management.v1.GetRolesRequest) | [GetRolesResponse](#application.management.v1.GetRolesResponse) | Get all roles of an application. When get roles, we need to get all infos of roles which include its policies. Need APIs: https://anubis.npool.top/v1/role/{role_id}/policies, https://pyramids.npool.top/v1/role/{RoleId}/users, https://user.npool.top/v1/user/{UserId}. |
| GetUsers | [GetUsersRequest](#application.management.v1.GetUsersRequest) | [GetUsersResponse](#application.management.v1.GetUsersResponse) | Get users from app. |
| GetUserRole | [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest) | [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse) | Get all roles of the user. |
| CreateUsers | [CreateUsersRequest](#application.management.v1.CreateUsersRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Add users into app. When add users, admin can choose to authorise policies to user and authorise roles to user. Need APIs: https://pyramids.npool.top/set-user-role, https://pyramids.npool.top/user/policies. |
| CreateGroup | [CreateGroupRequest](#application.management.v1.CreateGroupRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Create group in an application. When create a group, admin can add users into this group at the same time. |
| AddGroupUsers | [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Add users into group. |
| DeleteGroupUsers | [DeleteGroupUsersRequest](#application.management.v1.DeleteGroupUsersRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Remove users from group. |
| UpdateGroup | [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Update group info. |
| DeleteGroup | [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Delete group from app. When Delete group, we also need to remove users out from group. Need api: https://user.npool.top/v1/remove/group/users. |
| GetGroups | [GetGroupsRequest](#application.management.v1.GetGroupsRequest) | [GetGroupsResponse](#application.management.v1.GetGroupsResponse) | Get groups from app. |
| CreateResource | [CreateResourceRequest](#application.management.v1.CreateResourceRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Create resource for app. |
| UpdateResource | [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Update resource of app. |
| GetResources | [GetResourcesRequest](#application.management.v1.GetResourcesRequest) | [GetResourcesResponse](#application.management.v1.GetResourcesResponse) | Get all resources from app. |
| DeleteResource | [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest) | [NoDataResponse](#application.management.v1.NoDataResponse) | Delete resource from app. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

