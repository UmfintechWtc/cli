package shell

import (
	"cli/src/common"
	"fmt"
	"os"
	"time"

	xlogger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type SSHExec struct {
	SSHConfig
}
type SSHConfig struct {
}

func (C SSHConfig) SSHClientConfig(user, password string) *ssh.ClientConfig {
	ssh_config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout:         1800 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh_config
}

func (S SSHExec) SSHConnToRemoteHost(host, port string, config *ssh.ClientConfig) (*ssh.Session, error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	ssh_client, ssh_client_err := ssh.Dial("tcp", addr, config)
	if ssh_client_err != nil {
		fmt_message := fmt.Sprintf("%s:%s\n", host, port)
		xlogger.Error(common.Rsperror(
			common.XE_SSH_CONN_ERROR,
			fmt_message+ssh_client_err.Error()))
		os.Exit(5)
	}
	ssh_session, ssh_session_err := ssh_client.NewSession()
	if ssh_session_err != nil {
		fmt_message := fmt.Sprintf("%s:%s\n", host, port)
		xlogger.Error(common.Rsperror(
			common.XE_SSH_SESSION_ERROR,
			fmt_message+ssh_client_err.Error()))
		os.Exit(6)
	}
	return ssh_session, nil
}

func (S SSHExec) SSH_Exec(host, port, user, password, shellCmd string) {
	session, _ := S.SSHConnToRemoteHost(host, port, S.SSHConfig.SSHClientConfig(user, password))
	cmd_res, cmd_res_err := session.Output(shellCmd)
	if cmd_res_err != nil {
		fmt_message := fmt.Sprintf("%s:%s[%s]", host, port, shellCmd)
		xlogger.Error(common.Rsperror(
			common.XE_EXEC_ERROR,
			fmt_message+cmd_res_err.Error()))
		os.Exit(4)
	}
	defer session.Close()
	xlogger.Info("命令输出: ", string(cmd_res))
}
