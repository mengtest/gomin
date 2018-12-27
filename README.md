
+ 用户

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|login_name|varchar|登录名|
|name|varchar|姓名|
|mobile|varchar|手机号|
|email|varchar|邮箱|
|password|varchar|密码|
|status|tinyint|状态|
|is_valid|bit|是否有效


+ 角色表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|role_name|varchar|角色名称|
|remark|varchar|备注|
|status|tinyint|状态|
|is_valid|bit|是否有效|


+ 用户角色表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|user_id|int|用户id|
|role_id|int|角色id|
|status|tinyint|状态|
|is_valid|bit|是否有效|


+ 菜单表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|menu_name|varchar|菜单名称|
|menu_url|varchar|菜单url|
|remark|varchar|备注|
|status|tinyint|状态|
|is_valid|bit|是否有效|

+ 角色菜单表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|role_id|int|角色id|
|menu_id|int|菜单id|
|remark|varchar|备注|
|status|tinyint|状态|
|is_valid|bit|是否有效|

+ 权限表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|pms_code|varchar|权限代码|
|pms_name|varchar|权限名称
|remark|varchar|备注|
|status|tinyint|状态|
|is_valid|bit|是否有效|

+ 角色权限表

|字段|类型|描述|
|----|----|----|
|id|int|主键|
|create_time|int|创建时间|
|update_time|int|更新时间|
|role_id|int|角色id|
|pms_id|int|权限id|
|remark|varchar|备注|
|status|tinyint|状态|
|is_valid|bit|是否有效|