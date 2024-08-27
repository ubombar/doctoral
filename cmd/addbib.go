package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addBibCmd = &cobra.Command{
	Use:     "addbib",
	Short:   "Add Bib Notes to your Obsidian Vault.",
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not working yet.")
	},
}

func init() {
	rootCmd.AddCommand(addBibCmd)
}
