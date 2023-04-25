package command

import (
	"cli/src/client/shell"
	"strings"
)

type SSHCommand struct {
	RemoteHost     string
	RemotePort     string
	RemoteUser     string
	RemotePassword string
}

func (S *SSHCommand) Exec(args []string) {
	cmd_element := strings.Split(args[0], ";")
	for _, shell_cmd := range cmd_element {
		shellApi := shell.SSHExec{}
		shellApi.SSH_Exec(S.RemoteHost, S.RemotePort, S.RemoteUser, S.RemotePassword, shell_cmd)
	}

}
