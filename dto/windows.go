package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Windows
type Windows struct {
	RdpAllowedIps []string `json:"rdpAllowedIps"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto Windows) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("Windows dto can not be converted to io.Reader:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
