package cmd

import (
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "This command reset mouse setting to default",
	Long: `This command reset mouse setting to default
`,
	Run: func(cmd *cobra.Command, args []string) {
		myMouse.Reset(myScreen.width, myScreen.height)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
