package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "ModelBox client commands provides sub-commands to access all the objects exposed by the api",
	Long:  `Follow the modelbox client subcommands to access various modelbox features exposed by the api`,
}

var clientInitConfigCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Creates a sample client config",
	Long: `Creates a sample client config. By default the config
is created in the current directory. Help:
./modelbox client --init-config path/to/new/config`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		err := WriteClientConfigToFile(path)
		if err == nil {
			fmt.Printf("config written to path: %v\n", path)
		} else {
			fmt.Printf("unable to write config: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(clientInitConfigCmd)
	clientInitConfigCmd.Flags().String("path", "./modelbox_client.toml", "path to write the client config")

	// Initialize logger for the client
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
