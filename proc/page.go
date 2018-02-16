package proc

import (
	"bytes"
	"io"
	"text/template"

	"github.com/spiegel-im-spiegel/markdown-preview/data"
)

//PageData is field set for HTML template
type PageData struct {
	Title string
	CSS   string
	Body  string
}

//NewPage returns instance of NewPage
func NewPage(body []byte, css, title string) *PageData {
	return &PageData{Title: title, CSS: css, Body: string(body)}
}

//Render returns HTML page text with template
func (p *PageData) Render() ([]byte, error) {
	f, err := data.Assets.Open("/template.html")
	if err != nil {
		return nil, err
	}
	tmpData := &bytes.Buffer{}
	io.Copy(tmpData, f)

	t, err := template.New("Markdown Processing").Parse(tmpData.String())
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, p)
	return buf.Bytes(), err
}
