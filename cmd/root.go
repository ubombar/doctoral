/*
Copyright © 2024 Ufuk BOMBAR <ufukbombar@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ubombar/doctoral/pkg/doctoral"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doctoral",
	Short: "Doctoral Articles & Materials Manager",
	Long:  `This is a CLI application for managing the Notes in an Obsidian Vault.`,
	Run: func(cmd *cobra.Command, args []string) {
		// The config file, all of the information is stored here.
		var err error
		var config *doctoral.Config

		// Get the config file from the env variable, if it doesn't exist get the default one
		configPath := doctoral.GetConfigPathOrDefault()

		// Try to read the config file
		config, err = doctoral.ReadFromConfig(configPath)

		if err != nil {
			fmt.Printf("Cannot found the config file on %q creating under %q with default values\n", configPath, doctoral.GetDefaultConfigPath())

			// Try to create a new config file
			config, err = doctoral.CreateNewConfig(configPath)

			if err != nil {
				fmt.Println("Cannot create the config file, exitting")
				return
			}
		}

		// If it is run with no arguments, run in interactive mode.
		if len(args) == 0 {
			fmt.Printf("config.TemplateFile: %v\n", config.TemplateFile)
		} else { // Run in non-interactive modem just get the pdf/blog/article etc.
			fmt.Println("Running in non-interactive mode is not supported yet.")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doctoral.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
