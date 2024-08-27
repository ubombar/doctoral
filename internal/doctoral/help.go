package doctoral

import "fmt"

// Output the help message to the screen.
func DisplayHelp() {
	msg := `Command is not recognised. Here are the available commands:
	bib: Bib Note related functionality.
		add: Adds a new Bib Note to the Vault.

for help type: doctoral help`

	fmt.Printf("%v\n", msg)
}
