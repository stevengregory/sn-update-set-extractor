package main

import (
	"fmt"
	"os"

	"github.com/stevengregory/sn-update-set-extractor/internal/fileops"
	"github.com/stevengregory/sn-update-set-extractor/internal/xmlparser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the XML file path and the output directory as arguments.")
		os.Exit(1)
	}

	xmlFilePath := os.Args[1]
	outputDir := os.Args[2]

	unload, err := xmlparser.ParseXMLFile(xmlFilePath)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		os.Exit(1)
	}

	err = fileops.CreateDirsAndFiles(unload, outputDir)
	if err != nil {
		fmt.Println("Error creating directory structure and files:", err)
		os.Exit(1)
	}

	fmt.Println("Directories and files successfully created in", outputDir)
}
