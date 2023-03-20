# \DatabaseApi

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDocs**](DatabaseApi.md#BulkDocs) | **Post** /{db}/_bulk_docs | The bulk document API allows you to create and update multiple documents at the same time within a single request.
[**BulkGet**](DatabaseApi.md#BulkGet) | **Post** /{db}/_bulk_get | This method can be called to query several documents in bulk.
[**DbSecurityGet**](DatabaseApi.md#DbSecurityGet) | **Get** /{db}/_security | Returns the current security object from the specified database.
[**Delete**](DatabaseApi.md#Delete) | **Delete** /{db} | Deletes the specified database, and all the documents and attachments contained within it.
[**DesignDocAllGet**](DatabaseApi.md#DesignDocAllGet) | **Get** /{db}/_design_docs | Returns a JSON structure of all of the design documents in a given database.
[**DesignDocAllPost**](DatabaseApi.md#DesignDocAllPost) | **Post** /{db}/_design_docs | POST _design_docs functionality supports identical parameters and behavior as specified in the GET /{db}/_design_docs
[**DocGetAll**](DatabaseApi.md#DocGetAll) | **Get** /{db}/_all_docs | Executes the built-in _all_docs view
[**DocPostAll**](DatabaseApi.md#DocPostAll) | **Post** /{db}/_all_docs | Executes the built-in _all_docs view
[**Exists**](DatabaseApi.md#Exists) | **Head** /{db} | Returns the HTTP Headers containing a minimal amount of information about the specified database.
[**Get**](DatabaseApi.md#Get) | **Get** /{db} | Gets information about the specified database.
[**Put**](DatabaseApi.md#Put) | **Put** /{db} | Creates a new database.
[**SbSecurityPut**](DatabaseApi.md#SbSecurityPut) | **Put** /{db}/_security | Sets the security object for the given database.



## BulkDocs

> []BulkResponse BulkDocs(ctx, db).Body(body).Execute()

The bulk document API allows you to create and update multiple documents at the same time within a single request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := *openapiclient.NewBody3() // Body3 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.BulkDocs(context.Background(), db).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.BulkDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDocs`: []BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.BulkDocs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Body3**](Body3.md) |  | 

### Return type

[**[]BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGet

> Results BulkGet(ctx, db).Body(body).Revs(revs).Execute()

This method can be called to query several documents in bulk.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := *openapiclient.NewBody2() // Body2 | List of document objects, with id, and optionally rev and atts_since
    revs := true // bool | Give the revisions history (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.BulkGet(context.Background(), db).Body(body).Revs(revs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.BulkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGet`: Results
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.BulkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Body2**](Body2.md) | List of document objects, with id, and optionally rev and atts_since | 
 **revs** | **bool** | Give the revisions history | 

### Return type

[**Results**](Results.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/related, multipart/mixed
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbSecurityGet

> InlineResponse2005 DbSecurityGet(ctx, db).Execute()

Returns the current security object from the specified database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DbSecurityGet(context.Background(), db).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DbSecurityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DbSecurityGet`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DbSecurityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbSecurityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plan

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> OK Delete(ctx, db).Execute()

Deletes the specified database, and all the documents and attachments contained within it.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.Delete(context.Background(), db).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: OK
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OK**](OK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DesignDocAllGet

> Pagination DesignDocAllGet(ctx, db).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).IncludeDocs(includeDocs).InclusiveEnd(inclusiveEnd).Key(key).Keys(keys).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).UpdateSeq(updateSeq).Execute()

Returns a JSON structure of all of the design documents in a given database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    conflicts := true // bool | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. (optional)
    descending := true // bool | Return the documents in descending order by key. Default is false. (optional)
    endkey := "endkey_example" // string | Stop returning records when the specified key is reached. Optional. (optional)
    endKey := "endKey_example" // string | Alias for endkey param. (optional)
    endkeyDocid := "endkeyDocid_example" // string | Stop returning records when the specified design document ID is reached. Optional. (optional)
    endKeyDocId := "endKeyDocId_example" // string | Alias for endkey_docid param. (optional)
    includeDocs := true // bool | Include the full content of the design documents in the return. Default is false. (optional)
    inclusiveEnd := true // bool | Specifies whether the specified end key should be included in the result. Default is true. (optional)
    key := "key_example" // string | Return only design documents that match the specified key. Optional. (optional)
    keys := "keys_example" // string | Return only design documents that match the specified keys. Optional. (optional)
    limit := int32(56) // int32 | Limit the number of the returned design documents to the specified number. Optional. (optional)
    skip := int32(56) // int32 | Skip this number of records before starting to return the results. Default is 0. (optional)
    startkey := "startkey_example" // string | Return records starting with the specified key. Optional. (optional)
    startKey := "startKey_example" // string | Alias for startkey param. (optional)
    startkeyDocid := "startkeyDocid_example" // string | Return records starting with the specified design document ID. Optional. (optional)
    startKeyDocId := "startKeyDocId_example" // string | Alias for startkey_docid param (optional)
    updateSeq := true // bool | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DesignDocAllGet(context.Background(), db).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).IncludeDocs(includeDocs).InclusiveEnd(inclusiveEnd).Key(key).Keys(keys).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DesignDocAllGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DesignDocAllGet`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DesignDocAllGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDesignDocAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conflicts** | **bool** | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. | 
 **descending** | **bool** | Return the documents in descending order by key. Default is false. | 
 **endkey** | **string** | Stop returning records when the specified key is reached. Optional. | 
 **endKey** | **string** | Alias for endkey param. | 
 **endkeyDocid** | **string** | Stop returning records when the specified design document ID is reached. Optional. | 
 **endKeyDocId** | **string** | Alias for endkey_docid param. | 
 **includeDocs** | **bool** | Include the full content of the design documents in the return. Default is false. | 
 **inclusiveEnd** | **bool** | Specifies whether the specified end key should be included in the result. Default is true. | 
 **key** | **string** | Return only design documents that match the specified key. Optional. | 
 **keys** | **string** | Return only design documents that match the specified keys. Optional. | 
 **limit** | **int32** | Limit the number of the returned design documents to the specified number. Optional. | 
 **skip** | **int32** | Skip this number of records before starting to return the results. Default is 0. | 
 **startkey** | **string** | Return records starting with the specified key. Optional. | 
 **startKey** | **string** | Alias for startkey param. | 
 **startkeyDocid** | **string** | Return records starting with the specified design document ID. Optional. | 
 **startKeyDocId** | **string** | Alias for startkey_docid param | 
 **updateSeq** | **bool** | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. | 

### Return type

[**Pagination**](Pagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plan

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DesignDocAllPost

> Pagination DesignDocAllPost(ctx, db).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).IncludeDocs(includeDocs).InclusiveEnd(inclusiveEnd).Key(key).Keys(keys).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).UpdateSeq(updateSeq).Execute()

POST _design_docs functionality supports identical parameters and behavior as specified in the GET /{db}/_design_docs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := *openapiclient.NewBody4() // Body4 | 
    conflicts := true // bool | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. (optional)
    descending := true // bool | Return the documents in descending order by key. Default is false. (optional)
    endkey := "endkey_example" // string | Stop returning records when the specified key is reached. Optional. (optional)
    endKey := "endKey_example" // string | Alias for endkey param. (optional)
    endkeyDocid := "endkeyDocid_example" // string | Stop returning records when the specified design document ID is reached. Optional. (optional)
    endKeyDocId := "endKeyDocId_example" // string | Alias for endkey_docid param. (optional)
    includeDocs := true // bool | Include the full content of the design documents in the return. Default is false. (optional)
    inclusiveEnd := true // bool | Specifies whether the specified end key should be included in the result. Default is true. (optional)
    key := "key_example" // string | Return only design documents that match the specified key. Optional. (optional)
    keys := "keys_example" // string | Return only design documents that match the specified keys. Optional. (optional)
    limit := int32(56) // int32 | Limit the number of the returned design documents to the specified number. Optional. (optional)
    skip := int32(56) // int32 | Skip this number of records before starting to return the results. Default is 0. (optional)
    startkey := "startkey_example" // string | Return records starting with the specified key. Optional. (optional)
    startKey := "startKey_example" // string | Alias for startkey param. (optional)
    startkeyDocid := "startkeyDocid_example" // string | Return records starting with the specified design document ID. Optional. (optional)
    startKeyDocId := "startKeyDocId_example" // string | Alias for startkey_docid param (optional)
    updateSeq := true // bool | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DesignDocAllPost(context.Background(), db).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).IncludeDocs(includeDocs).InclusiveEnd(inclusiveEnd).Key(key).Keys(keys).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DesignDocAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DesignDocAllPost`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DesignDocAllPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDesignDocAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Body4**](Body4.md) |  | 
 **conflicts** | **bool** | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. | 
 **descending** | **bool** | Return the documents in descending order by key. Default is false. | 
 **endkey** | **string** | Stop returning records when the specified key is reached. Optional. | 
 **endKey** | **string** | Alias for endkey param. | 
 **endkeyDocid** | **string** | Stop returning records when the specified design document ID is reached. Optional. | 
 **endKeyDocId** | **string** | Alias for endkey_docid param. | 
 **includeDocs** | **bool** | Include the full content of the design documents in the return. Default is false. | 
 **inclusiveEnd** | **bool** | Specifies whether the specified end key should be included in the result. Default is true. | 
 **key** | **string** | Return only design documents that match the specified key. Optional. | 
 **keys** | **string** | Return only design documents that match the specified keys. Optional. | 
 **limit** | **int32** | Limit the number of the returned design documents to the specified number. Optional. | 
 **skip** | **int32** | Skip this number of records before starting to return the results. Default is 0. | 
 **startkey** | **string** | Return records starting with the specified key. Optional. | 
 **startKey** | **string** | Alias for startkey param. | 
 **startkeyDocid** | **string** | Return records starting with the specified design document ID. Optional. | 
 **startKeyDocId** | **string** | Alias for startkey_docid param | 
 **updateSeq** | **bool** | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. | 

### Return type

[**Pagination**](Pagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plan
- **Accept**: application/json, text/plan

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocGetAll

> Pagination DocGetAll(ctx, db).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()

Executes the built-in _all_docs view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    conflicts := true // bool | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. (optional)
    descending := true // bool | Return the documents in descending order by key. Default is false. (optional)
    endkey := "endkey_example" // string | Stop returning records when the specified key is reached. (optional)
    endKey := "endKey_example" // string | Alias for endkey param (optional)
    endkeyDocid := "endkeyDocid_example" // string | Stop returning records when the specified document ID is reached. Ignored if endkey is not set. (optional)
    endKeyDocId := "endKeyDocId_example" // string | Alias for endkey_docid. (optional)
    group := true // bool | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. (optional)
    groupLevel := int32(56) // int32 | Specify the group level to be used. Implies group is true. (optional)
    includeDocs := true // bool | Include the associated document with each row. Default is false. (optional)
    attachments := true // bool | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. (optional)
    attEncodingInfo := true // bool | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. (optional)
    inclusiveEnd := true // bool | Specifies whether the specified end key should be included in the result. Default is true. (optional)
    revsInfo := true // bool | Includes detailed information for all known document revisions. Default is false (optional)
    key := "key_example" // string | eturn only documents that match the specified key. (optional)
    keys := []string{"Inner_example"} // []string | Return only documents where the key matches one of the keys specified in the array. (optional)
    limit := int32(56) // int32 | Limit the number of the returned documents to the specified number. (optional)
    reduce := true // bool | Use the reduction function. Default is true when a reduce function is defined. (optional)
    skip := int32(56) // int32 | Skip this number of records before starting to return the results. Default is 0. (optional)
    sorted := true // bool | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. (optional)
    stable := true // bool | Whether or not the view results should be returned from a stable set of shards. Default is false. (optional)
    stale := "stale_example" // string | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable=true&update=false. update_after is equivalent to stable=true&update=lazy. The default behavior is equivalent to stable=false&update=true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  (optional)
    startkey := "startkey_example" // string | Return records starting with the specified key. (optional)
    startKey := "startKey_example" // string | Alias for startkey. (optional)
    startkeyDocid := "startkeyDocid_example" // string | Return records starting with the specified document ID. Ignored if startkey is not set. (optional)
    startKeyDocId := "startKeyDocId_example" // string | Alias for startkey_docid param (optional)
    update := "update_example" // string | Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.  (optional)
    updateSeq := true // bool | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DocGetAll(context.Background(), db).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DocGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocGetAll`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DocGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conflicts** | **bool** | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. | 
 **descending** | **bool** | Return the documents in descending order by key. Default is false. | 
 **endkey** | **string** | Stop returning records when the specified key is reached. | 
 **endKey** | **string** | Alias for endkey param | 
 **endkeyDocid** | **string** | Stop returning records when the specified document ID is reached. Ignored if endkey is not set. | 
 **endKeyDocId** | **string** | Alias for endkey_docid. | 
 **group** | **bool** | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. | 
 **groupLevel** | **int32** | Specify the group level to be used. Implies group is true. | 
 **includeDocs** | **bool** | Include the associated document with each row. Default is false. | 
 **attachments** | **bool** | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. | 
 **attEncodingInfo** | **bool** | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. | 
 **inclusiveEnd** | **bool** | Specifies whether the specified end key should be included in the result. Default is true. | 
 **revsInfo** | **bool** | Includes detailed information for all known document revisions. Default is false | 
 **key** | **string** | eturn only documents that match the specified key. | 
 **keys** | **[]string** | Return only documents where the key matches one of the keys specified in the array. | 
 **limit** | **int32** | Limit the number of the returned documents to the specified number. | 
 **reduce** | **bool** | Use the reduction function. Default is true when a reduce function is defined. | 
 **skip** | **int32** | Skip this number of records before starting to return the results. Default is 0. | 
 **sorted** | **bool** | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. | 
 **stable** | **bool** | Whether or not the view results should be returned from a stable set of shards. Default is false. | 
 **stale** | **string** | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable&#x3D;true&amp;update&#x3D;false. update_after is equivalent to stable&#x3D;true&amp;update&#x3D;lazy. The default behavior is equivalent to stable&#x3D;false&amp;update&#x3D;true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  | 
 **startkey** | **string** | Return records starting with the specified key. | 
 **startKey** | **string** | Alias for startkey. | 
 **startkeyDocid** | **string** | Return records starting with the specified document ID. Ignored if startkey is not set. | 
 **startKeyDocId** | **string** | Alias for startkey_docid param | 
 **update** | **string** | Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.  | 
 **updateSeq** | **bool** | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. | 

### Return type

[**Pagination**](Pagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocPostAll

> Pagination DocPostAll(ctx, db).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()

Executes the built-in _all_docs view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := *openapiclient.NewKeys() // Keys | 
    conflicts := true // bool | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. (optional)
    descending := true // bool | Return the documents in descending order by key. Default is false. (optional)
    endkey := "endkey_example" // string | Stop returning records when the specified key is reached. (optional)
    endKey := "endKey_example" // string | Alias for endkey param (optional)
    endkeyDocid := "endkeyDocid_example" // string | Stop returning records when the specified document ID is reached. Ignored if endkey is not set. (optional)
    endKeyDocId := "endKeyDocId_example" // string | Alias for endkey_docid. (optional)
    group := true // bool | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. (optional)
    groupLevel := int32(56) // int32 | Specify the group level to be used. Implies group is true. (optional)
    includeDocs := true // bool | Include the associated document with each row. Default is false. (optional)
    attachments := true // bool | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. (optional)
    attEncodingInfo := true // bool | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. (optional)
    inclusiveEnd := true // bool | Specifies whether the specified end key should be included in the result. Default is true. (optional)
    revsInfo := true // bool | Includes detailed information for all known document revisions. Default is false (optional)
    key := "key_example" // string | eturn only documents that match the specified key. (optional)
    keys := []string{"Inner_example"} // []string | Return only documents where the key matches one of the keys specified in the array. (optional)
    limit := int32(56) // int32 | Limit the number of the returned documents to the specified number. (optional)
    reduce := true // bool | Use the reduction function. Default is true when a reduce function is defined. (optional)
    skip := int32(56) // int32 | Skip this number of records before starting to return the results. Default is 0. (optional)
    sorted := true // bool | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. (optional)
    stable := true // bool | Whether or not the view results should be returned from a stable set of shards. Default is false. (optional)
    stale := "stale_example" // string | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable=true&update=false. update_after is equivalent to stable=true&update=lazy. The default behavior is equivalent to stable=false&update=true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  (optional)
    startkey := "startkey_example" // string | Return records starting with the specified key. (optional)
    startKey := "startKey_example" // string | Alias for startkey. (optional)
    startkeyDocid := "startkeyDocid_example" // string | Return records starting with the specified document ID. Ignored if startkey is not set. (optional)
    startKeyDocId := "startKeyDocId_example" // string | Alias for startkey_docid param (optional)
    update := "update_example" // string | Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.  (optional)
    updateSeq := true // bool | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DocPostAll(context.Background(), db).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DocPostAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocPostAll`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DocPostAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocPostAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Keys**](Keys.md) |  | 
 **conflicts** | **bool** | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. | 
 **descending** | **bool** | Return the documents in descending order by key. Default is false. | 
 **endkey** | **string** | Stop returning records when the specified key is reached. | 
 **endKey** | **string** | Alias for endkey param | 
 **endkeyDocid** | **string** | Stop returning records when the specified document ID is reached. Ignored if endkey is not set. | 
 **endKeyDocId** | **string** | Alias for endkey_docid. | 
 **group** | **bool** | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. | 
 **groupLevel** | **int32** | Specify the group level to be used. Implies group is true. | 
 **includeDocs** | **bool** | Include the associated document with each row. Default is false. | 
 **attachments** | **bool** | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. | 
 **attEncodingInfo** | **bool** | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. | 
 **inclusiveEnd** | **bool** | Specifies whether the specified end key should be included in the result. Default is true. | 
 **revsInfo** | **bool** | Includes detailed information for all known document revisions. Default is false | 
 **key** | **string** | eturn only documents that match the specified key. | 
 **keys** | **[]string** | Return only documents where the key matches one of the keys specified in the array. | 
 **limit** | **int32** | Limit the number of the returned documents to the specified number. | 
 **reduce** | **bool** | Use the reduction function. Default is true when a reduce function is defined. | 
 **skip** | **int32** | Skip this number of records before starting to return the results. Default is 0. | 
 **sorted** | **bool** | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. | 
 **stable** | **bool** | Whether or not the view results should be returned from a stable set of shards. Default is false. | 
 **stale** | **string** | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable&#x3D;true&amp;update&#x3D;false. update_after is equivalent to stable&#x3D;true&amp;update&#x3D;lazy. The default behavior is equivalent to stable&#x3D;false&amp;update&#x3D;true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  | 
 **startkey** | **string** | Return records starting with the specified key. | 
 **startKey** | **string** | Alias for startkey. | 
 **startkeyDocid** | **string** | Return records starting with the specified document ID. Ignored if startkey is not set. | 
 **startKeyDocId** | **string** | Alias for startkey_docid param | 
 **update** | **string** | Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.  | 
 **updateSeq** | **bool** | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. | 

### Return type

[**Pagination**](Pagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Exists

> Exists(ctx, db).Execute()

Returns the HTTP Headers containing a minimal amount of information about the specified database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseApi.Exists(context.Background(), db).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.Exists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Database Get(ctx, db).Execute()

Gets information about the specified database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.Get(context.Background(), db).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Database**](Database.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Put

> OK Put(ctx, db).Q(q).N(n).Partitioned(partitioned).Execute()

Creates a new database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    q := int32(56) // int32 | Shards, aka the number of range partitions. Default is 8, unless overridden in the cluster config. (optional)
    n := int32(56) // int32 | Replicas. The number of copies of the database in the cluster. The default is 3, unless overridden in the cluster config . (optional)
    partitioned := true // bool | Whether to create a partitioned database. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.Put(context.Background(), db).Q(q).N(n).Partitioned(partitioned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.Put``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Put`: OK
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.Put`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **int32** | Shards, aka the number of range partitions. Default is 8, unless overridden in the cluster config. | 
 **n** | **int32** | Replicas. The number of copies of the database in the cluster. The default is 3, unless overridden in the cluster config . | 
 **partitioned** | **bool** | Whether to create a partitioned database. Default is false. | 

### Return type

[**OK**](OK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SbSecurityPut

> OK SbSecurityPut(ctx, db).Body(body).Execute()

Sets the security object for the given database.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := *openapiclient.NewBody5() // Body5 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.SbSecurityPut(context.Background(), db).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.SbSecurityPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SbSecurityPut`: OK
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.SbSecurityPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSbSecurityPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Body5**](Body5.md) |  | 

### Return type

[**OK**](OK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plan
- **Accept**: application/json, text/plan

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

