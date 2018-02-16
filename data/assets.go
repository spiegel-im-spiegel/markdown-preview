package data

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets385a21e068f39970304cdfea5272ab235abe7fa9 = "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n  <meta name=\"generator\" content=\"markdown-preview\">{{with .CSS }}\n  <link rel=\"stylesheet\" type=\"text/css\" href=\"{{ . }}\">{{ end }}\n  <title>{{ .Title }}</title>\n</head>\r\n<body>\r\n\n{{ .Body }}\n\n</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template.html"}}, map[string]*assets.File{
	"/template.html": &assets.File{
		Path:     "/template.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1518810775, 1518810775444226300),
		Data:     []byte(_Assets385a21e068f39970304cdfea5272ab235abe7fa9),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518810466, 1518810466154181800),
		Data:     nil,
	}}, "")
