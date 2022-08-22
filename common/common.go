package common

import (
	"context"
	"github.com/megoo/utils"
	"strconv"
)

// CommonRsp 通用的Rsp Json结构
type CommonRsp struct {
	Code      int32  `json:"Code"`
	Msg       string `json:"Msg"`
	RequestId string `json:"RequestId"`
}

// SetCommonRsp 设置common rsp
func SetCommonRsp(code int32, ctx ...context.Context) CommonRsp {

	var requestId string
	if ctx != nil && ctx[0] != nil {
		requestId = ctx[0].Value("X-Request-Id").(string)
	} else {
		requestId = ""
	}
	if requestId == "" {
		uuid := utils.GenerateUID()
		requestId = strconv.FormatInt(uuid, 10)
	}

	return CommonRsp{
		Code:      code,
		Msg:       ErrCodeToMsg(code),
		RequestId: requestId,
	}
}

// CgiError 错误类型
type CgiError struct {
	Code  int32
	Error error
}

// CgiPanic ..
func CgiPanic(code int32, err error) {
	panic(CgiError{
		Code:  code,
		Error: err,
	})
}

type CgiCommonRsp struct {
	CommonRsp
}
