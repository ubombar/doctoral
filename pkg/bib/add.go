package bib

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ubombar/doctoral/internal/doctoral"
)

var (
	templateFile   string
	searchDirs     []string
	copy           bool
	bibtex         bool
	interactive    bool
	bibDir         string
	pdfDir         string
	pdfOnly        bool
	noTemplate     bool
	forceOverwrite bool
	tags           []string
	status         string
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

			switch itype {
			case doctoral.FILE:
				fmt.Println("\tDetected a file, copying contents and adding it to the Obsidian Box")
				candidates := doctoral.FindRequestedFile(identifier, searchDirs)
				var candidate string

				if len(candidates) == 0 {
					fmt.Printf("\tERROR: Cannot find any candidates for the given filename %q\n", identifier)
					continue
				} else if len(candidates) > 1 {
					// if interacfive ask user which one to pick, you can use hashes of the pandoc of the pdfs
					// so you can reduce the number of selections.
					if interactive {
						// TODO stuff
					} else {
						// pick the first one
						candidate = candidates[0]
					}
				} else {
					// Get the element
					candidate = candidates[0]
				}

				// Check if it is a pdf file
				if pdfOnly && !doctoral.IsAPDFFile(candidate) {
					fmt.Println("\tERROR: Given is not a pdf file, to allow all files use --pdf-only false flag.")
					continue
				}

				// Calcualte the destination path in the pdf dir
				destinationpath := doctoral.CalculateDestinationPath(candidate, pdfDir)

				// Transfer pdf file to the new place
				if err := doctoral.TransferFileContent(candidate, destinationpath, !copy, forceOverwrite); err != nil {
					fmt.Printf("Cannot move/copy the file %q\n", err)
					continue
				}

				// Get the tags + default tags
				allTags := append(doctoral.GetDefaultTags(), tags...)

				// If we want to create the template after transfering the file.
				if !noTemplate {
					// Create the bib note from the given tempalte
					if err := doctoral.CreateBibTemplate(templateFile, bibDir, candidate, identifier, forceOverwrite, allTags, status); err != nil {
						fmt.Printf("\tERROR: Cannot create the template file %q\n", err)
						continue
					}
				}

				// Done
				fmt.Println("\tAdded file to the Obsidian Vault")

			default:
				fmt.Println("\tI have no idea what media type this is, skipping.")
			}

		}

	},
}

func init() {
	// Set the flags
	addCmd.Flags().StringVar(&templateFile, "template", doctoral.GetDefaultTemplateFile(), "Template file to be added as a Bib Note, add an empty one with just the same name of the material.")
	addCmd.Flags().StringArrayVar(&searchDirs, "search-dirs", doctoral.GetDefaultSearchDirs(), "Directories to look for material in the local machine.")
	addCmd.Flags().BoolVar(&copy, "copy", false, "Copy the material in the search directory instead of move.")
	addCmd.Flags().BoolVar(&bibtex, "bibtex", false, "Try to find or generate bibtext of the resource.")
	addCmd.Flags().BoolVar(&interactive, "interactive", true, "Promt user in cases like multiple files.")
	addCmd.Flags().StringVar(&bibDir, "bib-dir", doctoral.GetDefaultBibDir(), "The location of the bib notes.")
	addCmd.Flags().StringVar(&pdfDir, "pdf-dir", doctoral.GetDefaultPDFDir(), "The location of the pdfs.")
	addCmd.Flags().BoolVar(&pdfOnly, "pdf-only", true, "Only transfer pdf files.")
	addCmd.Flags().BoolVar(&noTemplate, "no-template", false, "Only transfer the file, do not create a template.")
	addCmd.Flags().BoolVar(&forceOverwrite, "force-overwrite", false, "Overwrite the bib not even if already exist.")
	addCmd.Flags().StringArrayVar(&tags, "tags", []string{"#type/"}, "Tags that will be added.")
	addCmd.Flags().StringVar(&status, "status", "#status/waiting", "Status of the bib file.")
}
