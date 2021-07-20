package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//DeleteSshKeyResponse represents sshKey delete action response
type DeleteSshKeyResponse struct {
	Result   string `json:"result"`
	SshKeyID string `json:"sshKeyId"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *DeleteSshKeyResponse) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to DeleteSshKeyResponse dto:", err)
		os.Exit(1)
	}
}
