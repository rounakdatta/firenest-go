package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadMessageFromFile(fileName string) string {
	filePath := fmt.Sprintf("../resources/test/%s", fileName)
	message, err := ioutil.ReadFile(filePath)
	if err == nil {
		return string(message)
	}

	return ""
}
