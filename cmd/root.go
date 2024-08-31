/*
Copyright © 2024 Ufuk BOMBAR <ufukbombar@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
			// Get the documents under search dirs
			documents, err := doctoral.GetDocumentsUnderDirectories(config.SearchDirectories)

			if err != nil {
				fmt.Printf("Cannot list the files under search directories: %s\n", err)
				return
			}

			// Create the bib note file as well. It is a .md file
			template, err := doctoral.NewDocument(config.TemplateFile)
			if err != nil {
				fmt.Printf("\tCannot find template file %q: %s\n", config.TemplateFile, err)
				return
			}

			// Create a simple menu and get the choices from user
			menu := doctoral.NewSimpleMenu()
			choices := menu.GetChoices(documents)

			for i, choice := range choices {
				// Create a document pointing to the new location.
				choiceCopy, err := doctoral.NewDocument(filepath.Join(config.PDFDirectory, choice.FileName))
				if err != nil {
					fmt.Printf("\tCannot copy material file %q: %s\n", choice.FileName, err)
					continue
				}

				// Create the bib note file as well. It is a .md file
				bibNote, err := doctoral.NewDocument(filepath.Join(config.BibNotesDirectory, fmt.Sprint(choice.FileNameWithoutExt(), ".md")))
				if err != nil {
					fmt.Printf("\tCannot copy file %q: %s\n", bibNote.FileName, err)
					continue
				}

				// Skip since it exists
				if !config.OverwritePDFFiles && choiceCopy.ExistOnDisk() {
					fmt.Println("\tCannot copy/move material file because it already exists and 'overwritePDFFiles' is set to false")
					continue
				}

				// Skip since it exists
				if !config.OverwriteBibNoteFiles && bibNote.ExistOnDisk() {
					fmt.Println("\tCannot create bib note because it already exists and 'overwriteBibNoteFiles' is set to false")
					continue
				}

				// Then create the bib from template file
				if err := bibNote.TemplateContent(*template, doctoral.NewTemplateData(config, bibNote, choiceCopy)); err != nil {
					fmt.Printf("\tCannot create bib file %q: %s\n", bibNote.FileName, err)
					continue
				}

				if err := choice.CopyToFile(choiceCopy.AbsolutePath); err != nil {
					fmt.Printf("\tCannot copy material file %q: %s\n", choiceCopy.FileName, err)
					continue
				}

				if config.DeleteAfterCopyingPDFs {
					// Delete the old file so it will be moved instead of copied.
					if err := choice.Delete(); err != nil {
						fmt.Printf("\tCannot remove old file %q: %s\n", choice.FileName, err)
						continue
					}
					fmt.Printf("\t(%d/%d): Moved material file and created template for %q\n", i, len(choices), choiceCopy.FileNameWithoutExt())
				} else {
					fmt.Printf("\t(%d/%d): Copied material file and created template for %q\n", i, len(choices), choiceCopy.FileNameWithoutExt())
				}
			}
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
