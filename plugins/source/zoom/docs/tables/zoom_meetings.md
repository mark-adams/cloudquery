
# Table: zoom_meetings
Scheduled Zoom meetings.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|uuid|text|Unique Meeting ID.|
|id|bigint|Meeting ID - also known as the meeting number in long (int64) format.|
|host_id|text|ID of the user who is set as the host of the meeting.|
|topic|text|Meeting topic.|
|type|bigint|Meeting type.|
|status|text|Meeting status.|
|start_time|timestamp without time zone|Meeting start time.|
|duration|bigint|Meeting duration.|
|timezone|text|Timezone to format the meeting start time.|
|created_at|timestamp without time zone|Time of creation.|
|agenda|text|Meeting description.|
|start_url|text|The URL using which a host or an alternative host can start the Meeting.|
|join_url|text|URL for participants to join the meeting.|
|password|text|Meeting passcode.|
|h323_password|text|H.323/SIP room system passcode.|
|encrypted_password|text|Encrypted passcode for third party endpoints (H323/SIP).|
|personal_meeting_id|text|Personal Meeting ID|
|tracking_fields|jsonb||
|occurrences|jsonb|Array of occurrence objects.|
|settings|jsonb|Meeting settings.|
|recurrence|jsonb|Recurrence object.|
