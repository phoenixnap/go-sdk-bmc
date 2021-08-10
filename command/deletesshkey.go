package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// DeleteSshKeyCommand represents command for sshkey deleting
type DeleteSshKeyCommand struct {
	requester client.Requester
	sshKeyID  string
}

// Execute deleting the sshkey
func (command *DeleteSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Delete(apiPrefix + "ssh-keys/" + command.sshKeyID)
}

// SetRequester sets requester to the command
func (command *DeleteSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKeyID sets sshkey id to the command
func (command *DeleteSshKeyCommand) SetSshKeyID(sshKeyID string) {
	command.sshKeyID = sshKeyID
}

//NewDeleteSshKeyCommand constructs new commmand of this type
func NewDeleteSshKeyCommand(requester client.Requester, sshKeyID string) *DeleteSshKeyCommand {
	return &DeleteSshKeyCommand{requester, sshKeyID}
}
