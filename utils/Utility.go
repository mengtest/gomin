package utils

type Response struct {
	status int
	msg    string
	count  int64
	data   interface{}
}

func Success(data interface{}) *map[string]interface{} {
	return &map[string]interface{}{"code": SUCCESS, "msg": "succss", "data": data}
}

func ListSuccess(count int64, listData interface{}) *map[string]interface{} {
	return &map[string]interface{}{"code": SUCCESS, "msg": "succss", "data": listData, "count": count}

}

func ParamError() *map[string]interface{} {
	return &map[string]interface{}{"code": SYS_PARAM_ERROR, "msg": "param is invalid"}
}

func SysError() *map[string]interface{} {
	return &map[string]interface{}{"code": SYS_ERROR, "msg": "sys error"}
}
