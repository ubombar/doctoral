package doctoral

import (
	"fmt"
	"log"

	"github.com/pkg/term"
)

type Menu interface {
	// Asks use which of the files they want to choose. Returns the selected ones.
	GetChoices(options []Document) []Document
}

const (
	KEY_UP     string = "\x1b[A"
	KEY_DOWN   string = "\x1b[B"
	KEY_SPACE  string = " \x00\x00"
	KEY_RETURN string = "\r\x00\x00"
	KEY_Q      string = "q\x00\x00"
)

type simpleMenu struct {
	Menu
}

func NewSimpleMenu() Menu {
	return &simpleMenu{}
}

// Prompts user for a list of files to pick.
func (m *simpleMenu) GetChoices(options []Document) []Document {
	exit := false
	cursor := 0
	selectedIndicies := map[int]bool{}

	// The main loop
	for !exit {
		m.clearTerminal()
		fmt.Println("UP/DOWN arrows to move, SPACE to select, ENTER to end, Q to quit.")

		// Check if there are no available options to select, exit with an empty list
		if len(options) == 0 {
			fmt.Println("\tThere is noting to display")
			return []Document{}
		}

		for i, option := range options {
			optString := m.stringifyOption(i, &option, selectedIndicies[i], cursor == i)
			fmt.Printf("\t%v\n", optString)
		}

		userInput := m.getInput()

		switch userInput {
		case KEY_UP:
			cursor = (cursor + len(options) - 1) % len(options)
		case KEY_DOWN:
			cursor = (cursor + len(options) + 1) % len(options)
		case KEY_SPACE:
			selectedIndicies[cursor] = !selectedIndicies[cursor] // Toggle
		case KEY_Q:
			return []Document{} // Return empty list
		case KEY_RETURN:
			exit = true
		}
	}
	selectedOptions := make([]Document, 0)

	for index, selected := range selectedIndicies {
		if selected {
			selectedOptions = append(selectedOptions, options[index])
		}
	}

	return selectedOptions
}

func (m simpleMenu) clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func (m *simpleMenu) stringifyOption(index int, opt *Document, selected, cursor bool) string {
	cursorString := " "
	selectedString := " "

	if cursor {
		cursorString = ">"
	}
	if selected {
		selectedString = "*"
	}

	return fmt.Sprintf("%s%s [%d]: %s", cursorString, selectedString, index, opt.FileName)
}

func (m *simpleMenu) getInput() string {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	readBytes := make([]byte, 3)
	_, err = t.Read(readBytes)

	if err != nil {
		return ""
	}

	t.Restore()
	t.Close()

	return string(readBytes)
}
