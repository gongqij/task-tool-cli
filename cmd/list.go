/*
Copyright © 2020 ll

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permission s and
limitations under the License.
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
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
			err := mgr.GetTaskInfoById(cmd.Flag("task_id").Value.String())
			if err != nil {
				logrus.Fatal(err)
			}
			return
		}
		if (len(args) > 0 && args[0] != "all") || cmd.Flag("object_type").Value.String() != "" {
			resp, err := mgr.ListAllTasks()
			if err != nil {
				logrus.Fatal(err)
			}
			if len(args) > 0 {
				utils.PrintResult(resp, args[0])
				if !utils.IsObjectTypeExist(args[0]) {
					err := mgr.GetTaskInfoById(args[0])
					if err != nil {
						logrus.Fatal(err)
					}
				}
			} else {
				utils.PrintResult(resp, cmd.Flag("object_type").Value.String())
			}
			return
		}
		resp, err := mgr.ListAllTasks()
		if err != nil {
			logrus.Fatal(err)
		}
		utils.PrintResult(resp, "")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringP("task_id", "t", "", "optional, -task_id xxx -t xxx  get taskInfo by taskID")
	listCmd.Flags().StringP("object_type", "o", "", "optional, -object_type xxx -o xxx get tasks by objectType")
}
