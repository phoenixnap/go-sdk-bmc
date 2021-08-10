package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
)

// ShutDownCommand represents command that reboots the server
type ShutDownCommand struct {
	requester client.Requester
	serverID  string
}

// Execute reboot on the server
func (command *ShutDownCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix+"servers/"+command.serverID+"/actions/shutdown", nil)
}

// SetRequester sets requester to the command
func (command *ShutDownCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *ShutDownCommand) SetServerID(id string) {
	command.serverID = id
}

//NewShutDownCommand constructs new commmand of this type
func NewShutDownCommand(requester client.Requester, serverID string) *ShutDownCommand {

	return &ShutDownCommand{requester, serverID}
}
