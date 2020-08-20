package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

func newListCmd() *cobra.Command {
	// listCmd represents the list command
	listCmd := &cobra.Command{
		Use:               "list OBJECT_TYPE[all(获取全部任务),taskID(获取任务详细信息)",
		Short:             "show tasksInfo",
		Aliases:           []string{"show", "get", "ls", "info"},
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := client.Options{
				HTTPAuthFunc: utils.SetupAuth(accessKey, secretKey),
			}
			mgr = client.NewManager(endPoint, defaultTimeout, opts)
			//fmt.Println(cmd.Flag("task_id").Value)
			if cmd.Flag("task_id").Value.String() != "" {
				err := mgr.GetTaskInfoById(cmd.Flag("task_id").Value.String())
				if err != nil {
					return err
				}
			}
			if (len(args) > 0 && args[0] != "all") || cmd.Flag("object_type").Value.String() != "" {
				resp, err := mgr.ListAllTasks()
				if err != nil {
					return err
				}
				if len(args) > 0 {
					utils.PrintResult(resp, args[0])
					if !utils.IsObjectTypeExist(args[0]) {
						err := mgr.GetTaskInfoById(args[0])
						if err != nil {
							return err
						}
					}
				} else {
					utils.PrintResult(resp, cmd.Flag("object_type").Value.String())
				}
				return nil
			}
			resp, err := mgr.ListAllTasks()
			if err != nil {
				return err
			}
			utils.PrintResult(resp, "")
			return nil
		},
		ValidArgs: append(utils.AllObjectType(), "all"),
	}

	listCmd.Flags().StringP("task_id", "t", "", "optional, --task_id xxx -t xxx  get taskInfo by taskID")
	listCmd.Flags().StringP("object_type", "o", "", "optional, --object_type xxx -o xxx get tasks by objectType")

	err := listCmd.RegisterFlagCompletionFunc("object_type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return utils.AllObjectType(), cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return listCmd
}
