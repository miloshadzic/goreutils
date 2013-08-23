package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
)

type Score struct {
	WordCount, LineCount, CharCount, ByteCount int
	FileName                                   string
}

func (self *Score) Add(other *Score) *Score {
	self.WordCount += other.WordCount
	self.LineCount += other.LineCount
	self.CharCount += other.CharCount
	self.ByteCount += other.ByteCount
	return self
}

var w, l, m, c bool

func init() {
	flag.BoolVar(&l, "l", false, "Count the number of lines")
	flag.BoolVar(&w, "w", false, "Count the number of words")
	flag.BoolVar(&c, "c", false, "Count the number of bytes")
	flag.BoolVar(&m, "m", false, "Count the number of characters")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: wc [-clmw] [file ...]\n\n")
		flag.VisitAll(usage)
	}
}

func usage(f *flag.Flag) {
	fmt.Printf("  -%s\t%s\n", f.Name, f.Usage)
}

func words(str *string) int {
	scanner := bufio.NewScanner(strings.NewReader(*str))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func lines(str *string) int {
	scanner := bufio.NewScanner(strings.NewReader(*str))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func count_file(fileName string) *Score {
	bytes, _ := ioutil.ReadFile(fileName)
	score := count(bytes)
	score.FileName = fileName
	return score
}

func count(b []byte) *Score {
	str := string(b)
	return &Score{
		LineCount: lines(&str),
		WordCount: words(&str),
		CharCount: len(bytes.Runes(b)),
		ByteCount: len(b)}
}

func Printout(score *Score) {
	var out string
	if !(l || w || c || m) {
		out = fmt.Sprintf("%d\t%d\t%d\t %s",
			score.LineCount, score.WordCount, score.ByteCount, score.FileName)
	} else {
		if l {
			out = fmt.Sprintf("%d\t", score.LineCount)
		}
		if w {
			out = fmt.Sprintf("%s%d\t", out, score.WordCount)
		}
		if m {
			out = fmt.Sprintf("%s%d\t", out, score.CharCount)
		}
		if c && !m {
			out = fmt.Sprintf("%s%d\t", out, score.ByteCount)
		}
		out = fmt.Sprintf("%s %s", out, score.FileName)
	}
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, out)
	w.Flush()
}

func main() {
	flag.Parse()
	fileCount := len(flag.Args())
	var score *Score
	if fileCount == 1 {
		score = count_file(flag.Args()[0])
	} else if fileCount > 1 {
		score = &Score{WordCount: 0, LineCount: 0, ByteCount: 0, FileName: "total"}
		for i := 0; i < fileCount; i++ {
			s := count_file(flag.Args()[i])
			score.Add(s)
			Printout(s)
		}
	} else {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		score = count(bytes)
	}
	Printout(score)
	os.Exit(0)
}
