package doctoral

import (
	"os"
	"strings"
)

const (
	DOCTORAL_SEARCH_DIRS = "DOCTORAL_SEARCH_DIRS"
	DOCTORAL_PDF_DIR     = "DOCTORAL_PDF_DIR"
	DOCTORAL_BIB_DIR     = "DOCTORAL_BIB_DIR"
)

func GetDefaultPDFDir() string {
	return os.Getenv(DOCTORAL_PDF_DIR)
}

func GetDefaultBibDir() string {
	return os.Getenv(DOCTORAL_BIB_DIR)
}

func GetDefaultSearchDirs() []string {
	dirs := os.Getenv(DOCTORAL_SEARCH_DIRS)

	// Do some sanitization ...
	dirArray := strings.Split(dirs, ":")

	return dirArray
}
