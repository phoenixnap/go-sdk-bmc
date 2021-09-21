package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Esxi
type Esxi struct {
	RootPassword string `json:"rootPassword,omitempty"`
	ManagementUiUrl string `json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps []string `json:"managementAccessAllowedIps,omitempty"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto Esxi) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("Esxi dto can not be converted to io.Reader:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
