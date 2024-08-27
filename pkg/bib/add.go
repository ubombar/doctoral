package bib

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/ubombar/doctoral/internal/doctoral"
)

var (
	templateFile string
	searchDirs   []string
	copy         bool
	bibtex       bool
	interactive  bool
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add Bib Notes to your Obsidian Vault.",
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Waiting for at least one material name/url")
			os.Exit(-1)
		}
		for _, identifier := range args {
			fmt.Printf("Trying for %q\n", identifier)

			itype := doctoral.GetTypeOfIdentifier(identifier)

			if itype == doctoral.UNKNOWN {
				fmt.Println("Unknown identifier specified, skipping.")
				continue
			} else if itype == doctoral.PDF {
				// Do PDF matching
				re, err := regexp.Compile(fmt.Sprintf("(.*%s.*)|(.*%s\\.(?i)(pdf))", identifier, identifier))

				if err != nil {
					fmt.Printf("Cannot create a searching regex for %s. Skipping\n", identifier)
					continue
				}

				testDir := "/home/ubombar/Downloads/Strategies for Sound Internet Measurement.PDF"

				fmt.Printf("re.Match([]byte(testDir)): %v\n", re.Match([]byte(testDir)))
				// Then get the finding directories, search for matched cases.
			}

		}

	},
}

func init() {
	// Set the flags
	addCmd.Flags().StringVar(&templateFile, "template", "", "Template file to be added as a Bib Note, add an empty one with just the same name of the material.")
	addCmd.Flags().StringArrayVar(&searchDirs, "search-dirs", []string{}, "Directories to look for material in the local machine.")
	addCmd.Flags().BoolVar(&copy, "copy", false, "Copy the material in the search directory instead of move.")
	addCmd.Flags().BoolVar(&bibtex, "bibtex", false, "Try to find or generate bibtext of the resource.")
	addCmd.Flags().BoolVar(&interactive, "interactive", true, "Promt user in cases like multiple files.")
}
