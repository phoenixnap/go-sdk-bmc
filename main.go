package main

import (
	"fmt"
	//"os"

	//"github.com/PNAP/bmc-api-sdk/client"
	newClient "github.com/phoenixnap/go-sdk-bmc/client/pnapClient"
	"github.com/phoenixnap/go-sdk-bmc/command"
	"github.com/phoenixnap/go-sdk-bmc/dto"
)

func main() {

	auth := dto.Authentication{ClientID : "",
	ClientSecret: "",
	TokenURL: "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token",
	ApiHostName:"https://api-dev.phoenixnap.com/",
	PoweredBy: "go-sdk"}
	cl := newClient.NewPNAPClient(auth)
	cl.SetAuthentication(auth)


	requestCommand0 := &command.CreateServerCommand{}
	requestCommand0.SetRequester(cl)
	server0 := dto.ProvisionedServer{}
	server0.Name = "test-create-from-sdk"
	server0.Location = "PHX"
	server0.Type = "s1.c1.medium"
	server0.Os = "ubuntu/bionic"
	keys := []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDF9LdAFElNCi7JoWh6KUcchrJ2Gac1aqGRPpdZNowObpRtmiRCecAMb7bUgNAaNfcmwiQi7tos9TlnFgprIcfMWb8MSs3ABYHmBgqEEt3RWYf0fAc9CsIpJdMCUG28TPGTlRXCEUVNKgLMdcseAlJoGp1CgbHWIN65fB3he3kAZcfpPn5mapV0tsl2p+ZyuAGRYdn5dJv2RZDHUZBkOeUobwsij+weHCKAFmKQKtCP7ybgVHaQjAPrj8MGnk1jBbjDt5ws+Be+9JNjQJee9zCKbAOsIo3i+GcUIkrw5jxPU/RTGlWBcemPaKHdciSzGcjWboapzIy49qypQhZe1U75 user@my_ip"}
	server0.SshKeys = keys
	requestCommand0.SetServer(server0)
	resp0,_ := requestCommand0.Execute()
	code := resp0.StatusCode
	if code == 200 {
		response0 := &dto.LongServer{}
		response0.FromBytes(resp0)
		
	} else {
		response0 := &dto.ErrorMessage{}
		response0.FromBytes(resp0)
		fmt.Println("API Returned" + response0.Message)
		

	}

	requestCommand := &command.GetServersCommand{}
	requestCommand.SetRequester(cl)
	resp, err := requestCommand.Execute()
	response := &dto.Servers{}
	if err == nil {
		response.FromBytes(resp)
		fmt.Println("Done!")
	}


	

	
}
