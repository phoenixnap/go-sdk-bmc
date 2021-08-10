package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// GetSshKeysCommand represents command for sshkeys retrieval
type GetSshKeysCommand struct {
	requester client.Requester
}

// Execute retrieves all ssh-keys
func (command *GetSshKeysCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Get(apiPrefix + "ssh-keys")
}

// SetRequester sets requester to the command
func (command *GetSshKeysCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

//NewGetSshKeysCommand constructs new commmand of this type
func NewGetSshKeysCommand(requester client.Requester) *GetSshKeysCommand {
	return &GetSshKeysCommand{requester}
}
