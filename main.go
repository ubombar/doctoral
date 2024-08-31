/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/ubombar/doctoral/internal/doctoral"
	doctoralpkg "github.com/ubombar/doctoral/pkg/doctoral"
)

func main() {
	// cmd.Execute()

	m := doctoral.NewSimpleMenu()

	options := []doctoralpkg.Document{
		*doctoralpkg.NewDocumentWithoutError("~/hello 1.pdf"),
		*doctoralpkg.NewDocumentWithoutError("~/hello 2.pdf"),
		*doctoralpkg.NewDocumentWithoutError("~/hello 3.pdf"),
		*doctoralpkg.NewDocumentWithoutError("~/hello 4.pdf"),
	}

	selectedDocuments := m.GetChoices(options)

	fmt.Printf("selectedDocuments: %v\n", selectedDocuments)
}
