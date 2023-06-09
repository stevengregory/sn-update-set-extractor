package fileops

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stevengregory/sn-update-set-extractor/internal/xmlparser"
)

func CreateDirsAndFiles(unload *xmlparser.Unload, outputDir string) error {
	widgetFileTypes := getWidgetFileTypes()

	for _, script := range unload.XMLScripts {
		if shouldExclude(script.Type) {
			continue
		}

		dirPath := filepath.Join(outputDir, script.Type)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}

		if script.Type == "Widget" || script.Type == "Header | Footer" {
			var recordUpdate xmlparser.RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse widget: %v\n", err)
				continue
			}

			widgetDirPath := filepath.Join(dirPath, script.Name)
			if err := os.MkdirAll(widgetDirPath, 0755); err != nil {
				return err
			}

			content := getWidgetContentType(script.Type, recordUpdate)

			for key, value := range content {
				fileName := widgetFileTypes[key]
				filePath := filepath.Join(widgetDirPath, fileName)
				if err := ioutil.WriteFile(filePath, []byte(value), 0644); err != nil {
					return err
				}
			}
		} else {
			fileName := fmt.Sprintf("%s.js", strings.ToLower(script.Name))
			filePath := filepath.Join(dirPath, fileName)
			jsContent := xmlparser.ExtractCDATA(script.Payload)
			if err := ioutil.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

func doWidgetOperation(widgetContent xmlparser.WidgetContent) map[string]string {
	content := map[string]string{
		"client_script": xmlparser.ExtractCDATA(widgetContent.GetClientScript()),
		"css":           xmlparser.ExtractCDATA(widgetContent.GetCss()),
		"script":        xmlparser.ExtractCDATA(widgetContent.GetScript()),
		"template":      xmlparser.ExtractCDATA(widgetContent.GetTemplate()),
		"option_schema": xmlparser.ExtractCDATA(widgetContent.GetOptionSchema()),
		"link":          xmlparser.ExtractCDATA(widgetContent.GetLink()),
	}
	return content
}

func excludedFileTypes() map[string]struct{} {
	return map[string]struct{}{
		"System Property":       {},
		"Access Roles":          {},
		"Dictionary":            {},
		"Field Label":           {},
		"Form Layout":           {},
		"HTTP Method":           {},
		"REST Message":          {},
		"Scripted REST API":     {},
		"Scripted REST Version": {},
	}
}

func getWidgetContentType(fileType string, recordUpdate xmlparser.RecordUpdate) map[string]string {
	var content map[string]string
	switch fileType {
	case "Widget":
		content = doWidgetOperation(recordUpdate.Widget)
	case "Header | Footer":
		content = doWidgetOperation(recordUpdate.HeaderFooter)
	}
	return content
}

func getWidgetFileTypes() map[string]string {
	return map[string]string{
		"client_script": "client-script.js",
		"css":           "style.scss",
		"script":        "script.js",
		"template":      "template.html",
		"option_schema": "options.json",
		"link":          "link.js",
	}
}

func shouldExclude(fileType string) bool {
	_, exclude := excludedFileTypes()[fileType]
	return exclude
}
