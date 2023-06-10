package main

import (
	"fmt"
	"os"

	"github.com/stevengregory/sn-update-set-extractor/internal/fileops"
	"github.com/stevengregory/sn-update-set-extractor/internal/xmlparser"
)

func main() {
	dataDir := "./data"
	outputDir := "./dist"

	files, err := os.ReadDir(dataDir)
	if err != nil {
		fmt.Println("Error reading data directory:", err)
		os.Exit(1)
	}

	for _, file := range files {
		xmlFilePath := dataDir + "/" + file.Name()

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
	}

	fmt.Println("Directories and files successfully created in", outputDir)
}
