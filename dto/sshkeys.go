package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//SshKeys represents list of sshKeys
type SshKeys []SshKey

func (dto *SshKeys) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to SshKeys dto:", err)
		os.Exit(1)
	}
}
