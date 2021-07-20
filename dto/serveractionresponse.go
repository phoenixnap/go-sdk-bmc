package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Result of a successful server action.
type ServerActionResponse struct {
	Result   string `json:"result"`
	Password string `json:"password"`
	ServerID string `json:"serverId"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ServerActionResponse) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to ServerActionResponse dto:", err)
		os.Exit(1)
	}
}
