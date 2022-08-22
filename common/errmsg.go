package common

import (
	"sync"
)

var mapErrMsg map[int32]string
var once sync.Once

func InitErrMsgMap() {
	once.Do(func() {
		mapErrMsg = make(map[int32]string)
		mapErrMsg[CodeSuccess] = "成功"
		mapErrMsg[CodeFailed] = "失败"
		mapErrMsg[CodeBusy] = "服务忙，请稍后再试"
		mapErrMsg[CodeInternalErr] = "系统错误，请稍后再试"
		mapErrMsg[CodeUnknown] = "未知错误，错误码无说明"
		mapErrMsg[CodeDBErr] = "数据库操作失败"
		mapErrMsg[CodeInvalidParam] = "不合法的参数"
		mapErrMsg[CodeErrorConvertParam] = "参数类型转换异常"
		mapErrMsg[CodeNotPermission] = "没有权限"
		mapErrMsg[CodeNetErr] = "host连接超时"
	})
}

// ErrCodeToMsg 转换code到错误说明
func ErrCodeToMsg(errCode int32) string {
	if len(mapErrMsg) == 0 {
		InitErrMsgMap()
	}
	strMsg, ok := mapErrMsg[errCode]
	if !ok {
		strMsg = mapErrMsg[CodeUnknown]
	}

	return strMsg
}
