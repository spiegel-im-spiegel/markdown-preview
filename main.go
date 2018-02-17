package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/markdown-preview/facade"
)

//go:generate go-assets-builder -p data -s="/data/assets" -o data/assets.go data/assets/

func main() {
	facade.Execute(
		rwi.New(
			rwi.Reader(os.Stdin),
			rwi.Writer(os.Stdout),
			rwi.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
