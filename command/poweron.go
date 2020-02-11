package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// PowerOnCommand represents command that powers on the server
type PowerOnCommand struct {
	requester client.Requester
	serverID  string
}

// Execute powers on the server
func (command *PowerOnCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Post("servers/"+command.serverID+"/actions/power-on", nil)
}

// SetRequester sets requester to the command
func (command *PowerOnCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *PowerOnCommand) SetServerID(id string) {
	command.serverID = id
}
