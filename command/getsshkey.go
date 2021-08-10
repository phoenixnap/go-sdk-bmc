package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// GetSshKeyCommand represents command that pulls details about specific sshKey
type GetSshKeyCommand struct {
	requester client.Requester
	sshKeyID  string
}

// Execute pulls details about specific sshKey
func (command *GetSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Get(apiPrefix + "ssh-keys/" + command.sshKeyID)
}

// SetRequester sets requester to the command
func (command *GetSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKeyID sets ssh-key id to the command
func (command *GetSshKeyCommand) SetSshKeyID(id string) {
	command.sshKeyID = id
}

//NewGetSshKeyCommand constructs new commmand of this type
func NewGetSshKeyCommand(requester client.Requester, sshKeyID string) *GetSshKeyCommand {
	return &GetSshKeyCommand{requester, sshKeyID}
}
