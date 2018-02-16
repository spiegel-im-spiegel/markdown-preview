package proc

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/google/go-github/github"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spiegel-im-spiegel/markdown-preview/options"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

//Run Processing markdown file
func Run(r io.Reader, opt *options.Options) (io.Reader, error) {
	md, title, err := ReadMarkdown(r)
	if err != nil {
		return nil, err
	}

	var html []byte
	switch opt.ProcType() {
	case options.Blackfriday:
		html, err = renderWIthBlackfriday(md, opt)
		if err != nil {
			return nil, err
		}
	case options.GitHub:
		html, err = renderWIthGitHub(md)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unknown Processor")
	}

	if opt.Sanitizing() {
		html = bluemonday.UGCPolicy().SanitizeBytes(html)
	}

	if opt.PageFlag() {
		page := NewPage(html, opt.CSS(), title)
		html, err = page.Render()
		if err != nil {
			return nil, err
		}
	}

	return bytes.NewReader(html), nil
}

func renderWIthBlackfriday(md []byte, opt *options.Options) ([]byte, error) {
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: opt.HTMLFlags(), FootnoteReturnLinkContents: "&#x23CE;"})
	return blackfriday.Run(md, blackfriday.WithExtensions(opt.Extensions()), blackfriday.WithRenderer(renderer)), nil
}

func renderWIthGitHub(md []byte) ([]byte, error) {
	client := github.NewClient(nil)
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}
	body, _, err := client.Markdown(context.Background(), string(md), opt)
	return []byte(body), err
}
