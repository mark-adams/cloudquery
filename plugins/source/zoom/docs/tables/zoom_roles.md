
# Table: zoom_roles
Role represents an account role.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|Role ID.|
|name|text|Name of the role.|
|description|text|Description of the role.|
|total_members|bigint|Total members assigned to that role.|
|privileges|text[]|Privileges assigned to the role.|
|sub_account_privileges|jsonb|This field will only be displayed to accounts that are enrolled in a partner plan and follow the master accounts and sub accounts structure.|
