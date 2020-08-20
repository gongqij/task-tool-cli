package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

func newAddCmd() *cobra.Command {
	// addCmd represents the add command
	var addCmd = &cobra.Command{
		Use:               "add",
		Short:             "add tasks",
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := client.Options{
				HTTPAuthFunc: utils.SetupAuth(accessKey, secretKey),
			}
			mgr = client.NewManager(endPoint, defaultTimeout, opts)
			if cmd.Flag("num").Value.String() == "" || cmd.Flag("rtsp").Value.String() == "" || cmd.Flag("object_type").Value.String() == "" {
				logrus.Warnln("请指定任务数、任务类型和rtsp源")
				_ = cmd.Help()
				return nil
			}
			num, _ := strconv.Atoi(cmd.Flag("num").Value.String())
			err := mgr.AddTasks(num, cmd.Flag("object_type").Value.String(), cmd.Flag("rtsp").Value.String(), cmd.Flag("minio_key").Value.String())
			if err != nil {
				return err
			}
			logrus.Println("添加任务成功！！！")
			return nil
		},
	}
	addCmd.Flags().StringP("num", "n", "", "required, -num xxx -n xxx task number")
	addCmd.Flags().StringP("object_type", "o", "", "required, -object_type xxx -o xxx task type")
	addCmd.Flags().StringP("rtsp", "r", "", "required, -rtsp xxx -r xxx  task rtsp")
	addCmd.Flags().StringP("minio_key", "m", "", "optional, -minio_key xxx -m xxx minio_key(only for traffic anomaly)")
	addCmd.MarkFlagRequired("num")
	addCmd.MarkFlagRequired("object_type")
	addCmd.MarkFlagRequired("rtsp")

	err := addCmd.RegisterFlagCompletionFunc("object_type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return utils.AllObjectType(), cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}
	err = addCmd.RegisterFlagCompletionFunc("rtsp", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultRtspSourceForCompletion, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return addCmd
}
