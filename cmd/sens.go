package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// sensCmd represents the sens command
var sensCmd = &cobra.Command{
	Use:   "sens",
	Short: "Set mouse sensitivity",
	Long: `Set mouse sensitivity. Example:
./goMouse sens 5 // set mouse sensitivity 5
`,
	RunE: func(cmd *cobra.Command, args []string) error{
		if len(args) == 0 {
			return fmt.Errorf("need argument of sensitivity")
		} else if len(args) > 1 {
			return fmt.Errorf("too much arguments for sensitivity")
		}
		mSens, err := strconv.ParseUint(args[0], 10, 8)
		if err != nil {
			return fmt.Errorf("invalid argument for sensitivity: %s", err)
		}
		myMouse.Sensitivity(uint8(mSens))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sensCmd)
}
