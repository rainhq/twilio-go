# TaskrouterV1WorkspaceWorkspaceRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActivityStatistics** | Pointer to **[]map[string]interface{}** | The number of current Workers by Activity |
**LongestTaskWaitingAge** | Pointer to **int32** | The age of the longest waiting Task |
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task |
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status |
**TotalTasks** | Pointer to **int32** | The total number of Tasks |
**TotalWorkers** | Pointer to **int32** | The total number of Workers in the Workspace |
**Url** | Pointer to **string** | The absolute URL of the Workspace statistics resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

