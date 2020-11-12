package cmd

import (
	"github.com/axllent/wiregod/app"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View current wireguard connection status",
	Long:  `View current wireguard connection status`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.WGStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
