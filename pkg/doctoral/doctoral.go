package doctoral

import "path/filepath"

type Document struct {
	AbsolutePath string
	FileName     string
	Extension    string
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
