package command

import (
	"cli/src/common"
	"fmt"

	"github.com/spf13/cobra"
)

func InitParser() *cobra.Command {
	var rootCmd = &cobra.Command{Use: common.APP_NAME}
	rootCmd.AddCommand(initPodCli())
	rootCmd.AddCommand(initSSHCli())
	rootCmd.AddCommand(Version())

	return rootCmd
}

func initPodCli() *cobra.Command {
	cliPodCmdObj := PodCommand{}
	var cliPodCmd = &cobra.Command{
		Use:   common.SUB_CMD_CLI,
		Short: "Exec CLI Command With Remote Pod",
		Long:  "Exec CLI Command With Remote Pod",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cliPodCmdObj.Exec(args)
		},
	}
	cliPodCmd.Flags().StringVarP(&cliPodCmdObj.QueryCondition, common.SUB_CMD_QUERY, common.SUB_CMD_QUERY_SHORT, "", "Container名称或PodIP")
	cliPodCmd.Flags().StringVarP(&cliPodCmdObj.ContainerName, common.SUB_CMD_CONTAINER, common.SUB_CMD_CONTAINER_SHORT, "Defaulted container", "Container名称[当提供PodIP时使用此参数]")
	cliPodCmd.Flags().StringVarP(&cliPodCmdObj.NameSpace, common.SUB_CMD_NAMESPACE, common.SUB_CMD_NAMESPACE_SHORT, "default", "NameSpace名称")
	cliPodCmd.Flags().StringVarP(&cliPodCmdObj.Mode, common.SUB_CMD_CLI_MODE, common.SUB_CMD_CLI_MODE_SHORT, "host", "当前执行环境类型[host or pod]")
	cliPodCmd.Flags().IntVarP(&cliPodCmdObj.Index, common.SUB_CMD_INDEX, common.SUB_CMD_INDEX_SHORT, 1, "在第几个Pod中执行[当匹配到多个Pod时]")
	cliPodCmd.PersistentFlags().String("", "", "指明需要执行的CLI Command")
	cliPodCmd.MarkFlagRequired(common.SUB_CMD_QUERY)
	cliPodCmd.MarkFlagsMutuallyExclusive(common.SUB_CMD_CONTAINER, common.SUB_CMD_INDEX)

	return cliPodCmd
}

func initSSHCli() *cobra.Command {
	cliSSHCmdObj := SSHCommand{}
	var cliSSHCmd = &cobra.Command{
		Use:   common.SUB_CMD_SSH,
		Short: "Exec CLI Command With Remote Host",
		Long:  "Exec CLI Command With Remote Host",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cliSSHCmdObj.Exec(args)
		},
	}
	cliSSHCmd.Flags().StringVarP(&cliSSHCmdObj.RemoteHost, common.SUB_CMD_SSH_HOST, common.SUB_CMD_SSH_HOST_SHORT, "127.0.0.1", "目标主机IP")
	cliSSHCmd.Flags().StringVarP(&cliSSHCmdObj.RemotePort, common.SUB_CMD_SSH_PORT, common.SUB_CMD_SSH_PORT_SHORT, "22", "目标主机端口")
	cliSSHCmd.Flags().StringVarP(&cliSSHCmdObj.RemoteUser, common.SUB_CMD_SSH_USER, common.SUB_CMD_SSH_USER_SHORT, "root", "目标主机用户")
	cliSSHCmd.Flags().StringVarP(&cliSSHCmdObj.RemotePassword, common.SUB_CMD_SSH_PASS, common.SUB_CMD_SSH_PASS_SHORT, "", "目标主机用户密码")
	cliSSHCmd.PersistentFlags().String("", "", "指明需要执行的CLI Command,多个命令以分号分割并引用双引号")
	cliSSHCmd.MarkFlagRequired(common.SUB_CMD_SSH_PASS)
	return cliSSHCmd
}

func Version() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get the latest version",
		Long:  "get the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Client Version: %s", common.APP_VERSON)
		},
	}
	versionCmd.PersistentFlags().BoolP("", "", false, "get the version")
	return versionCmd
}
