package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/dto"
)

// UpdateSshKeyCommand represents command for sshkey updating. Use NewUpdateSshKeyCommand to initilize command properly.
type UpdateSshKeyCommand struct {
	requester client.Requester
	sshKey    dto.SshKey
}

// Execute update the ssh key
func (command *UpdateSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Put("ssh-keys/"+command.sshKey.ID, command.sshKey.ToBytes())
}

// SetRequester - sets requester to the command
func (command *UpdateSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKey - sets sshKey details to the command
func (command *UpdateSshKeyCommand) SetSshKey(sshKey dto.SshKey) {
	command.sshKey = sshKey
}

//NewUpdateSshKeyCommand - constructs new command of this type
func NewUpdateSshKeyCommand(requester client.Requester, sshKey dto.SshKey) *UpdateSshKeyCommand {
	return &UpdateSshKeyCommand{requester, sshKey}
}
