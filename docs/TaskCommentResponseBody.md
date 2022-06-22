# TaskCommentResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Comment content of TaskComment | 
**CreatedAt** | **time.Time** | created time of TaskComment | 
**Id** | **int64** | ID of TaskComment | 
**Type** | **string** | Type of TaskComment | 
**UpdatedAt** | **time.Time** | updated time of TaskComment | 
**UserID** | **int64** | UserID of TaskComment | 
**UserName** | **string** | UserName of TaskComment | 

## Methods

### NewTaskCommentResponseBody

`func NewTaskCommentResponseBody(comment string, createdAt time.Time, id int64, type_ string, updatedAt time.Time, userID int64, userName string, ) *TaskCommentResponseBody`

NewTaskCommentResponseBody instantiates a new TaskCommentResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCommentResponseBodyWithDefaults

`func NewTaskCommentResponseBodyWithDefaults() *TaskCommentResponseBody`

NewTaskCommentResponseBodyWithDefaults instantiates a new TaskCommentResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *TaskCommentResponseBody) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TaskCommentResponseBody) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TaskCommentResponseBody) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *TaskCommentResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskCommentResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskCommentResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *TaskCommentResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskCommentResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskCommentResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *TaskCommentResponseBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskCommentResponseBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskCommentResponseBody) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *TaskCommentResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskCommentResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskCommentResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserID

`func (o *TaskCommentResponseBody) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *TaskCommentResponseBody) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *TaskCommentResponseBody) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetUserName

`func (o *TaskCommentResponseBody) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TaskCommentResponseBody) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TaskCommentResponseBody) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


