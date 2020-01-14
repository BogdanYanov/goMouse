package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scrlUp bool
var scrlDown bool

// scrollCmd represents the scroll command
var scrollCmd = &cobra.Command{
	Use:   "scroll",
	Short: "This command scroll the mouse wheel up or down",
	Long: `This command scroll the mouse wheel up or down. Example:
./goMouse scroll -u // scroll up the mouse wheel`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !scrlUp && !scrlDown {
			return fmt.Errorf("need flag --up [-u] or --down [-d]")
		}
		if scrlUp && scrlDown {
			return fmt.Errorf("need only one flag: --up [-u] or --down [-d]")
		}
		if scrlUp {
			myMouse.ScrollUp()
		} else {
			myMouse.ScrollDown()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scrollCmd)
	scrollCmd.Flags().BoolVarP(&scrlUp, "up", "u", false, "scroll up")
	scrollCmd.Flags().BoolVarP(&scrlDown, "down", "d", false, "scroll down")
}
