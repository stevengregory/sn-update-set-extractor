package xmlparser

import "strings"

func ExtractCDATA(content string) string {
	jsContent := strings.TrimSpace(content)

	jsCodeStart := strings.Index(jsContent, "<![CDATA[")
	jsCodeEnd := strings.LastIndex(jsContent, "]]>")
	if jsCodeStart != -1 && jsCodeEnd != -1 {
		jsContent = jsContent[jsCodeStart+len("<![CDATA[") : jsCodeEnd]
	}

	return jsContent
}
