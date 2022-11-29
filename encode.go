package utils

import (
	"encoding/base64"
	"fmt"
)

func EncodeBase64(input string) string {

	output := base64.StdEncoding.EncodeToString([]byte(input))

	return output
}

func DecodeBase64(input string) (string, error) {

	output, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(output), nil
}
