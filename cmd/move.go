package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Moving mouse cursor to X Y coordinates relative to the screen",
	Long: `Moving mouse cursor to X Y coordinates relative to the screen. Example:
./goMouse move 312 312 // move mouse cursor to coords x - 312, y - 312

! X Y must be greater than or equal to zero !
`,
	RunE: func(cmd *cobra.Command, args []string) error{
		if len(args) != 2 {
			return fmt.Errorf("must be two arguments: X Y")
		}
		x, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid argument X: %s\n", err)
		}
		y, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid argument Y: %s\n", err)
		}
		myMouse.Move(uint32(x), uint32(y), myScreen.width, myScreen.height)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
