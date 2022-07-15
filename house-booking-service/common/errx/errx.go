package errx

import "fmt"

type CommonError struct {
	errCode uint32
	errMsg  string
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("ErrCode: %d,ErrMsg:%s", e.errCode, e.errMsg)
}

func (e *CommonError) GetErrCode() uint32 {
	return e.errCode
}

func (e *CommonError) GetErrMsg() string {
	return e.errMsg
}

func NewCommonError(errCode uint32, errMsg string) *CommonError {
	return &CommonError{
		errCode: errCode,
		errMsg:  errMsg,
	}
}

func NewErrMsg(errMsg string) *CommonError {
	return &CommonError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CommonError {
	return &CommonError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}
