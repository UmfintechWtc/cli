package common

import "encoding/json"

type XDeliverException struct {
	Code             string
	Code_description string
	Error_message    interface{}
}

func (X *XDeliverException) Error() string {
	err, _ := json.Marshal(X)
	return string(err)
}

func Rsperror(code string, error_message ...interface{}) *XDeliverException {
	return &XDeliverException{
		Code:             code,
		Code_description: Code_description_definbtion[code],
		Error_message:    error_message,
	}
}
