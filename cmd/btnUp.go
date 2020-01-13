package cmd

import (
	"github.com/spf13/cobra"
)

var lFlagU bool
var rFlagU bool
// btnUpCmd represents the btnUp command
var btnUpCmd = &cobra.Command{
	Use:   "btn-up",
	Short: "The command performs a button release",
	Long: `The command performs a button release. Example:
./goMouse btn-up -l // release the left button
`,
	RunE: func(cmd *cobra.Command, args []string) error{
		if !lFlagU && !rFlagU {
			myMouse.LeftBtnUp()
			myMouse.RightBtnUp()
			return nil
		}
		if lFlagU {
			myMouse.LeftBtnUp()
		}
		if rFlagU {
			myMouse.RightBtnUp()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(btnUpCmd)
	btnUpCmd.Flags().BoolVarP(&lFlagU, "left", "l", false, "release the left mouse button")
	btnUpCmd.Flags().BoolVarP(&rFlagU, "right", "r", false, "release the right mouse button")
}
