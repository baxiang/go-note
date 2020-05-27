package errcode

type ErrCode int

//go:generate stringer -type ErrCode -linecomment
const (
	 ERR_CODE_OK ErrCode=0 //OK
	 ERR_CODE_INVALID_PARAMS ErrCode=1 //参数错误
	 ERR_CODE_TIMEOUT ErrCode =2 //请求超时

)