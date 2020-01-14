package cmd

import (
	"github.com/spf13/cobra"
)

var lFlagD bool
var rFlagD bool

// btnDownCmd represents the btnDown command
var btnDownCmd = &cobra.Command{
	Use:   "btn-down",
	Short: "The command performs a button push",
	Long: `The command performs a button push. Example:
./goMouse btn-down -r // push the right button
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !lFlagD && !rFlagD {
			myMouse.LeftBtnDown()
			myMouse.RightBtnDown()
			return nil
		}
		if lFlagD {
			myMouse.LeftBtnDown()
		}
		if rFlagD {
			myMouse.RightBtnDown()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(btnDownCmd)
	btnDownCmd.Flags().BoolVarP(&lFlagD, "left", "l", false, "push the left button")
	btnDownCmd.Flags().BoolVarP(&rFlagD, "right", "r", false, "push the right button")
}
