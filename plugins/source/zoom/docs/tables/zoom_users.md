
# Table: zoom_users
User represents an account user.s
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The user's ID.|
|first_name|text|The user's first name.|
|last_name|text|The user's last name.|
|email|text|The user's email address.|
|type|bigint|The user's assigned plan type.|
|status|text|The user's status.|
|personal_meeting_id|bigint|The user's Personal Meeting ID.|
|timezone|text|The user's timezone.|
|dept|text|The user's department.|
|role_id|text|The unique ID of the user's assigned role.|
|created_at|timestamp without time zone|The time at which the user's account was created.|
|last_login_time|timestamp without time zone|The user's last login time.|
|last_client_version|text|The last client version that user used to log in.|
|group_ids|text[]|The IDs of groups where the user is a member.|
|im_group_ids|text[]|The IDs of IM directory groups where the user is a member.|
|verified|bigint|Whether the user's email address for the Zoom account is verified.|
|custom_attributes|jsonb|Information about the user's custom attributes.|
|plan_united_type|text|The user's Zoom United plan.|
|use_pmi|boolean||
|language|text|Default language for the Zoom Web Portal.|
|phone_numbers|jsonb|The user's phone numbers.|
|vanity_url|text|Personal meeting room URL, if the user has one.|
|personal_meeting_url|text|User's personal meeting url.|
|pic_url|text|The URL for user's profile picture.|
|cms_user_id|text|CMS ID of the user.|
|account_id|text|User's account ID.|
|jid|text|The user's Jabber ID (JID).|
|job_title|text|The user's job title.|
|company|text|The user's company.|
|location|text|User's location.|
|login_type|bigint|The user's login method.|
