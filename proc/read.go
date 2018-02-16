package proc

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

//ReadMarkdown returns text and title
func ReadMarkdown(r io.Reader) ([]byte, string, error) {
	buf := &bytes.Buffer{}
	title := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		str := scanner.Text()
		buf.WriteString(str)
		buf.WriteString("\n")
		str = strings.TrimLeft(str, " \t")
		if strings.HasPrefix(str, "# ") {
			title = strings.TrimLeft(str[2:], " \t")
		}
	}
	return buf.Bytes(), title, scanner.Err()
}
