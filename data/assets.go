package data

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9ebc522023c36de07ec03c8945f46e909e76684b = "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n  <meta name=\"generator\" content=\"markdown-preview\">{{with .CSS }}\n  <link rel=\"stylesheet\" type=\"text/css\" href=\"{{ . }}\">{{ end }}\n  <title>{{ .Title }}</title>\n</head>\r\n<body>\r\n\n{{ .Body }}\n\n</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518810466, 1518810466154181800),
		Data:     nil,
	}, "/template.html": &assets.File{
		Path:     "/template.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1518810775, 1518810775444226300),
		Data:     []byte(_Assets9ebc522023c36de07ec03c8945f46e909e76684b),
	}}, "")
