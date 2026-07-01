# elder-care-backend YApi接口文档

生成来源：`internal/router`、`internal/handler`、`internal/data`。

可导入文件：[yapi-swagger.json](/Users/choveylee/Documents/choveylee/tmp-project/elder-care-backend/docs/yapi-swagger.json)

## 通用说明

- 成功响应统一为：`{"code":0,"data":...}`。无返回数据的写操作返回：`{"code":0}`。
- 失败响应统一为：`{"code":错误码,"message":"错误信息","detail":"错误详情"}`。
- 管理端鉴权接口需要请求头：`Authorization: Bearer <access_token>`。
- 分页参数：`page_num` 默认 `1`，范围 `1-100`；`page_size` 默认 `10`，范围 `1-100`。
- 时间字段请求格式按代码校验为 RFC3339，例如 `2026-06-30T10:00:00+08:00`。
- 当前工作区未包含独立前端工程；本文将 `/api/v1/public` 视为前端公开端接口，`/api/v1/admin` 视为管理端接口。

## 健康检查

| 方法 | 路径 | 名称 | 鉴权 |
|---|---|---|---|
| GET | `/cpu-check` | CPU健康检查 | 否 |
| GET | `/ram-check` | 内存健康检查 | 否 |

响应：`status`、`detail`。

## 公开端接口

| 方法 | 路径 | 名称 | 参数 |
|---|---|---|---|
| POST | `/api/v1/public/captcha` | 创建验证码 | 无 |
| GET | `/api/v1/public/banners` | 轮播图列表 | 无 |
| GET | `/api/v1/public/articles/categories` | 文章分类列表 | 无 |
| GET | `/api/v1/public/articles` | 文章列表 | `category_id` 必填，`page_num`，`page_size` |
| GET | `/api/v1/public/articles/{id}` | 文章详情 | `id` |
| GET | `/api/v1/public/courses/categories` | 课程分类列表 | 无 |
| GET | `/api/v1/public/courses` | 课程列表 | `category_id` 必填，`course_type`，`sort_by`，`page_num`，`page_size` |
| GET | `/api/v1/public/courses/{id}` | 课程详情 | `id` |
| GET | `/api/v1/public/courses/{id}/catalogs` | 课程目录列表 | `id`，响应含 `catalogs[].video` |

公开端枚举：

- `course_type`：`0` 普通课程，`1` 视频课程。
- `sort_by`：`publish` 发布时间，`view` 浏览量，`favourite` 收藏量。

## 管理端通用与登录

| 方法 | 路径 | 名称 | 鉴权 | 参数 |
|---|---|---|---|---|
| POST | `/api/v1/admin/captcha` | 创建验证码 | 否 | 无 |
| POST | `/api/v1/admin/sessions` | 登录 | 否 | `mobile`、`password`、`captcha_id`、`captcha_code` |
| POST | `/api/v1/admin/image` | 上传图片 | 是 | multipart `file`，支持 `.img/.png/.jpg/.jpeg` |

## 管理端用户

| 方法 | 路径 | 名称 | 参数 |
|---|---|---|---|
| GET | `/api/v1/admin/users` | 用户列表 | `role_id`，`status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/users` | 创建用户 | `role_id`、`name`、`mobile`、`password`、`status` |
| PUT | `/api/v1/admin/users/{id}` | 编辑用户 | `id`，body: `role_id`、`name`、`mobile`、`status` |
| PUT | `/api/v1/admin/users/{id}/password` | 重置用户密码 | `id`，body: `password` |
| PUT | `/api/v1/admin/users/password` | 修改当前用户密码 | `old_password`、`new_password` |

## 管理端轮播图

| 方法 | 路径 | 名称 | 参数 |
|---|---|---|---|
| GET | `/api/v1/admin/banners` | 轮播图列表 | `status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/banners` | 创建轮播图 | `title`、`abstract`、`image_url`、`link_url`、`weight`、`status` |
| GET | `/api/v1/admin/banners/{id}` | 轮播图详情 | `id` |
| PUT | `/api/v1/admin/banners/{id}` | 编辑轮播图 | `id`，同创建字段 |
| DELETE | `/api/v1/admin/banners/{id}` | 删除轮播图 | `id` |

## 管理端文章

| 方法 | 路径 | 名称 | 参数 |
|---|---|---|---|
| GET | `/api/v1/admin/articles/categories` | 文章分类列表 | `status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/articles/categories` | 创建文章分类 | `name`、`weight`、`status` |
| GET | `/api/v1/admin/articles/categories/{id}` | 文章分类详情 | `id` |
| PUT | `/api/v1/admin/articles/categories/{id}` | 编辑文章分类 | `id`，body: `name`、`weight`、`status` |
| DELETE | `/api/v1/admin/articles/categories/{id}` | 删除文章分类 | `id` |
| GET | `/api/v1/admin/articles` | 文章列表 | `category_id`，`status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/articles` | 创建文章 | `category_id`、`title`、`abstract`、`content`、`cover_url`、`link_url`、`publish_at`、`status` |
| GET | `/api/v1/admin/articles/{id}` | 文章详情 | `id` |
| PUT | `/api/v1/admin/articles/{id}` | 编辑文章 | `id`，body: `title`、`abstract`、`content`、`cover_url`、`link_url`、`publish_at`、`status` |
| DELETE | `/api/v1/admin/articles/{id}` | 删除文章 | `id` |

## 管理端课程

| 方法 | 路径 | 名称 | 参数 |
|---|---|---|---|
| GET | `/api/v1/admin/courses/categories` | 课程分类列表 | `status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/courses/categories` | 创建课程分类 | `name`、`weight`、`status` |
| GET | `/api/v1/admin/courses/categories/{id}` | 课程分类详情 | `id` |
| PUT | `/api/v1/admin/courses/categories/{id}` | 编辑课程分类 | `id`，body: `name`、`weight`、`status` |
| DELETE | `/api/v1/admin/courses/categories/{id}` | 删除课程分类 | `id` |
| GET | `/api/v1/admin/courses` | 课程列表 | `category_id`，`course_type`，`status`，`page_num`，`page_size` |
| POST | `/api/v1/admin/courses` | 创建课程 | `category_id`、`course_type`、`author`、`source`、`title`、`tags`、`abstract`、`cover_url`、`link_url`、`detail`、`summary`、`objective`、`outline`、`references`、`publish_at`、`status` |
| GET | `/api/v1/admin/courses/{id}` | 课程详情 | `id` |
| PUT | `/api/v1/admin/courses/{id}` | 编辑课程 | `id`，body 同创建课程但不含 `course_type` |
| DELETE | `/api/v1/admin/courses/{id}` | 删除课程 | `id` |
| GET | `/api/v1/admin/courses/{id}/catalogs` | 课程目录列表 | `id`，响应含 `catalogs[].video` |
| POST | `/api/v1/admin/courses/{id}/catalogs` | 创建课程目录 | `id`，body: `parent_id`、`name`、`weight`、`status`、`video`；`video` 不含 `weight/status` |
| PUT | `/api/v1/admin/courses/catalogs/{id}` | 编辑课程目录 | `id`，body 同创建课程目录，增删改目录时同步处理 `video` 信息 |
| DELETE | `/api/v1/admin/courses/catalogs/{id}` | 删除课程目录 | `id`，同时删除对应视频 |

管理端枚举：

- 轮播图、文章、文章分类、课程分类 `status`：`1` 正常，`0` 禁用。
- 课程与课程目录 `status`：`0` 正常，`1` 禁用。
- 课程 `course_type`：`0` 普通课程，`1` 视频课程。

