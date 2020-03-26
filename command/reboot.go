package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// RebootCommand represents command that reboots the server
type RebootCommand struct {
	requester client.Requester
	serverID  string
}

// Execute reboot on the server
func (command *RebootCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Post("servers/"+command.serverID+"/actions/reboot", nil)
}

// SetRequester sets requester to the command
func (command *RebootCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *RebootCommand) SetServerID(id string) {
	command.serverID = id
}

//NewRebootCommand constructs new commmand of this type
func NewRebootCommand(requester client.Requester, serverID string) *RebootCommand {

	return &RebootCommand{requester, serverID}
}
