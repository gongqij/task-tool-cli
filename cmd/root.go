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
)

var (
	accessKey string
	secretKey string
	endPoint  string
	mgr       *client.Manager
)

const defaultEndPoint = "http://127.0.0.1:8081"
const defaultAccessKey = "UzEatlPmdhdU9b3nSEp61I6Y"
const defaultSecretKey = "Hwibpcoa9yaDBVuGU9kOEJo6"

const defaultTimeout = client.DefaultTimeout

var globalUsage = `The task tool cli

Common actions for task tool cli:

- task-tool-cli add - add tasks
- task-tool-cli delete - delete tasks
- task-tool-cli list - show tasksInfo
- task-tool-cli version - show task-tool-cli version info
`

func NewRootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "task-tool-cli",
		Short:        "The Helm package manager for Kubernetes.",
		Long:         globalUsage,
		SilenceUsage: true,
		// This breaks completion for 'helm help <TAB>'
		// The Cobra release following 1.0 will fix this
		//ValidArgsFunction: noCompletions, // Disable file completion
	}

	//设置全局参数
	cmd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", defaultEndPoint, "HTTP endpoint for task-tool-cli")
	cmd.PersistentFlags().StringVarP(&accessKey, "access_key", "a", defaultAccessKey, "JWT Access Key (optional)")
	cmd.PersistentFlags().StringVarP(&secretKey, "secret_key", "s", defaultSecretKey, "JWT Secret Key (optional)")
	cmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	//Setup shell completion for the access_key flag
	flagName1 := "access_key"
	err := cmd.RegisterFlagCompletionFunc(flagName1, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{defaultAccessKey}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}

	flagName2 := "secret_key"
	err = cmd.RegisterFlagCompletionFunc(flagName2, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{defaultSecretKey}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		logrus.Fatal(err)
	}
	// Add subcommands
	cmd.AddCommand(
		// chart commands
		newListCmd(),
		newDeleteCmd(),
		newAddCmd(),
		newCompletionCmd(),
		newVersionCmd(),
	)
	return cmd, nil
}
