package doctoral

import "fmt"

// Output the help message to the screen.
func DisplayHelp() {
	msg := `Command is not recognised. Here are the available commands:
	addbib: Adds a Bib Note to the vault.
	addhard: Adds a Hard Note to the vaul.
	`

	fmt.Printf("%v\n", msg)
}
