package main

import (
	"fmt"
	"os"

	"github.com/perelin/pdf2cbz/internal/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pdf2cbz <pdf-file>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	fmt.Printf("Converting %s to CBZ format...\n", pdfPath)

	err := converter.ConvertPDFToCBZ(pdfPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully!")
}
