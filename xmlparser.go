package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Unload struct {
	XMLName     xml.Name    `xml:"unload"`
	UpdateSets  []UpdateSet `xml:"sys_remote_update_set"`
	XMLScripts  []XMLScript `xml:"sys_update_xml"`
}

type UpdateSet struct {
	Action          string `xml:"action,attr"`
	Application     string `xml:"application"`
	ApplicationName string `xml:"application_name"`
}

type XMLScript struct {
	Action    string `xml:"action"`
	Application   string `xml:"application"`
	Name      string `xml:"name"`
	Type      string `xml:"type"`
	Payload   string `xml:"payload"`
}

func parseXMLFile(filePath string) (*Unload, error) {
	xmlData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var unload Unload
	err = xml.Unmarshal(xmlData, &unload)
	if err != nil {
		return nil, err
	}

	return &unload, nil
}

func extractJavaScriptCode(unload *Unload) []string {
	var jsCode []string

	for _, script := range unload.XMLScripts {
		jsCode = append(jsCode, script.Payload)
	}

	return jsCode
}

func createDirectoryStructureAndFiles(unload *Unload, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	for _, script := range unload.XMLScripts {

		// Skip creating directory for "System Property" type
		if script.Type == "System Property" {
			continue
		}

		dirPath := filepath.Join(outputDir, script.Type)

		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s.js", strings.ToLower(script.Name))
		filePath := filepath.Join(dirPath, fileName)

		jsContent := removeXMLTags(script.Payload)

		if err := ioutil.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
			return err
		}
	}

	return nil
}


func removeXMLTags(content string) string {
	// Remove XML tags using regular expressions
	xmlTagRegex := regexp.MustCompile(`<.*?>`)
	jsContent := xmlTagRegex.ReplaceAllString(content, "")

	// Remove leading and trailing whitespaces
	jsContent = strings.TrimSpace(jsContent)

	// Extract only the JavaScript code from the payload
	jsCodeStart := strings.Index(jsContent, "<![CDATA[")
	jsCodeEnd := strings.LastIndex(jsContent, "]]>")
	if jsCodeStart != -1 && jsCodeEnd != -1 {
		jsContent = jsContent[jsCodeStart+len("<![CDATA[") : jsCodeEnd]
	}

	return jsContent
}



func main() {
	xmlFilePath := os.Args[1]
	outputDir := os.Args[2]

	unload, err := parseXMLFile(xmlFilePath)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	jsCode := extractJavaScriptCode(unload)

	err = createDirectoryStructureAndFiles(unload, outputDir)
	if err != nil {
		fmt.Println("Error creating directory structure and files:", err)
		return
	}

	for _, code := range jsCode {
		fmt.Println(code)
	}
}
