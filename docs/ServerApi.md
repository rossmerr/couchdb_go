# \ServerApi

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveTasks**](ServerApi.md#ActiveTasks) | **Get** /_active_tasks | List of running tasks, including the task type, name, status and process ID.
[**AllDBs**](ServerApi.md#AllDBs) | **Get** /_all_dbs | Returns a list of all the databases in the CouchDB instance.
[**ClusterSetupGet**](ServerApi.md#ClusterSetupGet) | **Get** /_cluster_setup | Returns the status of the node or cluster, per the cluster setup wizard.
[**ClusterSetupPost**](ServerApi.md#ClusterSetupPost) | **Post** /_cluster_setup | Configure a node as a single (standalone) node, as part of a cluster, or finalise a cluster.
[**DBsInfo**](ServerApi.md#DBsInfo) | **Post** /_dbs_info | Returns information of a list of the specified databases in the CouchDB instance.
[**Membership**](ServerApi.md#Membership) | **Get** /_membership | Displays the nodes that are part of the cluster as cluster_nodes.
[**MetaInformation**](ServerApi.md#MetaInformation) | **Get** / | Accessing the root of a CouchDB instance returns meta information about the instance.
[**Replication**](ServerApi.md#Replication) | **Post** /_replicate | Request, configure, or stop, a replication operation.
[**SearchAnalyze**](ServerApi.md#SearchAnalyze) | **Post** /_search_analyze | Tests the results of Lucene analyzer tokenization on sample text.
[**Up**](ServerApi.md#Up) | **Get** /_up | Confirms that the server is up, running, and ready to respond to requests.



## ActiveTasks

> []Task ActiveTasks(ctx).Execute()

List of running tasks, including the task type, name, status and process ID.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ActiveTasks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ActiveTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveTasks`: []Task
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ActiveTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActiveTasksRequest struct via the builder pattern


### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllDBs

> []string AllDBs(ctx).Descending(descending).Endkey(endkey).EndKey(endKey).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).Execute()

Returns a list of all the databases in the CouchDB instance.

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
    descending := true // bool | Return the databases in descending order by key. Default is false. (optional)
    endkey := "endkey_example" // string | Stop returning databases when the specified key is reached. (optional)
    endKey := "endKey_example" // string | Alias for endkey param (optional)
    limit := int32(56) // int32 | Limit the number of the returned databases to the specified number. (optional)
    skip := int32(56) // int32 | Skip this number of databases before starting to return the results. Default is 0. (optional)
    startkey := "startkey_example" // string | Return databases starting with the specified key. (optional)
    startKey := "startKey_example" // string | Alias for startkey. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.AllDBs(context.Background()).Descending(descending).Endkey(endkey).EndKey(endKey).Limit(limit).Skip(skip).Startkey(startkey).StartKey(startKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.AllDBs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllDBs`: []string
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.AllDBs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllDBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descending** | **bool** | Return the databases in descending order by key. Default is false. | 
 **endkey** | **string** | Stop returning databases when the specified key is reached. | 
 **endKey** | **string** | Alias for endkey param | 
 **limit** | **int32** | Limit the number of the returned databases to the specified number. | 
 **skip** | **int32** | Skip this number of databases before starting to return the results. Default is 0. | 
 **startkey** | **string** | Return databases starting with the specified key. | 
 **startKey** | **string** | Alias for startkey. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterSetupGet

> InlineResponse2001 ClusterSetupGet(ctx).Execute()

Returns the status of the node or cluster, per the cluster setup wizard.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ClusterSetupGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ClusterSetupGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterSetupGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ClusterSetupGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClusterSetupGetRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterSetupPost

> OK ClusterSetupPost(ctx).Body(body).Execute()

Configure a node as a single (standalone) node, as part of a cluster, or finalise a cluster.

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
    body := *openapiclient.NewBody() // Body | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.ClusterSetupPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.ClusterSetupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterSetupPost`: OK
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.ClusterSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClusterSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body**](Body.md) |  | 

### Return type

[**OK**](OK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DBsInfo

> []InlineResponse200 DBsInfo(ctx).Body(body).Execute()

Returns information of a list of the specified databases in the CouchDB instance.



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
    body := *openapiclient.NewKeys() // Keys | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DBsInfo(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DBsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DBsInfo`: []InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DBsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDBsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Keys**](Keys.md) |  | 

### Return type

[**[]InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Membership

> InlineResponse2002 Membership(ctx).Execute()

Displays the nodes that are part of the cluster as cluster_nodes.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.Membership(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.Membership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Membership`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.Membership`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMembershipRequest struct via the builder pattern


### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetaInformation

> []Server MetaInformation(ctx).Execute()

Accessing the root of a CouchDB instance returns meta information about the instance.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.MetaInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.MetaInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetaInformation`: []Server
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.MetaInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetaInformationRequest struct via the builder pattern


### Return type

[**[]Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replication

> Replication Replication(ctx).Body(body).Execute()

Request, configure, or stop, a replication operation.

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
    body := *openapiclient.NewReplicate() // Replicate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.Replication(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.Replication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Replication`: Replication
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.Replication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Replicate**](Replicate.md) |  | 

### Return type

[**Replication**](Replication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAnalyze

> InlineResponse2003 SearchAnalyze(ctx).Body(body).Execute()

Tests the results of Lucene analyzer tokenization on sample text.



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
    body := *openapiclient.NewBody1() // Body1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.SearchAnalyze(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.SearchAnalyze``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAnalyze`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.SearchAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body1**](Body1.md) |  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Up

> InlineResponse2004 Up(ctx).Execute()

Confirms that the server is up, running, and ready to respond to requests.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.Up(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.Up``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Up`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.Up`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpRequest struct via the builder pattern


### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

