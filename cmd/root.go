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
var defaultRtspSourceForCompletion = []string{"rtsp://10.4.176.152:5454", "rtsp://10.4.176.227:5454", "rtsp://10.4.196.29:5454"}

const defaultEndPoint91 = "https://10.4.192.91:30443/engine/video-process"
const defaultEndPoint92 = "https://10.4.192.92:30443/engine/video-process"
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
		ValidArgsFunction: noCompletions, // Disable file completion
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
	flagName3 := "endpoint"
	err = cmd.RegisterFlagCompletionFunc(flagName3, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{defaultEndPoint92, defaultEndPoint91}, cobra.ShellCompDirectiveNoFileComp
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
