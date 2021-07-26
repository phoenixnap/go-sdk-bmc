package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//SshKey represents sshkey
type SshKey struct {
	ID            string `json:"id"`
	Default       bool   `json:"default"`
	Name          string `json:"name"`
	Key           string `json:"key"`
	Fingerprint   string `json:"fingerprint"`
	CreatedOn     string `json:"createdOn"`
	LastUpdatedOn string `jason:"lastUpdatedOn"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *SshKey) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		//If body contains valid JSON that fits in SshKey, after the call err will be nil and the data from body will have been stored in the struct SshKey ()
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to SshKey dto:", err)
		os.Exit(1)
	}
}

//ToBytes performs conversion of struct to the io.Reader
func (dto SshKey) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("SshKey details entered in invalid format:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
