package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "This command shows the settings values of mouse",
	Long: `This command shows the settings values of mouse
`,
	Run: func(cmd *cobra.Command, args []string) {
		myMouse.Info()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
