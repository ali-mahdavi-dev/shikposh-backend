package commands

import (
	"os"

	"github.com/spf13/cobra"

	config "shikposh-backend/config"
	"github.com/shikposh/framework/infrastructure/logging"
)

var (
	cfg     config.Config
	rootCmd = &cobra.Command{
		Use: "",
		Run: func(cmd *cobra.Command, args []string) {
			initializeConfigs()
		},
	}
)

func initializeConfigs() {
	cfg = *config.GetConfig()
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(runHTTPServerCMD())
	rootCmd.AddCommand(migrateCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Logger might not be initialized, check if available
		if logging.GetLogger() != nil {
			logging.Error("Command execution failed").WithError(err).Log()
		}
		os.Exit(1)
	}
}
