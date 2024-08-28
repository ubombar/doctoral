package doctoral

import "fmt"

// Output the help message to the screen.
func DisplayHelp() {
	msg := `Command is not recognised. Here are the available commands:
	bib: Bib Note related functionality.
		add: Adds a new Bib Note to the Vault.

You can also set these environment variables.

export DOCTORAL_PDF_DIR=/home/ubombar/Documents/Projects/Doctoral/doctoral/test/pdfs
export DOCTORAL_SEARCH_DIRS=/home/ubombar/Documents/Projects/Doctoral/doctoral/test/searchdir
export DOCTORAL_TEMPLATE_FILE=/home/ubombar/Documents/Projects/Doctoral/doctoral/test/template/mylittletemplate.md
export DOCTORAL_BIB_DIR=/home/ubombar/Documents/Projects/Doctoral/doctoral/test/bibnotes


for help type: doctoral help`

	fmt.Printf("%v\n", msg)
}
