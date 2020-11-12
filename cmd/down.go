package cmd

import (
	"fmt"
	"os"

	"github.com/axllent/wiregod/app"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down [interface]",
	Short: "Take a Wireguard interface offline",
	Long:  `Take a Wireguard interface offline`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connections := app.WGConnections()
		if len(connections) == 0 {
			fmt.Println("Wireguard is not connected")
			os.Exit(1)
		}
		var wg string

		if len(connections) == 1 {
			wg = connections[0]
		}

		if len(args) == 1 {
			wg = args[0]
		}

		if wg == "" {
			wg = app.Select("Select available connection", connections)
		}

		if wg == "" {
			return
		}

		if !app.InArray(wg, connections) {
			fmt.Println(wg, "is not connected")
			os.Exit(1)
		}

		output, err := app.SudoExec("wg-quick down " + wg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(output)
		fmt.Printf("Wireguard on %s is down\n", wg)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
