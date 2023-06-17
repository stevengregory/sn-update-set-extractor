package fileops

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stevengregory/sn-update-set-extractor/internal/xmlparser"
)

func CreateDirsAndFiles(unload *xmlparser.Unload, outputDir string) error {
	widgetFileTypes := getWidgetFileTypes()

	for _, script := range unload.XMLScripts {
		if !shouldInclude(script.Type) {
			continue
		}

		dirPath := filepath.Join(outputDir, script.Type)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}

		fileExt := getFileExtForType(script.Type)
		fileName := fmt.Sprintf("%s.%s", script.TargetName, fileExt)
		filePath := filepath.Join(dirPath, fileName)

		if script.Type == "Widget" || script.Type == "Header | Footer" {
			var recordUpdate xmlparser.RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse widget: %v\n", err)
				continue
			}

			scriptName := script.TargetName

			widgetDirPath := filepath.Join(dirPath, scriptName)
			if err := os.MkdirAll(widgetDirPath, 0755); err != nil {
				return err
			}

			content := getWidgetContentType(script.Type, recordUpdate)

			for key, value := range content {
				fileName := widgetFileTypes[key]
				filePath := filepath.Join(widgetDirPath, fileName)
				if err := os.WriteFile(filePath, []byte(value), 0644); err != nil {
					return err
				}
			}
		} else if script.Type == "Fix Script" {
			var recordUpdate xmlparser.RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse fix script: %v\n", err)
				continue
			}

			jsContent := xmlparser.ExtractCDATA(recordUpdate.FixScript.Description)
			if err := os.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
				return err
			}
		} else if script.Type == "UI Action" {
			var recordUpdate xmlparser.RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse UI action: %v\n", err)
				continue
			}

			jsContent := xmlparser.ExtractCDATA(recordUpdate.UIAction.Script)
			if err := os.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
				return err
			}
		} else if script.Type == "Theme" {
			var recordUpdate xmlparser.RecordUpdate
			err := xml.Unmarshal([]byte(script.Payload), &recordUpdate)
			if err != nil {
				fmt.Printf("Failed to parse theme: %v\n", err)
				continue
			}

			scssContent := xmlparser.ExtractCDATA(recordUpdate.Theme.CssVariables)
			if err := os.WriteFile(filePath, []byte(scssContent), 0644); err != nil {
				return err
			}
		} else {
			jsContent := xmlparser.ExtractCDATA(script.Payload)
			if err := os.WriteFile(filePath, []byte(jsContent), 0644); err != nil {
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

func getFileExtForType(fileType string) string {
	switch fileType {
	case "Theme":
		return "scss"
	default:
		return "js"
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
		"client_script": "client_script.js",
		"css":           "style.scss",
		"script":        "server_script.js",
		"template":      "template.html",
		"option_schema": "option_schema.json",
		"link":          "link.js",
	}
}

func shouldInclude(fileType string) bool {
	_, include := supportedFileTypes()[fileType]
	return include
}

func supportedFileTypes() map[string]struct{} {
	return map[string]struct{}{
		"Business Rule":          {},
		"Client Script":          {},
		"Fix Script":             {},
		"Header | Footer":        {},
		"Script Include":         {},
		"Scripted REST Resource": {},
		"UI Action":              {},
		"UI Script":              {},
		"Theme":                  {},
		"Widget":                 {},
	}
}
