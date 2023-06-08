package xmlparser

import (
	"encoding/xml"
	"io/ioutil"
)

func ParseXMLFile(filePath string) (*Unload, error) {
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
