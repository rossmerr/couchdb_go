# \PartitionApi

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartitionDesignDocSearch**](PartitionApi.md#PartitionDesignDocSearch) | **Get** /{db}/_partition/{partition}/_design/{ddoc}/_search/{index} | Executes a search request against the named index in the specified design document.
[**PartitionDesignDocView**](PartitionApi.md#PartitionDesignDocView) | **Get** /{db}/_partition/{partition}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
[**PartitionDesignDocViewPost**](PartitionApi.md#PartitionDesignDocViewPost) | **Post** /{db}/_partition/{partition}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
[**PartitionDocGetAll**](PartitionApi.md#PartitionDocGetAll) | **Get** /{db}/_partition/{partition}/_all_docs | Executes the built-in _all_docs view
[**PartitionInfo**](PartitionApi.md#PartitionInfo) | **Get** /{db}/_partition/{partition} | This endpoint returns information describing the provided partition. It includes document and deleted document counts along with external and active data sizes.



## PartitionDesignDocSearch

> Pagination PartitionDesignDocSearch(ctx, db, partition, ddoc, index).Bookmark(bookmark).Counts(counts).Drilldown(drilldown).GroupField(groupField).GroupSort(groupSort).HighlightFields(highlightFields).HighlightPreTag(highlightPreTag).HighlightPostTag(highlightPostTag).HighlightNumber(highlightNumber).HighlightSize(highlightSize).IncludeDocs(includeDocs).IncludeFields(includeFields).Limit(limit).Q(q).Query(query).Range_(range_).Sort(sort).Stale(stale).Execute()

Executes a search request against the named index in the specified design document.



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
    partition := "partition_example" // string | Partition name
    ddoc := "ddoc_example" // string | Design document id
    index := "index_example" // string | Search index name
    bookmark := "bookmark_example" // string | A bookmark received from a previous search. This parameter enables paging through the results. If there are no more results after the bookmark, you get a response with an empty rows array and the same bookmark, confirming the end of the result list. (optional)
    counts := "counts_example" // string | An array of names of string fields for which counts are requested. The response contains counts for each unique value of this field name among the documents that match the search query. Faceting must be enabled for this parameter to function. (optional)
    drilldown := "drilldown_example" // string | This field can be used several times. Each use defines a pair with a field name and a value. The search matches only documents containing the value that was provided in the named field. It differs from using \"fieldname:value\" in the q parameter only in that the values are not analyzed. Faceting must be enabled for this parameter to function. (optional)
    groupField := "groupField_example" // string | Field by which to group search matches. :query number group_limit: Maximum group count. This field can be used only if group_field is specified.  (optional)
    groupSort := "groupSort_example" // string | This field defines the order of the groups in a search that uses group_field. The default sort order is relevance. (optional)
    highlightFields := "highlightFields_example" // string | Specifies which fields to highlight. If specified, the result object contains a highlights field with an entry for each specified field. (optional)
    highlightPreTag := "highlightPreTag_example" // string | A string that is inserted before the highlighted word in the highlights output. (optional)
    highlightPostTag := "highlightPostTag_example" // string | A string that is inserted after the highlighted word in the highlights output. (optional)
    highlightNumber := int32(56) // int32 | Number of fragments that are returned in highlights. If the search term occurs less often than the number of fragments that are specified, longer fragments are returned. (optional)
    highlightSize := int32(56) // int32 | Number of characters in each fragment for highlights. (optional)
    includeDocs := true // bool | Include the full content of the documents in the response. (optional)
    includeFields := "includeFields_example" // string | A JSON array of field names to include in search results. Any fields that are included must be indexed with the store:true option. limit (number) – Limit the number of the returned documents to the specified number. For a grouped search, this parameter limits the number of documents per group.  (optional)
    limit := int32(56) // int32 | Limit the number of the returned documents to the specified number. For a grouped search, this parameter limits the number of documents per group. (optional)
    q := "q_example" // string | Alias for query. (optional)
    query := "query_example" // string | Required. The Lucene query string. (optional)
    range_ := "range__example" // string | This field defines ranges for faceted, numeric search fields. The value is a JSON object where the fields names are faceted numeric search fields, and the values of the fields are JSON objects. The field names of the JSON objects are names for ranges. The values are strings that describe the range, for example “[0 TO 10]”. (optional)
    sort := "sort_example" // string | Specifies the sort order of the results. In a grouped search (when group_field is used), this parameter specifies the sort order within a group. The default sort order is relevance. A JSON string of the form \"fieldname<type>\" or -fieldname<type> for descending order, where fieldname is the name of a string or number field, and type is either a number, a string, or a JSON array of strings. The type part is optional, and defaults to number. Some examples are \"foo\", \"-foo\", \"bar<string>\", \"-foo<number>\" and [\"-foo<number>\", \"bar<string>\"]. String fields that are used for sorting must not be analyzed fields. Fields that are used for sorting must be indexed by the same indexer that is used for the search query. (optional)
    stale := "stale_example" // string | Set to ok to allow the use of an out-of-date index. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartitionApi.PartitionDesignDocSearch(context.Background(), db, partition, ddoc, index).Bookmark(bookmark).Counts(counts).Drilldown(drilldown).GroupField(groupField).GroupSort(groupSort).HighlightFields(highlightFields).HighlightPreTag(highlightPreTag).HighlightPostTag(highlightPostTag).HighlightNumber(highlightNumber).HighlightSize(highlightSize).IncludeDocs(includeDocs).IncludeFields(includeFields).Limit(limit).Q(q).Query(query).Range_(range_).Sort(sort).Stale(stale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionApi.PartitionDesignDocSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartitionDesignDocSearch`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `PartitionApi.PartitionDesignDocSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**partition** | **string** | Partition name | 
**ddoc** | **string** | Design document id | 
**index** | **string** | Search index name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartitionDesignDocSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bookmark** | **string** | A bookmark received from a previous search. This parameter enables paging through the results. If there are no more results after the bookmark, you get a response with an empty rows array and the same bookmark, confirming the end of the result list. | 
 **counts** | **string** | An array of names of string fields for which counts are requested. The response contains counts for each unique value of this field name among the documents that match the search query. Faceting must be enabled for this parameter to function. | 
 **drilldown** | **string** | This field can be used several times. Each use defines a pair with a field name and a value. The search matches only documents containing the value that was provided in the named field. It differs from using \&quot;fieldname:value\&quot; in the q parameter only in that the values are not analyzed. Faceting must be enabled for this parameter to function. | 
 **groupField** | **string** | Field by which to group search matches. :query number group_limit: Maximum group count. This field can be used only if group_field is specified.  | 
 **groupSort** | **string** | This field defines the order of the groups in a search that uses group_field. The default sort order is relevance. | 
 **highlightFields** | **string** | Specifies which fields to highlight. If specified, the result object contains a highlights field with an entry for each specified field. | 
 **highlightPreTag** | **string** | A string that is inserted before the highlighted word in the highlights output. | 
 **highlightPostTag** | **string** | A string that is inserted after the highlighted word in the highlights output. | 
 **highlightNumber** | **int32** | Number of fragments that are returned in highlights. If the search term occurs less often than the number of fragments that are specified, longer fragments are returned. | 
 **highlightSize** | **int32** | Number of characters in each fragment for highlights. | 
 **includeDocs** | **bool** | Include the full content of the documents in the response. | 
 **includeFields** | **string** | A JSON array of field names to include in search results. Any fields that are included must be indexed with the store:true option. limit (number) – Limit the number of the returned documents to the specified number. For a grouped search, this parameter limits the number of documents per group.  | 
 **limit** | **int32** | Limit the number of the returned documents to the specified number. For a grouped search, this parameter limits the number of documents per group. | 
 **q** | **string** | Alias for query. | 
 **query** | **string** | Required. The Lucene query string. | 
 **range_** | **string** | This field defines ranges for faceted, numeric search fields. The value is a JSON object where the fields names are faceted numeric search fields, and the values of the fields are JSON objects. The field names of the JSON objects are names for ranges. The values are strings that describe the range, for example “[0 TO 10]”. | 
 **sort** | **string** | Specifies the sort order of the results. In a grouped search (when group_field is used), this parameter specifies the sort order within a group. The default sort order is relevance. A JSON string of the form \&quot;fieldname&lt;type&gt;\&quot; or -fieldname&lt;type&gt; for descending order, where fieldname is the name of a string or number field, and type is either a number, a string, or a JSON array of strings. The type part is optional, and defaults to number. Some examples are \&quot;foo\&quot;, \&quot;-foo\&quot;, \&quot;bar&lt;string&gt;\&quot;, \&quot;-foo&lt;number&gt;\&quot; and [\&quot;-foo&lt;number&gt;\&quot;, \&quot;bar&lt;string&gt;\&quot;]. String fields that are used for sorting must not be analyzed fields. Fields that are used for sorting must be indexed by the same indexer that is used for the search query. | 
 **stale** | **string** | Set to ok to allow the use of an out-of-date index. | 

### Return type

[**Pagination**](Pagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartitionDesignDocView

> Pagination PartitionDesignDocView(ctx, db, partition, ddoc, view).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()

Executes the specified view function from the specified design document.

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
    partition := "partition_example" // string | Partition name
    ddoc := "ddoc_example" // string | Design document id
    view := "view_example" // string | View function name
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
    keys := "keys_example" // string | Return only documents where the key matches one of the keys specified in the array. (optional)
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
    resp, r, err := apiClient.PartitionApi.PartitionDesignDocView(context.Background(), db, partition, ddoc, view).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionApi.PartitionDesignDocView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartitionDesignDocView`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `PartitionApi.PartitionDesignDocView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**partition** | **string** | Partition name | 
**ddoc** | **string** | Design document id | 
**view** | **string** | View function name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartitionDesignDocViewRequest struct via the builder pattern


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
 **keys** | **string** | Return only documents where the key matches one of the keys specified in the array. | 
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
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartitionDesignDocViewPost

> Pagination PartitionDesignDocViewPost(ctx, db, partition, ddoc, view).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()

Executes the specified view function from the specified design document.

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
    partition := "partition_example" // string | Partition name
    ddoc := "ddoc_example" // string | Design document id
    view := "view_example" // string | View function name
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
    keys := "keys_example" // string | Return only documents where the key matches one of the keys specified in the array. (optional)
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
    resp, r, err := apiClient.PartitionApi.PartitionDesignDocViewPost(context.Background(), db, partition, ddoc, view).Body(body).Conflicts(conflicts).Descending(descending).Endkey(endkey).EndKey(endKey).EndkeyDocid(endkeyDocid).EndKeyDocId(endKeyDocId).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).RevsInfo(revsInfo).Key(key).Keys(keys).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).Startkey(startkey).StartKey(startKey).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionApi.PartitionDesignDocViewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartitionDesignDocViewPost`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `PartitionApi.PartitionDesignDocViewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**partition** | **string** | Partition name | 
**ddoc** | **string** | Design document id | 
**view** | **string** | View function name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartitionDesignDocViewPostRequest struct via the builder pattern


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
 **keys** | **string** | Return only documents where the key matches one of the keys specified in the array. | 
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

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartitionDocGetAll

> Pagination PartitionDocGetAll(ctx, db, partition).Conflicts(conflicts).Descending(descending).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()

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
    partition := "partition_example" // string | Partition name
    conflicts := true // bool | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. (optional)
    descending := true // bool | Return the documents in descending order by key. Default is false. (optional)
    group := true // bool | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. (optional)
    groupLevel := int32(56) // int32 | Specify the group level to be used. Implies group is true. (optional)
    includeDocs := true // bool | Include the associated document with each row. Default is false. (optional)
    attachments := true // bool | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. (optional)
    attEncodingInfo := true // bool | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. (optional)
    inclusiveEnd := true // bool | Specifies whether the specified end key should be included in the result. Default is true. (optional)
    limit := int32(56) // int32 | Limit the number of the returned documents to the specified number. (optional)
    reduce := true // bool | Use the reduction function. Default is true when a reduce function is defined. (optional)
    skip := int32(56) // int32 | Skip this number of records before starting to return the results. Default is 0. (optional)
    sorted := true // bool | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. (optional)
    stable := true // bool | Whether or not the view results should be returned from a stable set of shards. Default is false. (optional)
    stale := "stale_example" // string | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable=true&update=false. update_after is equivalent to stable=true&update=lazy. The default behavior is equivalent to stable=false&update=true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  (optional)
    startkeyDocid := "startkeyDocid_example" // string | Return records starting with the specified document ID. Ignored if startkey is not set. (optional)
    startKeyDocId := "startKeyDocId_example" // string | Alias for startkey_docid param (optional)
    update := "update_example" // string | Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.  (optional)
    updateSeq := true // bool | Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartitionApi.PartitionDocGetAll(context.Background(), db, partition).Conflicts(conflicts).Descending(descending).Group(group).GroupLevel(groupLevel).IncludeDocs(includeDocs).Attachments(attachments).AttEncodingInfo(attEncodingInfo).InclusiveEnd(inclusiveEnd).Limit(limit).Reduce(reduce).Skip(skip).Sorted(sorted).Stable(stable).Stale(stale).StartkeyDocid(startkeyDocid).StartKeyDocId(startKeyDocId).Update(update).UpdateSeq(updateSeq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionApi.PartitionDocGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartitionDocGetAll`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `PartitionApi.PartitionDocGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**partition** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartitionDocGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conflicts** | **bool** | Include conflicts information in response. Ignored if include_docs isn’t true. Default is false. | 
 **descending** | **bool** | Return the documents in descending order by key. Default is false. | 
 **group** | **bool** | Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false. | 
 **groupLevel** | **int32** | Specify the group level to be used. Implies group is true. | 
 **includeDocs** | **bool** | Include the associated document with each row. Default is false. | 
 **attachments** | **bool** | Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false. | 
 **attEncodingInfo** | **bool** | Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false. | 
 **inclusiveEnd** | **bool** | Specifies whether the specified end key should be included in the result. Default is true. | 
 **limit** | **int32** | Limit the number of the returned documents to the specified number. | 
 **reduce** | **bool** | Use the reduction function. Default is true when a reduce function is defined. | 
 **skip** | **int32** | Skip this number of records before starting to return the results. Default is 0. | 
 **sorted** | **bool** | Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true. | 
 **stable** | **bool** | Whether or not the view results should be returned from a stable set of shards. Default is false. | 
 **stale** | **string** | Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable&#x3D;true&amp;update&#x3D;false. update_after is equivalent to stable&#x3D;true&amp;update&#x3D;lazy. The default behavior is equivalent to stable&#x3D;false&amp;update&#x3D;true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  | 
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


## PartitionInfo

> Partition PartitionInfo(ctx, db, partition).Execute()

This endpoint returns information describing the provided partition. It includes document and deleted document counts along with external and active data sizes.

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
    partition := "partition_example" // string | Partition name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartitionApi.PartitionInfo(context.Background(), db, partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionApi.PartitionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartitionInfo`: Partition
    fmt.Fprintf(os.Stdout, "Response from `PartitionApi.PartitionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**partition** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartitionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Partition**](Partition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

