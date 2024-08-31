package doctoral

import (
	"os"
	"path/filepath"
	"strings"
)

// This method is guaranteed to resolve the absolute path unlike filepath
// module which joins the pwd and the given path even though it starts
// with "~". For example "~/.doctoral/config" would result in a something
// like this "/home/jondoe/Development/myproject/~/.doctoral/config". This
// is clearly wrong and you need a syscall to resulve it correctly. This function
// does that.
// However, there is a possibiliy of an error thus it returns an error.
func GetAbsolutePath(path string) (string, error) {
	trimmedPath := strings.TrimSpace(path)
	// if strings has Prefix ~ then do the syscall, otherwise don't

	if strings.HasPrefix(trimmedPath, "~") {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		// Discard the prefix ~, and join it with the homedir. Then clean the result.
		absolutePath := filepath.Join(homeDir, trimmedPath[1:])
		return filepath.Clean(absolutePath), nil
	}

	return filepath.Abs(path)
}
