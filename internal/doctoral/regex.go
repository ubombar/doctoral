package doctoral

import (
	"fmt"
	"regexp"
)

type IdentifierType string

const (
	UNKNOWN IdentifierType = "unknown"
	URL     IdentifierType = "url"
	FILE    IdentifierType = "file"
	YOUTUBE IdentifierType = "youtube"
	ARXIV   IdentifierType = "arxiv"
)

const (
	// This represents a general URL, ADAM OLANA COK BILE
	URL_REGEX_STR = "^https?://(www\\.)?"

	// This represents a URL of youtube.com
	YOUTUBE_REGEX_STR = "^https?://(www\\.)?youtube.com"

	// This represents a URL of arxiv.org
	ARXIV_REGEX_STR = "^https?://(www\\.)?arxiv.org"

	// It can end with .pdf or not, for now just check if it is a file. I made this myself so probably it is shit.
	FILE_REGEX_STR = "^\\/?(\\.|\\.\\.|[a-zA-Z0-9\\.\\_\\-!?%&\\$@\\ ]+\\/)*[a-zA-Z0-9\\.\\_\\-!?%&\\$@\\ ]+$"
)

func GetTypeOfIdentifier(identifier string) IdentifierType {
	identifierBytes := []byte(identifier)
	// If it is a URL
	if match(URL_REGEX_STR, identifierBytes) {
		// Check for youtube
		// https://www.youtube.com/watch?v=6OPsH8PK7xM
		if match(YOUTUBE_REGEX_STR, identifierBytes) {
			return YOUTUBE
		}

		// Check fro arxiv
		// example: https://arxiv.org/abs/2408.14881
		if match(ARXIV_REGEX_STR, identifierBytes) {
			return ARXIV
		}

		// This case it tries to download the pdf in the next steps or put it as a blog
		// example: https://conservancy.umn.edu/server/api/core/bitstreams/49e5e823-2642-42eb-89d8-8ed962b0fc38/content
		return URL
	} else {
		// We dont know if it is a pdf or not so jsut return the media type as file.
		if match(FILE_REGEX_STR, identifierBytes) {
			return FILE
		}
	}

	return UNKNOWN
}

func match(regexString string, identifierBytes []byte) bool {
	if re, err := regexp.Compile(regexString); err == nil && re.Match(identifierBytes) {
		return true
	}
	return false
}

// Creates a regex from the string while sanitizing it. Soft means there can be other characters behind
// and after the string.
// NOT Implemented the sanitization part actually :)
// Add case insensitivity and removal of special characters.
func SanitizedSoftCompile(name string) (*regexp.Regexp, error) {
	return regexp.Compile(fmt.Sprintf(".*%s", name))
}
