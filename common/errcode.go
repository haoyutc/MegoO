package common

// 全局错误码
const (
	CodeSuccess     = 0
	CodeFailed      = 1
	CodeBusy        = -1
	CodeInternalErr = -2 // 内部错误，内部调试用，不要暴露到外面！！
	CodeUnknown     = -3 // 未知错误
	CodeDBErr       = -4 // 数据库错误
	CodeNetErr      = -5 // host连接超时

	CodeNotPermission     = 10000 // 接口无权限调用
	CodeInvalidParam      = 10001 // 不合法的参数
	CodeErrorConvertParam = 10002 // 参数类型转换异常
)
