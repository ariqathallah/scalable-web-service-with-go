package util

import "assignment-2/model"

func CreateResponse(success bool, data any, err string) model.Response {
	return model.Response{
		Success: success,
		Data:    data,
		Error:   err,
	}
}
