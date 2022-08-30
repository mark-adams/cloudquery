
# Table: zoom_cloud_recordings
Zoom meting recordings.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|uuid|text|Unique Meeting ID.|
|id|bigint|Meeting ID - also known as the meeting number.|
|account_id|text|Unique Identifier of the user account.|
|host_id|text|ID of the user set as host of meeting.|
|topic|text|Meeting topic.|
|type|bigint|The recording's associated type of meeting or webinar.|
|start_time|timestamp without time zone|The time at which the meeting started.|
|duration|bigint|Meeting duration.|
|total_size|bigint|The total file size of the recording.|
|recording_count|bigint|Number of recording files.|
