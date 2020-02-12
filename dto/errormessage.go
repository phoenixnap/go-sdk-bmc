package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//ErrorMessage represents API response error messages.
type ErrorMessage struct {
	Message          string `json:"message"`
	ValidationErrors string `json:"validationErrors"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ErrorMessage) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to ErrorMessage dto:", err)
		os.Exit(1)
	}
}
