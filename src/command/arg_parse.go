package command

import (
	"cli/src/common"

	"github.com/spf13/cobra"
)

func InitParser() *cobra.Command {
	var rootCmd = &cobra.Command{Use: common.APP_NAME}
	rootCmd.AddCommand(initCli())
	return rootCmd
}

func initCli() *cobra.Command {
	cliCmdObj := CliCommand{}
	var cliCmd = &cobra.Command{
		Use:   common.SUB_CMD_CLI,
		Short: "Exec CLI Command",
		Long:  "Exec CLI Command",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cliCmdObj.Exec(args)
		},
	}
	cliCmd.Flags().StringVarP(&cliCmdObj.PodName, common.SUB_CMD_PODNAME, common.SUB_CMD_PODNAME_SHORT, "", "Pod名称")
	cliCmd.Flags().StringVarP(&cliCmdObj.Container, common.SUB_CMD_CONTAINER, common.SUB_CMD_CONTAINER_SHORT, "", "Container名称")
	cliCmd.Flags().StringVarP(&cliCmdObj.NameSpace, common.SUB_CMD_NAMESPACE, common.SUB_CMD_NAMESPACE_SHORT, "default", "NameSpace名称")
	cliCmd.PersistentFlags().String("", "", "指明需要执行的CLI Command")
	cliCmd.MarkFlagRequired(common.SUB_CMD_PODNAME)
	cliCmd.MarkFlagRequired(common.SUB_CMD_CONTAINER)
	return cliCmd
}
