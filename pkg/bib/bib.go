package bib

import (
	"github.com/spf13/cobra"
)

var BibCmd = &cobra.Command{
	Use:     "bib",
	Short:   "Bib Notes related commands",
	Version: "v0.1.0",
}

func init() {
	BibCmd.AddCommand(addCmd)
}
