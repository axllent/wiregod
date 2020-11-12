package cmd

import (
	"fmt"
	"os"

	"github.com/axllent/wiregod/app"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up [interface]",
	Short: "Bring a Wireguard interface online",
	Long:  `Brings a Wireguard interface online`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		available := app.Configs()
		if len(available) == 0 {
			fmt.Println("You do not have any Wireguard configs in /etc/wireguard")
			os.Exit(1)
		}
		var wg string

		if len(available) == 1 {
			wg = available[0]
		}

		if len(args) == 1 {
			wg = args[0]
		}

		if wg == "" {
			wg = app.Select("Select available config", available)
		}

		if wg == "" {
			return
		}

		if !app.InArray(wg, available) {
			fmt.Println(wg, "is not a valid interface")
			os.Exit(1)
		}

		output, err := app.SudoExec("wg-quick up " + wg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
		fmt.Printf("Wireguard on %s is up\n", wg)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
