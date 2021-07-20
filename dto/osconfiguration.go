package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//OsConfiguration
type OsConfiguration struct {
	Windows Windows `json:"windows"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto OsConfiguration) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("OsConfiguration dto can not be converted to io.Reader:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
