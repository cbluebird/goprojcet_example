# 这是一个接口说明文档
> ## 绝对不能让用用戶輸入id，会引起十分严重的错误

> ## 1. login接口
**路由接口**：`user/login`

**接口方法**：post

| 参数        | 类型 | 值   |
|-----------| --- |-----|
| user_name |string| 用户名 |
| password  |string | 密码  |

传入示例：
````json
{
    "user_name": "刘凌赫",
    "password": "1234"
}
````
返回值

|code| msg         |说明|
|---|-------------|---|
|200| ok          |操作成功|
|404| User not found |用户找不到|
|400||连接错误|
|400|Wrong Password|密码错误|

如果200,还会返回data，里面包含了user_id
返回示例：
```json
{
   "data": {
      "user_id": 1
   },
   "msg": "OK"
}

```
> # 2. register接口
**路由接口**：`user/register`

**接口方法**：post

|参数|类型|说明|
|---|---|---|
|user_name|string|用户名|
|user_sex|string|用户性别|
|password|string|密码|
|user_email|string|用户邮箱|

传入示例：
````json
{
    "user_sex": "女",
    "password": "123456789",
    "user_email": "llh@163.com",
    "user_name": "刘凌赫"
}
````

|code| msg         | 说明    |
|---|-------------|-------|
|200| ok          | 操作成功  |
|404| username repeat | 用户名重复 |
|400||连接错误|
|500||网络错误|

返回示例：
````json
{
    "msg": "OK"
}
````

> ## 3. chang user information接口
**路由接口**：`user/information`

**接口方法**：post


|参数|类型|说明|
|---|---|---|
|user_name|string|用户名|
|user_sex|string|用户性别|
|password|string|密码|
|user_email|string|用户邮箱|
|user_id|int|1|

````json
{
"user_sex":"男",
"password":"12345",
"user_email":"llh38@163.com",
"user_name":"刘凌赫",
"user_id":1
}
````

返回参数：

|code| msg         | 说明    |
|---|-------------|-------|
|200| ok        | 操作成功  |
|500||网络错误|

````json
{
"msg": "OK"
}
````

> ## 4. 查询用户历史简历接口
**路由接口**：`user/read`

**接口方法**：get

| 参数      | 类型  | 值    |
|---------|-----|------|
| uesr_id | int | 账户编号 |

添加在query参数里进行请求

返回值：

| 字段   |类型|说明|
|------|---|---|
| code |string|状态码|
| data |object|对象|
|list|arrey|包含在对象中的数组|
|number|int|简历数量|

list包含的参数：（注意大小写）

| 字段       |类型|说明|
|----------|---|---|
| Title    |string|标题|
| Date     |string|时间|
| Resumeid |int|简历编号|

返回示例
````json
{
    "data": {
        "list": [
            {
                "Title": "测试",
                "Date": "2022_09_28",
                "Resumeid": 1
            }
        ],
        "number": 1
    },
    "msg": "OK"
}
````


> ## 5. 新建简历接口
**路由接口**`resume/add`

**接口方法**：post

| 参数      | 类型  | 值    |
|---------|-----|------|
| uesr_id | int | 账户编号 |
|title|string|简历标题|

````json
{
    "user_id": 1,
    "title": "简历1"
}
````

返回参数：

|code| msg         | 说明    |
|---|-------------|-------|
|200| ok        | 操作成功  |
|500||网络错误|
|404|resume name repeat|用户名重复|

示例：

````json
{
  "data": {
  "resume_id": 1
  },
  "msg": "OK"
}
````

> ##  6. 修改简历


**路由接口**`resume/template`

**接口方法**：post

|字段|类型|说明|是否必要|
|---|---|---|---|
|name|string|用户姓名|否|
|phone|string|用户电话|否|
|user_email|String|用户邮箱|否|
|graduate_data|String|毕业时间|否|
|education_background|String|教育背景|否|
|previous_jobs|String|历史工作|否|
|work_experience|String|工作经历|否|
|self_introduce|String|自我介绍|否|
|personal_skills|string|个人技能|否|
|user_id|number|用户id|是|
|resume_id|number|简历id|是|
|router|string|ly|是|

示例：

````json
{
    "user_id": 1,
    "resume_id": 1,
    "title": "简历1",
    "name": "倪佳旗",
    "phone": "19841184447",
    "user_email": "njq@qq.com",
    "education_background": "北大博士",
    "graduate_data": "2002-09-30 19:37:59",
    "previous_jobs": "学生",
    "work_experience": "无",
    "self_introduce": "我叫njq，来自浙江，擅长。。。",
    "personal_skills": "无"
}
````

返回值：

|code| msg         | 说明    |
|---|-------------|-------|
|200| ok        | 操作成功  |
|500||网络错误|
|404|resume not found| 找不到简历 |


>##  7. 读取简历


**路由接口**`resume/read`

**接口方法**：get

| 参数      | 类型  | 值    |
|---------|-----|------|
| uesr_id | int | 账户编号 |
| id      |string| 简历编号 |

query进行请求

返回参数：

|code| msg         | 说明    |
|---|-------------|-------|
|200| ok        | 操作成功  |
|404|resume not found| 找不到简历 |

data:

|字段|类型|说明|
|---|---|---|
|name|string|用户姓名|
|phone|string|用户电话|
|user_email|String|用户邮箱|
|graduate_data|String|毕业时间|
|education_background|String|教育背景|
|previous_jobs|String|历史工作|
|work_experience|String|工作经历|
|self_introduce|String|自我介绍|
|personal_skills|string|个人技能|
|user_id|number|用户id|
|resume_id|number|简历id|
|creation_time|string|创建时间|
|router|string|路由|

````json
{
"data": {
"creation_time": "2002-09-30 19:37:59",
"education_background": "北大博士",
"graduate_data": "2002-09-30 19:37:59",
"name": "倪佳旗",
"personal_skills": "计算机",
"phone": "19841184447",
"previous_jobs": "后端工程师",
"resume_id": 1,
"self_introduce": "我叫llh，哈哈哈哈",
"user_email": "u.qlhjdqs@qq.com",
"user_id": 1,
"work_experience": "阿里"
},
"msg": "OK"
}
````

>##  8. 修改密码

**路由接口**：`user/repassword`

**接口方法**：post

| 参数         | 类型  | 值    |
|------------|-----|------|
| user_email | int | 账户编号 |
| password   |string| 密码   |

返回参数：

|code| msg         | 说明   |
|---|-------------|------|
|200| ok        | 操作成功 |
|404|email not found| 找不到简历|
|500|网络错误||



>##  9. 邮箱验证

**路由接口**：`user/email`

**接口方法**：post


| 参数         | 类型  | 值    |
|------------|-----|------|
| user_email | int | 账户邮箱 |

返回参数：

| 参数  | 类型      | 值     |
|-----|---------|-------|
| msg | string  | 消息    |
|code|int| 状态码   |
|data|对象| 包含验证码 |

|code| msg         | 说明   |
|---|-------------|------|
|200| ok        | 操作成功 |
|500|网络错误||

















