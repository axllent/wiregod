package cmd

import (
	"fmt"

	"github.com/axllent/wiregod/app"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show your public IP",
	Long:  `Show your public IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := app.PublicIP()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Printf("Public IP: %s\n", ip)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
