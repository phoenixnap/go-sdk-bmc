package main

import (
	"fmt"
	"os"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/command"
	"github.com/PNAP/bmc-api-sdk/dto"
)

func main() {

	client, confError := client.Create()
	if confError != nil {
		fmt.Println("Error reading config file:", confError)
		os.Exit(3)
	}
	requestCommand := &command.GetServersCommand{}
	requestCommand.SetRequester(client)
	resp, err := requestCommand.Execute()
	response := &dto.Servers{}
	if err == nil {
		response.FromBytes(resp)
		fmt.Println("Done!")
	}

}
