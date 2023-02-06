package api_response

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid parameters",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "auth check token failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "auth check token timeout",
	ERROR_AUTH_TOKEN:               "error auth token",
	ERROR_AUTH:                     "error auth",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
