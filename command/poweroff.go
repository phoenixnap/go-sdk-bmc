package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
)

// PowerOffCommand represents command that powers off the server
type PowerOffCommand struct {
	requester client.Requester
	serverID  string
}

// Execute powers off the server
func (command *PowerOffCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix + "servers/"+command.serverID+"/actions/power-off", nil)
}

// SetRequester sets requester to the command
func (command *PowerOffCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *PowerOffCommand) SetServerID(id string) {
	command.serverID = id
}

//NewPowerOffCommand constructs new commmand of this type
func NewPowerOffCommand(requester client.Requester, serverID string) *PowerOffCommand {

	return &PowerOffCommand{requester, serverID}
}
