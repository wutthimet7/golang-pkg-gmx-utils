package utils

import (
	"github.com/gofiber/fiber/v2"
	"runtime"
	"strings"
)

var (
	TransactionId = "transactionId"
)

// ResponseBase for default error
type ResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetTransactionId(c *fiber.Ctx) string {
	transaction, ok := c.Locals(TransactionId).(string)
	if ok {
		return transaction
	} else {
		return ""
	}
}

func GetFuncName(temp interface{}) string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	return funcName[strings.LastIndex(funcName, ".")+1:]
}

func GetStringPointer(input string) (output *string) {
	output = &input
	return
}

func GetFloatPointer(input float64) (output *float64) {
	output = &input
	return
}
