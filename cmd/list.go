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
	"task-tool-cli/api"
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
		ValidArgs: []string{"all", api.ObjectType_OBJECT_TRAFFIC_ANOMALY_EVENT.String(), api.ObjectType_OBJECT_TRAFFIC_MULTI_PACH.String(), api.ObjectType_OBJECT_TRAFFIC_AUTOMOBILE_COUNT.String(), api.ObjectType_OBJECT_TRAFFIC_CAMERA_VISION_INFO.String()},
	}

	listCmd.Flags().StringP("task_id", "t", "", "optional, --task_id xxx -t xxx  get taskInfo by taskID")
	listCmd.Flags().StringP("object_type", "o", "", "optional, --object_type xxx -o xxx get tasks by objectType")
	flagName := "object_type"
	err := listCmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{api.ObjectType_OBJECT_TRAFFIC_ANOMALY_EVENT.String(), api.ObjectType_OBJECT_TRAFFIC_MULTI_PACH.String(), api.ObjectType_OBJECT_TRAFFIC_AUTOMOBILE_COUNT.String(), api.ObjectType_OBJECT_TRAFFIC_CAMERA_VISION_INFO.String()}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return listCmd
}
