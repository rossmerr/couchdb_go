# \IndexApi

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbFindGet**](IndexApi.md#DbFindGet) | **Post** /{db}/_find | Finds the document.
[**DbIndexGet**](IndexApi.md#DbIndexGet) | **Get** /{db}/_index | Returns the current indexes object from the specified database.
[**IndexDelete**](IndexApi.md#IndexDelete) | **Delete** /{db}/_index/{designdoc}/json/{name} | 
[**SbIndexPost**](IndexApi.md#SbIndexPost) | **Post** /{db}/_index | Sets the index for the given database.



## DbFindGet

> InlineResponse2006 DbFindGet(ctx, db).Body(body).Execute()

Finds the document.



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
    body := *openapiclient.NewQuery() // Query | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexApi.DbFindGet(context.Background(), db).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.DbFindGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DbFindGet`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.DbFindGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbFindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Query**](Query.md) |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbIndexGet

> Indexes DbIndexGet(ctx, db).Execute()

Returns the current indexes object from the specified database.



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
    resp, r, err := apiClient.IndexApi.DbIndexGet(context.Background(), db).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.DbIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DbIndexGet`: Indexes
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.DbIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Indexes**](Indexes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plan

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexDelete

> OK IndexDelete(ctx, db, designdoc, name).Execute()



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
    designdoc := "designdoc_example" // string | Design document name
    name := "name_example" // string | Search index name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexApi.IndexDelete(context.Background(), db, designdoc, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.IndexDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndexDelete`: OK
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.IndexDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**designdoc** | **string** | Design document name | 
**name** | **string** | Search index name | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OK**](OK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SbIndexPost

> IndexResponse SbIndexPost(ctx, db).Body(body).Execute()

Sets the index for the given database.

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
    body := *openapiclient.NewIndex() // Index | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexApi.SbIndexPost(context.Background(), db).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.SbIndexPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SbIndexPost`: IndexResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.SbIndexPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSbIndexPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Index**](Index.md) |  | 

### Return type

[**IndexResponse**](IndexResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

