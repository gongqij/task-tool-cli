package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

func newDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:               "delete OBJECT_TYPE[all(全部删除),taskID(可以指定多个taskID)]",
		Short:             "delete tasks",
		Aliases:           []string{"del", "remove", "rm"},
		SuggestFor:        []string{"un"},
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := client.Options{
				HTTPAuthFunc: utils.SetupAuth(accessKey, secretKey),
			}
			mgr = client.NewManager(endPoint, defaultTimeout, opts)
			//fmt.Println(cmd.Flag("task_id").Value)
			if cmd.Flag("task_id").Value.String() != "" {
				err := mgr.DeleteTaskById(cmd.Flag("task_id").Value.String())
				if err != nil {
					return err
				}
				return nil
			}
			//在delete命令后直接输入objectType或-object_type xxx,都可以直接删除该类型所有任务
			if (len(args) > 0 && args[0] != "all") || cmd.Flag("object_type").Value.String() != "" {
				if len(args) > 0 {
					for _, arg := range args {
						if utils.IsObjectTypeExist(arg) {
							err := mgr.DeleteTaskByObjectType(arg)
							if err != nil {
								logrus.Warnln(err)
							}
						} else {
							err := mgr.DeleteTaskById(arg)
							if err != nil {
								logrus.Warnln(err)
							}
						}
					}
				} else {
					err := mgr.DeleteTaskByObjectType(cmd.Flag("object_type").Value.String())
					if err != nil {
						logrus.Warnln(err)
					}
				}
				return nil
			}
			if len(args) > 0 && args[0] == "all" {
				err := mgr.DeleteAllTasks()
				if err != nil {
					return err
				}
				return nil
			}
			logrus.Warnln("请指定参数！！！")
			_ = cmd.Help()
			return nil
		},
		ValidArgs: append(utils.AllObjectType(), "all"),
	}
	deleteCmd.Flags().StringP("task_id", "t", "", "optional, --task_id xxx -t xxx  delete task by taskID")
	deleteCmd.Flags().StringP("object_type", "o", "", "optional, --object_type xxx -o xxx delete tasks by objectType")
	flagName := "object_type"
	err := deleteCmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return utils.AllObjectType(), cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return deleteCmd
}
