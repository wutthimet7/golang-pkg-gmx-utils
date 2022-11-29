package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// GetErrorCode for get error code
// -99 status code can not convert from string to int
func GetErrorCode(e error) int {
	s := strings.Split(e.Error(), ", ")
	for _, x := range s {
		q := strings.Split(x, "code=")
		if len(q) > 1 {
			i, err := strconv.Atoi(q[1])
			if err != nil {
				return -99
			}
			return i
		}
	}
	return -99
}

// GetErrorMessage for get error message
func GetErrorMessage(e error) string {
	s := strings.Split(e.Error(), ", ")
	for _, x := range s {
		q := strings.Split(x, "message=")
		if len(q) > 1 {
			return q[1]
		}
	}
	return e.Error()
}

const (
	Required string = "required"
	Min      string = "min"
	Max      string = "max"
	Length   string = "len"
	Numeric  string = "numeric"
)

func ValidateInputs(dataSet interface{}) (resp string) {
	var values []string
	validate := validator.New()
	err := validate.Struct(dataSet)
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			var message string
			switch vErr.Tag() {
			case Required:
				message = "%s is required.%s"
			case Min:
				message = "%s is less than %s."
			case Max:
				message = "%s is not over %s."
			case Length:
				message = "%s is length more than %s."
			case Numeric:
				message = "%s is not numeric.%s"
			}

			values = append(values, fmt.Sprintf(message, vErr.Field(), vErr.Param()))
		}
		resp = strings.Join(values, ", ")
	}
	return resp
}
