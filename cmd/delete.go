/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := client.Options{
			HTTPAuthFunc: utils.SetupAuth(accessKey, secretKey),
		}
		mgr = client.NewManager(endPoint, defaultTimeout, opts)
		//fmt.Println(cmd.Flag("task_id").Value)
		if cmd.Flag("task_id").Value.String() != "" {
			err := mgr.DeleteTaskById(cmd.Flag("task_id").Value.String())
			if err != nil {
				logrus.Fatal(err)
			}
			return
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
			return
		}
		if len(args) > 0 && args[0] == "all" {
			err := mgr.DeleteAllTasks()
			if err != nil {
				logrus.Fatal(err)
			}
			return
		}
		logrus.Warnln("请指定参数！！！")
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("task_id", "t", "", "optional, -task_id xxx -t xxx  delete task by taskID")
	deleteCmd.Flags().StringP("object_type", "o", "", "optional, -object_type xxx -o xxx delete tasks by objectType")
}