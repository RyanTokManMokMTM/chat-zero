package errx

var errorMessage map[uint32]string

func init() {
	errorMessage := make(map[uint32]string)
	errorMessage[OK] = "SUCCESS"
	errorMessage[SERVER_COMMON_ERROR] = "SERVER INTERNAL ERROR"
	errorMessage[REQ_PARAM_ERROR] = "REQUEST PARAMETER ERROR"
	errorMessage[TOKEN_EXPIRED_ERROR] = "TOKEN HAS BEEN EXPIRED"
	errorMessage[TOKEN_INVALID_ERROR] = "TOKEN HAS BEEN Invalid"
	errorMessage[TOKEN_GENERATE_ERROR] = "TOKEN GENERATE FAILED"
	errorMessage[DB_ERROR] = "DATABASE ERROR"
	errorMessage[DB_UPDATE_AFFECTED_ZERO_ERROR] = "DATABASE AFFECTED 0 rows"
	errorMessage[USER_HAS_BEEN_REGISTERED] = "USER HAS BEEN REGISTERED"
	errorMessage[USER_NOT_EXIST] = "USER NOT EXIST"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := errorMessage[errCode]; ok {
		return msg
	}
	return "SERVER INTERNAL ERROR"
}

func IsCodeError(errCode uint32) bool {
	if _, ok := errorMessage[errCode]; ok {
		return true
	}
	return false
}
