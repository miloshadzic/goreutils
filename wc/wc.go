package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Score struct {
	FileName                        string
	WordCount, LineCount, ByteCount int
}

func (self *Score) Add(other *Score) *Score {
	self.WordCount += other.WordCount
	self.LineCount += other.LineCount
	self.ByteCount += other.ByteCount
	return self
}

var w, l, c bool

func init() {
	flag.BoolVar(&l, "l", true, "Count the number of lines")
	flag.BoolVar(&w, "w", true, "Count the number of words")
	flag.BoolVar(&c, "c", true, "Count the number of bytes")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: wc [-clmw] file\n\n")
		flag.VisitAll(usage)
	}
}

func usage(f *flag.Flag) {
	fmt.Printf("  -%s\t%s\n", f.Name, f.Usage)
}

func words(str *string) int {
	scanner := bufio.NewScanner(strings.NewReader(*str))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
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

func count(bytes []byte) *Score {
	str := string(bytes)
	return &Score{
		LineCount: lines(&str),
		WordCount: words(&str),
		ByteCount: len(bytes)}
}

func Printout(score *Score) {
	if !(w && l && c) {
		fmt.Printf("%d\t%d\t%d ", score.LineCount, score.WordCount, score.ByteCount)
	} else {
		if l {
			fmt.Printf("%d\t ", score.LineCount)
		}
		if w {
			fmt.Printf("%d\t ", score.WordCount)
		}
		if c {
			fmt.Printf("%d\t ", score.ByteCount)
		}
		fmt.Printf("%s\n", score.FileName)
	}
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
}
