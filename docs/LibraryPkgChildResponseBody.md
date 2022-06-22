# LibraryPkgChildResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | crated time of package or cpe | 
**Id** | **int64** | ID of library package | 
**Name** | **string** | Name of library package | 
**UpdatedAt** | **time.Time** | updated time of package or cpe | 
**Version** | **string** | Version of library package | 

## Methods

### NewLibraryPkgChildResponseBody

`func NewLibraryPkgChildResponseBody(createdAt time.Time, id int64, name string, updatedAt time.Time, version string, ) *LibraryPkgChildResponseBody`

NewLibraryPkgChildResponseBody instantiates a new LibraryPkgChildResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryPkgChildResponseBodyWithDefaults

`func NewLibraryPkgChildResponseBodyWithDefaults() *LibraryPkgChildResponseBody`

NewLibraryPkgChildResponseBodyWithDefaults instantiates a new LibraryPkgChildResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LibraryPkgChildResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LibraryPkgChildResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LibraryPkgChildResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *LibraryPkgChildResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryPkgChildResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryPkgChildResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *LibraryPkgChildResponseBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryPkgChildResponseBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryPkgChildResponseBody) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *LibraryPkgChildResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LibraryPkgChildResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LibraryPkgChildResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *LibraryPkgChildResponseBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LibraryPkgChildResponseBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LibraryPkgChildResponseBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


