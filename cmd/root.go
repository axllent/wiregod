package cmd

import (
	"fmt"
	"os"

	"github.com/axllent/wiregod/app"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wiregod",
	Short: "Wiregod",
	Long: `Wiregod is a helper utility for handling one or more Wireguard client configurations.
Issues & documentation: https://github.com/axllent/wiregod`,
	SilenceUsage:  true, // suppress help screen on error
	SilenceErrors: true, // suppress duplicate error on error
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.WGStatus()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if err := rootCmd.Execute(); err != nil {
			fmt.Printf("Error: %v\n", err)

			// detect if subcommand is valid
			help := "\nSee: `wiregod -h` for help"
			if len(os.Args) > 1 {
				for _, t := range rootCmd.Commands() {
					if t.Name() == os.Args[1] {
						help = "\nSee: `wiregod " + os.Args[1] + " -h` for help"
					}
				}
			}

			fmt.Println(help)

			os.Exit(1)
		}
	}
}

func init() {
	cobra.OnInitialize(app.CheckSystemApps)

	// hide the `help` command
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")

	rootCmd.PersistentFlags().Lookup("help").Hidden = true
}
