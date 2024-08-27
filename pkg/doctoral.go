package doctoral

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ubombar/doctoral/internal/doctoral"
	"github.com/ubombar/doctoral/pkg/bib"
)

var rootCmd = &cobra.Command{
	Use:     "doctoral",
	Short:   "Maintain your Obsidian Vault.",
	Long:    `This is a CLI for adding, deleting, updating, and analysing material from the Obsidian vault. Specifically designed for Ufuk's doctoral studies.`,
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		doctoral.DisplayHelp()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bib.BibCmd)
}
