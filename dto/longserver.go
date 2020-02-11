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
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Name               string   `header:"name"`
	Description        string   `header:"description"`
	PrivateIPAddresses []string `header:"Private Ips"`
	PublicIPAddresses  []string `header:"Public Ips"`
	Os                 string   `header:"os"`
	Type               string   `header:"type"`
	Location           string   `header:"location"`
	CPU                string   `header:"cpu"`
	RAM                string   `header:"ram"`
	Storage            string   `header:"storage"`
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
