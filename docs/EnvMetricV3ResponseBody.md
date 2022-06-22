# EnvMetricV3ResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | created time | 
**CveID** | **string** | CveID of envMetricV3 | 
**Ma** | **string** | MA of envMetricV3 | 
**Mac** | **string** | MAC of envMetricV3 | 
**Mav** | **string** | MAV of envMetricV3 | 
**Mc** | **string** | MC of envMetricV3 | 
**Mpr** | **string** | MPR of envMetricV3 | 
**Ms** | **string** | MS of envMetricV3 | 
**Mui** | **string** | MUI of envMetricV3 | 
**RoleID** | **int64** | ServerRoleID of envMetricV3 | 
**RoleName** | **string** | ServerRoleName of envMetricV3 | 
**UpdatedAt** | **time.Time** | updated time | 

## Methods

### NewEnvMetricV3ResponseBody

`func NewEnvMetricV3ResponseBody(createdAt time.Time, cveID string, ma string, mac string, mav string, mc string, mpr string, ms string, mui string, roleID int64, roleName string, updatedAt time.Time, ) *EnvMetricV3ResponseBody`

NewEnvMetricV3ResponseBody instantiates a new EnvMetricV3ResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvMetricV3ResponseBodyWithDefaults

`func NewEnvMetricV3ResponseBodyWithDefaults() *EnvMetricV3ResponseBody`

NewEnvMetricV3ResponseBodyWithDefaults instantiates a new EnvMetricV3ResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EnvMetricV3ResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvMetricV3ResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvMetricV3ResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *EnvMetricV3ResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *EnvMetricV3ResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *EnvMetricV3ResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetMa

`func (o *EnvMetricV3ResponseBody) GetMa() string`

GetMa returns the Ma field if non-nil, zero value otherwise.

### GetMaOk

`func (o *EnvMetricV3ResponseBody) GetMaOk() (*string, bool)`

GetMaOk returns a tuple with the Ma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMa

`func (o *EnvMetricV3ResponseBody) SetMa(v string)`

SetMa sets Ma field to given value.


### GetMac

`func (o *EnvMetricV3ResponseBody) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EnvMetricV3ResponseBody) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EnvMetricV3ResponseBody) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMav

`func (o *EnvMetricV3ResponseBody) GetMav() string`

GetMav returns the Mav field if non-nil, zero value otherwise.

### GetMavOk

`func (o *EnvMetricV3ResponseBody) GetMavOk() (*string, bool)`

GetMavOk returns a tuple with the Mav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMav

`func (o *EnvMetricV3ResponseBody) SetMav(v string)`

SetMav sets Mav field to given value.


### GetMc

`func (o *EnvMetricV3ResponseBody) GetMc() string`

GetMc returns the Mc field if non-nil, zero value otherwise.

### GetMcOk

`func (o *EnvMetricV3ResponseBody) GetMcOk() (*string, bool)`

GetMcOk returns a tuple with the Mc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMc

`func (o *EnvMetricV3ResponseBody) SetMc(v string)`

SetMc sets Mc field to given value.


### GetMpr

`func (o *EnvMetricV3ResponseBody) GetMpr() string`

GetMpr returns the Mpr field if non-nil, zero value otherwise.

### GetMprOk

`func (o *EnvMetricV3ResponseBody) GetMprOk() (*string, bool)`

GetMprOk returns a tuple with the Mpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpr

`func (o *EnvMetricV3ResponseBody) SetMpr(v string)`

SetMpr sets Mpr field to given value.


### GetMs

`func (o *EnvMetricV3ResponseBody) GetMs() string`

GetMs returns the Ms field if non-nil, zero value otherwise.

### GetMsOk

`func (o *EnvMetricV3ResponseBody) GetMsOk() (*string, bool)`

GetMsOk returns a tuple with the Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMs

`func (o *EnvMetricV3ResponseBody) SetMs(v string)`

SetMs sets Ms field to given value.


### GetMui

`func (o *EnvMetricV3ResponseBody) GetMui() string`

GetMui returns the Mui field if non-nil, zero value otherwise.

### GetMuiOk

`func (o *EnvMetricV3ResponseBody) GetMuiOk() (*string, bool)`

GetMuiOk returns a tuple with the Mui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMui

`func (o *EnvMetricV3ResponseBody) SetMui(v string)`

SetMui sets Mui field to given value.


### GetRoleID

`func (o *EnvMetricV3ResponseBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *EnvMetricV3ResponseBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *EnvMetricV3ResponseBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.


### GetRoleName

`func (o *EnvMetricV3ResponseBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *EnvMetricV3ResponseBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *EnvMetricV3ResponseBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUpdatedAt

`func (o *EnvMetricV3ResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvMetricV3ResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvMetricV3ResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


