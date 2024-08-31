package doctoral

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"
)

type BibNoteTemplateData struct {
	EmbeddingSymbol                  string
	Tags                             []string
	Date                             string
	MaterialFileNameWithoutExtension string
	MaterialFileName                 string
	Status                           string
}

type Document struct {
	AbsolutePath string
	FileName     string
	Extension    string
}

// Gets all the documents under the given directory.
func GetDocumentsUnderDirectory(directory string) ([]Document, error) {
	directoryAbsPath, err := GetAbsolutePath(directory)

	if err != nil {
		return []Document{}, err
	}

	dirEntry, err := os.ReadDir(directoryAbsPath)

	if err != nil {
		return []Document{}, err
	}

	documents := make([]Document, len(dirEntry))

	for i, entry := range dirEntry {
		temp, err := NewDocument(filepath.Join(directoryAbsPath, entry.Name()))

		if err != nil {
			return []Document{}, err
		}

		documents[i] = *temp
	}

	return documents, nil
}

func GetDocumentsUnderDirectories(directories []string, searchRegex string) ([]Document, error) {
	re, err := regexp.Compile(searchRegex)
	if err != nil {
		return []Document{}, err
	}
	docs := []Document{}

	for _, directory := range directories {
		batch, err := GetDocumentsUnderDirectory(directory)

		if err != nil {
			return []Document{}, err
		}

		for _, e := range batch {
			if re.Match([]byte(e.AbsolutePath)) {
				docs = append(docs, e)
			}
		}
	}

	return docs, nil
}

// Creates a new Document file representation.
func NewDocument(path string) (*Document, error) {
	absPath, err := GetAbsolutePath(path)

	if err != nil {
		return nil, err
	}

	return &Document{
		AbsolutePath: absPath,
		FileName:     filepath.Base(absPath),
		Extension:    filepath.Ext(absPath),
	}, nil
}

// Creates a new Document file representation.
func NewDocumentWithoutError(path string) *Document {
	doc, err := NewDocument(path)
	if err != nil {
		return nil
	}
	return doc
}

func (d Document) ExistOnDisk() bool {
	_, err := os.Stat(d.AbsolutePath)
	return !errors.Is(err, os.ErrNotExist)
}

func (d Document) ParentDirectoryAbsPAth() string {
	return filepath.Dir(d.AbsolutePath)
}

func (d Document) CopyToDirectory(directory string) error {
	destinationFilePathAbs, err := GetAbsolutePath(filepath.Join(directory, d.FileName))
	if err != nil {
		return err
	}
	return d.CopyToFile(destinationFilePathAbs)
}

func (d Document) CopyToFile(destinationFilePathAbs string) error {
	sourceFileStat, err := os.Stat(d.AbsolutePath)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", d.AbsolutePath)
	}

	source, err := os.Open(d.AbsolutePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFilePathAbs)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func (d Document) Delete() error {
	if !d.ExistOnDisk() {
		return errors.New("file does not exist, cannot delete non-existing file")
	}

	return os.Remove(d.AbsolutePath)
}

func (d Document) FileNameWithoutExt() string {
	s := strings.TrimSuffix(d.FileName, d.Extension)
	return s
}

// From the template and TemplateScheme, put the contents inside the file
func (d Document) TemplateContent(templateDocument Document, templateData BibNoteTemplateData) error {
	templateFileContent, err := os.ReadFile(templateDocument.AbsolutePath)
	if err != nil {
		return err
	}

	t, err := template.New(d.FileName).Parse(string(templateFileContent))

	if err != nil {
		return err
	}

	outputFile, err := os.Create(d.AbsolutePath)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	if err := t.Execute(outputFile, templateData); err != nil {
		return err
	}

	return nil
}

func NewTemplateData(config *Config, bibNote *Document, material *Document) BibNoteTemplateData {
	embeddingSymbol := ""

	if config.EmbedPDFs {
		embeddingSymbol = "!"
	}
	return BibNoteTemplateData{
		EmbeddingSymbol:                  embeddingSymbol,
		Tags:                             config.DefaultTags, // TODO: Just add default tags for now.
		Date:                             time.Now().Format("02-01-2006"),
		MaterialFileNameWithoutExtension: bibNote.FileNameWithoutExt(),
		MaterialFileName:                 material.FileName,
		Status:                           config.DefaultStatus,
	}
}
