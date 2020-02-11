package main

import (
	"fmt"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/command"
	"github.com/PNAP/bmc-api-sdk/dto"
)

func main() {

	client := client.Create()
	requestCommand := &command.GetServersCommand{}
	requestCommand.SetRequester(client)
	resp, err := requestCommand.Execute()
	response := &dto.Servers{}
	if err == nil {
		response.FromBytes(resp)
		fmt.Println("Done!")
	}

}
