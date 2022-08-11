package main

import(
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
		<title>Markdown Preview Tool</title>
	</head>
<body>
`
footer = `
	</body>
	</html>
`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()
	if *filename = "" {
		flag.Usage()
		os.exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename String) 


