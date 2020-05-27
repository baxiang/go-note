package err

import (
	"fmt"
)

type ErrCode int

const (
	ERR_CODE_OK             ErrCode = 0 // success
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 参数无效
	ERR_CODE_TIMEOUT        ErrCode = 2 // 请求超时
	// ...
)

// 定义错误码与描述信息的映射
var mapErrDesc = map[ErrCode]string{
	ERR_CODE_OK:             "success",
	ERR_CODE_INVALID_PARAMS: "参数无效",
	ERR_CODE_TIMEOUT:        "请求超时",
	// ...
}

func getErrDescByCode(code ErrCode) string {
	if desc, ok := mapErrDesc[code]; ok {
		return desc
	}
	return fmt.Sprintf("未知错误：%d", code)
}

func (err ErrCode) String() string {
	return getErrDescByCode(err)
}
