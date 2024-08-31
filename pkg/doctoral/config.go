package doctoral

import (
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

// Name of the
const DOCTORALCONFIG = "DOCTORALCONFIG"

func GetDefaultConfigPath() string {
	homeDir := "~"
	absHomeDir, err := os.UserHomeDir()
	if err == nil {
		homeDir = absHomeDir
	}
	return filepath.Join(homeDir, ".doctoral", "config")
}

// Trys to get the doctoral config file path from the env variables.
// If it fails, then returns the default path.
func GetConfigPathOrDefault() string {
	configPath, isSet := os.LookupEnv(DOCTORALCONFIG)

	if !isSet {
		configPath = GetDefaultConfigPath()
	}

	return configPath
}

type Config struct {
	// The template file will be used to generate the Bib Note under the BibNotesDirectory
	TemplateFile string `yaml:"templateFile"`

	// Search directories for search to be performed. The PDF file will be transfered from
	// here to the PDFDirectory
	SearchDirectories []string `yaml:"searchDirectories"`

	// Where to put the Bib Note generated
	BibNotesDirectory string `yaml:"bibNotesDirectory"`

	// Where to transfer the new pdf file
	PDFDirectory string `yaml:"pdfDirectory"`

	// Will the PDF file be overwritten if it already exist
	OverwritePDFFiles bool `yaml:"overwritePDFFiles"`

	// Will the Bib Note file be overwritten if it already exist
	OverwriteBibNoteFiles bool `yaml:"overwriteBibNoteFiles"`

	// Deletes the original pdf file after copying it to the PDF directory
	DeleteAfterCopyingPDFs bool `yaml:"deleteAfterCopyingPDFs"`

	// Put a ! symbol in the begining of the PDFs are embedded instead of linked
	EmbedPDFs bool `yaml:"embedPDFs"`

	// Default tags to append the given tags by the command line flags
	DefaultTags []string `yaml:"defaultTags"`

	// Default status
	DefaultStatus string `yaml:"defaultStatus"`

	// The search regex
	DefaultSearchRegex string `yaml:"defaultSearchRegex"`
}

func NewConfigWithDefaultValues() *Config {
	return &Config{
		TemplateFile:           "~/Documents/template.md", // possibly doesn't exist
		SearchDirectories:      []string{},
		BibNotesDirectory:      "~/Documents",
		PDFDirectory:           "~/Documents",
		OverwritePDFFiles:      false,
		OverwriteBibNoteFiles:  false,
		DeleteAfterCopyingPDFs: true,
		EmbedPDFs:              false,
		DefaultTags:            []string{"#type/bibnote", "#topic/"},
		DefaultStatus:          "#status/waiting",
		DefaultSearchRegex:     ".*\\.pdf",
	}
}

// Trys to read from the config file. If it doesn't exist retuns an error.
// Note that if a field is not present in the yaml, it is pupulated by the
// default values by NewConfigWithDefaultValues function.
func ReadFromConfig(configPath string) (*Config, error) {
	configBytes, err := os.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	defaultConfig := NewConfigWithDefaultValues()
	err = yaml.Unmarshal(configBytes, defaultConfig)

	if err != nil {
		return nil, err
	}

	return defaultConfig, nil
}

// Creates a new config path and pupulates it with default values. Then
// returns it.
// For example the default config path is, $HOME/.doctoral/config
func CreateNewConfig(configPath string) (*Config, error) {
	err := os.MkdirAll(filepath.Dir(configPath), 0755)

	// Here is the default values of the ConfigFile
	defaultConfig := NewConfigWithDefaultValues()

	if err != nil {
		return nil, err
	}

	defaultConfigBytes, err := yaml.Marshal(defaultConfig)

	if err != nil {
		return nil, err
	}

	err = os.WriteFile(configPath, defaultConfigBytes, 0770)

	if err != nil {
		return nil, err
	}

	return defaultConfig, nil
}
