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

type HeaderFooter struct {
	XMLName      xml.Name `xml:"sp_header_footer"`
	ClientScript string   `xml:"client_script"`
	Css          string   `xml:"css"`
	Script       string   `xml:"script"`
	Template     string   `xml:"template"`
	OptionSchema string   `xml:"option_schema"`
	Link         string   `xml:"link"`
}

type RecordUpdate struct {
	XMLName      xml.Name     `xml:"record_update"`
	Widget       Widget       `xml:"sp_widget"`
	HeaderFooter HeaderFooter `xml:"sp_header_footer"`
}

type WidgetContent interface {
	GetClientScript() string
	GetCss() string
	GetScript() string
	GetTemplate() string
	GetOptionSchema() string
	GetLink() string
}

func (w Widget) GetClientScript() string { return w.ClientScript }
func (w Widget) GetCss() string          { return w.Css }
func (w Widget) GetScript() string       { return w.Script }
func (w Widget) GetTemplate() string     { return w.Template }
func (w Widget) GetOptionSchema() string { return w.OptionSchema }
func (w Widget) GetLink() string         { return w.Link }

func (h HeaderFooter) GetClientScript() string { return h.ClientScript }
func (h HeaderFooter) GetCss() string          { return h.Css }
func (h HeaderFooter) GetScript() string       { return h.Script }
func (h HeaderFooter) GetTemplate() string     { return h.Template }
func (h HeaderFooter) GetOptionSchema() string { return h.OptionSchema }
func (h HeaderFooter) GetLink() string         { return h.Link }

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

			widgetDirPath := filepath.Join(dirPath, script.Name)
			if err := os.MkdirAll(widgetDirPath, 0755); err != nil {
				return err
			}

			jsContent := doWidgetOperation(recordUpdate.Widget)
			headerContent := doWidgetOperation(recordUpdate.HeaderFooter)

			if script.Type == "Widget" {
				for key, value := range jsContent {
					fileName := widgetFileTypes[key]
					filePath := filepath.Join(widgetDirPath, fileName)
					if err := ioutil.WriteFile(filePath, []byte(value), 0644); err != nil {
						return err
					}
				}
			}

			if script.Type == "Header | Footer" {
				for key, value := range headerContent {
					fileName := widgetFileTypes[key]
					filePath := filepath.Join(widgetDirPath, fileName)
					if err := ioutil.WriteFile(filePath, []byte(value), 0644); err != nil {
						return err
					}
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

func doWidgetOperation(widgetContent WidgetContent) map[string]string {
	content := map[string]string{
		"client_script": extractCDATA(widgetContent.GetClientScript()),
		"css":           extractCDATA(widgetContent.GetCss()),
		"script":        extractCDATA(widgetContent.GetScript()),
		"template":      extractCDATA(widgetContent.GetTemplate()),
		"option_schema": extractCDATA(widgetContent.GetOptionSchema()),
		"link":          extractCDATA(widgetContent.GetLink()),
	}
	return content
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
