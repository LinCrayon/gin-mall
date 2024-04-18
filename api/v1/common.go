package v1

import (
	"encoding/json"
	"errors"
	"gin-mall/serializer"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) serializer.Response {

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
