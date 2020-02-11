package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Servers represents list of servers
type Servers []LongServer

//FromBytes performs conversion of http response to the representing struct
func (dto *Servers) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to LongServers dto:", err)
		os.Exit(1)
	}
}
