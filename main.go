package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Unload struct {
	XMLName    xml.Name    `xml:"unload"`
	UpdateSets []UpdateSet `xml:"sys_remote_update_set"`
	XMLScripts []XMLScript `xml:"sys_update_xml"`
}

type UpdateSet struct {
	Action          string `xml:"action,attr"`
	Application     string `xml:"application"`
	ApplicationName string `xml:"application_name"`
}

type XMLScript struct {
	Action      string `xml:"action"`
	Application string `xml:"application"`
	Name        string `xml:"name"`
	Type        string `xml:"type"`
	Payload     string `xml:"payload"`
}

type Widget struct {
	XMLName      xml.Name `xml:"sp_widget"`
	ClientScript string   `xml:"client_script"`
	Css          string   `xml:"css"`
	Script       string   `xml:"script"`
	Template     string   `xml:"template"`
	OptionSchema string   `xml:"option_schema"`
	Link         string   `xml:"link"`
}

type RecordUpdate struct {
	XMLName xml.Name `xml:"record_update"`
	Widget  Widget   `xml:"sp_widget"`
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

func extractCDATA(content string) string {
	jsContent := strings.TrimSpace(content)

	jsCodeStart := strings.Index(jsContent, "<![CDATA[")
	jsCodeEnd := strings.LastIndex(jsContent, "]]>")
	if jsCodeStart != -1 && jsCodeEnd != -1 {
		jsContent = jsContent[jsCodeStart+len("<![CDATA[") : jsCodeEnd]
	}

	return jsContent
}

func createDirectoryStructureAndFiles(unload *Unload, outputDir string) error {
	widgetFileTypes := map[string]string{
		"client_script": "client-script.js",
		"css":           "style.scss",
		"script":        "script.js",
		"template":      "template.html",
		"option_schema": "options.json",
		"link":          "link.js",
	}

	for _, script := range unload.XMLScripts {
		if shouldExclude(script.Type) {
			continue
		}

		dirPath := filepath.Join(outputDir, script.Type)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}

		if script.Type == "Widget" || script.Type == "Header | Footer" {
			var recordUpdate RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse widget: %v\n", err)
				continue
			}

			// Create directory for the widget
			widgetDirPath := filepath.Join(dirPath, script.Name)
			if err := os.MkdirAll(widgetDirPath, 0755); err != nil {
				return err
			}

			widget := recordUpdate.Widget
			jsContent := map[string]string{
				"client_script": extractCDATA(widget.ClientScript),
				"css":           extractCDATA(widget.Css),
				"script":        extractCDATA(widget.Script),
				"template":      extractCDATA(widget.Template),
				"option_schema": extractCDATA(widget.OptionSchema),
				"link":          extractCDATA(widget.Link),
			}

			for key, value := range jsContent {
				fileName := widgetFileTypes[key]
				filePath := filepath.Join(widgetDirPath, fileName)
				if err := ioutil.WriteFile(filePath, []byte(value), 0644); err != nil {
					return err
				}
			}

		} else {
			fileName := fmt.Sprintf("%s.js", strings.ToLower(script.Name))
			filePath := filepath.Join(dirPath, fileName)
			jsContent := extractCDATA(script.Payload)
			if err := ioutil.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

func shouldExclude(fileType string) bool {
	if fileType == "System Property" ||
		fileType == "Access Roles" ||
		fileType == "Dictionary" ||
		fileType == "Field Label" ||
		fileType == "Form Layout" ||
		fileType == "HTTP Method" ||
		fileType == "REST Message" ||
		fileType == "Scripted REST API" ||
		fileType == "Scripted REST Version" {
		return true
	}
	return false
}

func main() {
	xmlFilePath := os.Args[1]
	outputDir := os.Args[2]

	unload, err := parseXMLFile(xmlFilePath)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		os.Exit(1)
	}

	err = createDirectoryStructureAndFiles(unload, outputDir)
	if err != nil {
		fmt.Println("Error creating directory structure and files:", err)
		os.Exit(1)
	}

	fmt.Println("Directories and files successfully created in", outputDir)
}
