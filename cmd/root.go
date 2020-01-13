package cmd

import (
	"fmt"
	"github.com/BogdanYanov/goMouse/mouse"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type screen struct {
	width, height uint32
}

var cfgFile string
var myMouse *mouse.Mouse
var myScreen screen
var width, height uint32 = 1024, 768


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goMouse",
	Short: "goMouse is CLI tool which simulates mouse control",
	Long: `goMouse is CLI tool which simulates mouse control by commands.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	myScreen = screen{width, height}
	myMouse = mouse.NewMouse(myScreen.width, myScreen.height)
	if mouse.FileExists() {
		mouse.GetMouse(myMouse)
	} else {
		myMouse.WriteJSON()
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goMouse.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".goMouse" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goMouse")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
