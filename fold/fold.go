package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

const ENDL = '\n'

func fold(b []byte, width int) []byte {
	runes := bytes.Runes(b)
	if len(runes) < width {
		return b
	} else {
		lineBreak := bytes.IndexRune(b, ENDL)
		if len(bytes.Runes(b[:lineBreak])) < width {
			lineLen := lineBreak + utf8.RuneLen(ENDL)
			return append(b[0:lineLen], fold(b[lineLen:], width)...)
		} else {
			curLine := []byte(string(append(runes[0:width], '\n')))
			nextStart := len(curLine) - utf8.RuneLen(runes[width])
			return append(curLine, fold(b[nextStart:], width)...)
		}
	}
}

func main() {
	flag.Parse()
	fileCount := len(flag.Args())
	var input []byte
	if fileCount > 0 {
		for i := 0; i < fileCount; i++ {
			input, _ = ioutil.ReadFile(flag.Args()[i])

		}
	} else {
		input, _ = ioutil.ReadAll(os.Stdin)
	}
	result := string(fold(input, 60))
	fmt.Print(result)
}
