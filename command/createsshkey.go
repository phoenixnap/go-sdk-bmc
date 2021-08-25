package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
	"github.com/phoenixnap/go-sdk-bmc/dto"
)

// CreateSshKeyCommand represents command for sshkey creating. Use NewCreateSshKeyCommand to initilize command properly.
type CreateSshKeyCommand struct {
	requester client.Requester
	sshKey    dto.SshKey
}

// Execute create new ssh key
func (command *CreateSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix + "ssh-keys", command.sshKey.ToBytes())
}

// SetRequester - sets requester to the command
func (command *CreateSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKey - sets sshKey details to the command
func (command *CreateSshKeyCommand) SetSshKey(sshKey dto.SshKey) {
	command.sshKey = sshKey
}

//NewCreateSshKeyCommand - constructs new command of this type
func NewCreateSshKeyCommand(requester client.Requester, sshKey dto.SshKey) *CreateSshKeyCommand {
	return &CreateSshKeyCommand{requester, sshKey}
}