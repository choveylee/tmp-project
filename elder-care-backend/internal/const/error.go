package constant

import (
	"net/http"
)

var errorCodes = make(map[int]string)

func register(errCode int, errMsg string) int {
	if _, ok := errorCodes[errCode]; ok {
		panic("duplicate error code registration")
	}

	errorCodes[errCode] = errMsg

	return errCode
}

// ErrMsg returns the registered business error message for errCode.
func ErrMsg(errCode int) (string, bool) {
	errMsg, ok := errorCodes[errCode]

	return errMsg, ok
}

var (
	ErrorCodeOK = register(0, "")

	ErrorCodeMysqlServerAbnormal = register(100001, "Mysql服务器异常")
	ErrorCodeRedisServerAbnormal = register(100002, "Redis服务器异常")
	ErrorCodeHttpServerAbnormal  = register(100003, "Http服务器异常")

	ErrorCodeUnknownServerAbnormal = register(100011, "未知服务器异常")

	ErrorCodeRequestBodyInvalid  = register(100021, "请求Body非法")
	ErrorCodeRequestParamInvalid = register(100022, "请求参数非法")

	ErrorCodeAccessTokenInvalid = register(100031, "AccessToken非法")
	ErrorCodeAccessTokenExpired = register(100032, "AccessToken已过期")

	ErrorCodePermissionForbidden = register(100041, "权限禁止访问")

	ErrorCodeCaptchaInvalid = register(200001, "验证码非法")

	ErrorCodeRoleNotExist = register(200011, "角色不存在")

	ErrorCodePasswordRetryLimit = register(200021, "密码重试上限")
	ErrorCodePasswordNotMatch   = register(200022, "密码错误")

	ErrorCodeUserNotExist      = register(200031, "用户不存在")
	ErrorCodeUserStatusInvalid = register(200032, "用户状态非法")
	ErrorCodeUserMobileExist   = register(200033, "用户手机号已存在")
	ErrorCodeUserRoleInvalid   = register(200034, "用户角色非法")

	ErrorCodeBannerNotExist = register(200101, "轮播图不存在")

	ErrorCodeArticleCategoryNameExist = register(200111, "文章分类名称已存在")
	ErrorCodeArticleCategoryNotExist  = register(200112, "文章分类不存在")
	ErrorCodeArticleCategoryInUse     = register(200113, "文章分类正在使用中")

	ErrorCodeArticleNotExist = register(200121, "文章不存在")

	ErrorCodeCourseCategoryNameExist = register(200131, "课程分类名称已存在")
	ErrorCodeCourseCategoryNotExist  = register(200132, "课程分类不存在")
	ErrorCodeCourseCategoryInUse     = register(200133, "课程分类正在使用中")

	ErrorCodeCourseNotExist       = register(200211, "课程不存在")
	ErrorCodeCourseInvalid        = register(200212, "课程无效")
	ErrorCodeCourseTypeInvalid    = register(200216, "课程类型非法")
	ErrorCodeCourseDetailNotExist = register(200217, "课程详情不存在")

	ErrorCodeCourseVideoNotExist = register(200221, "课程视频不存在")
)

// StatusCode returns the HTTP status code mapped to errCode.
func StatusCode(errCode int) int {
	switch errCode {
	case ErrorCodeOK:
		return http.StatusOK

	case ErrorCodeMysqlServerAbnormal, ErrorCodeRedisServerAbnormal, ErrorCodeHttpServerAbnormal:
		return http.StatusInternalServerError

	case ErrorCodeUnknownServerAbnormal:
		return http.StatusInternalServerError

	case ErrorCodeRequestBodyInvalid, ErrorCodeRequestParamInvalid:
		return http.StatusBadRequest

	case ErrorCodeAccessTokenInvalid:
		return http.StatusUnauthorized

	case ErrorCodeAccessTokenExpired:
		return http.StatusUnauthorized

	case ErrorCodePermissionForbidden:
		return http.StatusForbidden

	case ErrorCodeCaptchaInvalid:
		return http.StatusBadRequest

	case ErrorCodeRoleNotExist:
		return http.StatusBadRequest

	case ErrorCodePasswordRetryLimit, ErrorCodePasswordNotMatch:
		return http.StatusBadRequest

	case ErrorCodeUserNotExist, ErrorCodeUserStatusInvalid, ErrorCodeUserMobileExist, ErrorCodeUserRoleInvalid:
		return http.StatusBadRequest

	case ErrorCodeBannerNotExist:
		return http.StatusBadRequest

	case ErrorCodeArticleCategoryNameExist, ErrorCodeArticleCategoryNotExist, ErrorCodeArticleCategoryInUse:
		return http.StatusBadRequest

	case ErrorCodeArticleNotExist:
		return http.StatusBadRequest

	case ErrorCodeCourseCategoryNameExist, ErrorCodeCourseCategoryNotExist, ErrorCodeCourseCategoryInUse:
		return http.StatusBadRequest

	case ErrorCodeCourseNotExist, ErrorCodeCourseInvalid, ErrorCodeCourseTypeInvalid, ErrorCodeCourseDetailNotExist:
		return http.StatusBadRequest

	case ErrorCodeCourseVideoNotExist:
		return http.StatusBadRequest

	default:
		panic("unrecognized error code")
	}
}
