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
	"strconv"
	"task-tool-cli/api"
	"task-tool-cli/client"
	"task-tool-cli/utils"
)

func newAddCmd() *cobra.Command {
	// addCmd represents the add command
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "add tasks",
		Run: func(cmd *cobra.Command, args []string) {
			opts := client.Options{
				HTTPAuthFunc: utils.SetupAuth(accessKey, secretKey),
			}
			mgr = client.NewManager(endPoint, defaultTimeout, opts)
			if cmd.Flag("num").Value.String() == "" || cmd.Flag("rtsp").Value.String() == "" || cmd.Flag("object_type").Value.String() == "" {
				logrus.Warnln("请指定任务数、任务类型和rtsp源")
				_ = cmd.Help()
				return
			}
			num, _ := strconv.Atoi(cmd.Flag("num").Value.String())
			err := mgr.AddTasks(num, cmd.Flag("object_type").Value.String(), cmd.Flag("rtsp").Value.String(), cmd.Flag("minio_key").Value.String())
			if err != nil {
				logrus.Fatal(err)
			}
			logrus.Println("添加任务成功！！！")

		},
	}
	addCmd.Flags().StringP("num", "n", "", "required, -num xxx -n xxx task number")
	addCmd.Flags().StringP("object_type", "o", "", "required, -object_type xxx -o xxx task type")
	addCmd.Flags().StringP("rtsp", "r", "", "required, -rtsp xxx -r xxx  task rtsp")
	addCmd.Flags().StringP("minio_key", "m", "", "optional, -minio_key xxx -m xxx minio_key(only for traffic anomaly)")
	addCmd.MarkFlagRequired("num")
	addCmd.MarkFlagRequired("object_type")
	addCmd.MarkFlagRequired("rtsp")

	flagName := "object_type"
	err := addCmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{api.ObjectType_OBJECT_TRAFFIC_ANOMALY_EVENT.String(), api.ObjectType_OBJECT_TRAFFIC_MULTI_PACH.String(), api.ObjectType_OBJECT_TRAFFIC_AUTOMOBILE_COUNT.String(), api.ObjectType_OBJECT_TRAFFIC_CAMERA_VISION_INFO.String()}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return addCmd
}
