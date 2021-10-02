package utils

import (
	"fmt"
	"io/ioutil"
)

type Direction uint32

const (
	DEBIT  Direction = 0
	CREDIT Direction = 1
	OWN    Direction = 2
)

func ReadMessageFromFile(fileName string) string {
	filePath := fmt.Sprintf("../resources/test/%s", fileName)
	message, err := ioutil.ReadFile(filePath)
	if err == nil {
		return string(message)
	}

	return ""
}
