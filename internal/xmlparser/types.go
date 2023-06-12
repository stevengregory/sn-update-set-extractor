package xmlparser

import "encoding/xml"

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
	TargetName  string `xml:"target_name"`
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

type FixScript struct {
	Description string `xml:"description"`
}

type RecordUpdate struct {
	XMLName      xml.Name     `xml:"record_update"`
	TargetName   string       `xml:"target_name"`
	Widget       Widget       `xml:"sp_widget"`
	HeaderFooter HeaderFooter `xml:"sp_header_footer"`
	FixScript    FixScript    `xml:"sys_script_fix"`
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
