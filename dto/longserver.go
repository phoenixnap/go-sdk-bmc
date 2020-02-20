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

//LongServer represents server
type LongServer struct {
	ID                 string   `json:"id"`
	Status             string   `json:"status"`
	Name               string   `json:"hostname"`
	Description        string   `json:"description"`
	PrivateIPAddresses []string `json:"privateIpAddresses"`
	PublicIPAddresses  []string `json:"publicIpAddresses"`
	Os                 string   `json:"os"`
	Type               string   `json:"type"`
	Location           string   `json:"location"`
	CPU                string   `json:"cpu"`
	RAM                string   `json:"ram"`
	Storage            string   `json:"storage"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto LongServer) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("LongServer dto can not be converted to io.Reader:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}

//FromBytes performs conversion of http response to the representing struct
func (dto *LongServer) FromBytes(resp *http.Response) {
	//construct := &LongServer{}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
		//newDTO. = *construct
	} else {
		fmt.Println("API response body can not be converted to LongServer dto:", err)
		os.Exit(1)
	}
}
