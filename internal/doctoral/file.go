package doctoral

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type BibNote struct {
	Date       string
	Identifier string
	Tags       []string
	Status     string
}

// Searches the given file on the search dirs, returns candidates.
// In the feature I can add an option for recursive search.
func FindRequestedFile(identifier string, searchDirs []string) []string {
	// Return no candidates of no search dir os specified.
	if len(searchDirs) == 0 {
		return []string{}
	}

	// List of possible files
	absolutePathCandidates := make([]string, 0)

	for _, directory := range searchDirs {
		if entries, err := os.ReadDir(directory); err == nil {
			for _, entry := range entries {
				// Filter out the dirs
				if entry.IsDir() {
					continue
				}

				// fmt.Printf("%q %q\n", directory, entry.Name())

				// Filter out the ones doesn't the name.
				// Instead of using a RegEx just check for substring.
				if !strings.Contains(entry.Name(), identifier) {
					// fmt.Printf("%q does not contain %q\n", entry.Name(), identifier)
					continue
				}

				fmt.Printf("%q does contain %q\n", entry.Name(), identifier)

				candidatePathString := filepath.Join(directory, entry.Name())
				// fmt.Printf("candidatePathString: %v\n", candidatePathString)
				absolutePathCandidates = append(absolutePathCandidates, candidatePathString)
			}
		} else {
			fmt.Printf("WARNING: cannot read from one of the search directories %q\n", directory)
		}
	}

	fmt.Printf("absolutePathCandidates: %v\n", absolutePathCandidates)

	return absolutePathCandidates
}

func TransferFileContent(sourcePath, destPath string, deleteOriginal, overwrite bool) error {
	if !overwrite && fileExists(destPath) {
		return fmt.Errorf("cannot create pdf, file already exists")
	}
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to dest from source: %v", err)
	}

	if deleteOriginal {
		err = os.Remove(sourcePath)
	}

	if err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// The actual job. This is why we are here.
// Check for override
func CreateBibTemplate(templateFile, bibDir, candidatePath, identifier string, overwrite bool, tags []string, status string) error {
	// Read the template from the template file
	tmp, err := template.ParseFiles(templateFile)

	if err != nil {
		return err
	}

	bibFilePath := filepath.Join(bibDir, fmt.Sprintf("%s.md", identifier))

	if !overwrite && fileExists(bibFilePath) {
		return fmt.Errorf("cannot create bib note, file already exists")
	}

	// Create the bib note in the bib note folder.
	bibNote, err := os.Create(bibFilePath)

	if err != nil {
		return err
	}

	defer bibNote.Close()

	// Use go templates to apply the template, the tags etc.
	return tmp.Execute(bibNote, BibNote{
		Date:       time.Now().Format("02-01-2006"),
		Identifier: identifier,
		Tags:       tags,
		Status:     status,
	})
}

// Note: Do some sanitization on the filename! like Remove unncessary capitilzation.
// You can also take the name with pandoc.
// But for now leave it as it is.
func CalculateDestinationPath(candidatePath, pdfDir string) string {
	filename := filepath.Base(candidatePath)
	return filepath.Join(pdfDir, filename)
}

// Just check if it ends with .pdf
func IsAPDFFile(path string) bool {
	return strings.Replace(strings.ToLower(filepath.Ext(path)), ".", "", -1) == "pdf"
}
