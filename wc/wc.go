package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var count_words, count_bytes, count_lines bool

func init() {
	flag.BoolVar(&count_lines, "l", true, "Count the number of lines")
	flag.BoolVar(&count_words, "w", true, "Count the number of words")
	flag.BoolVar(&count_bytes, "c", true, "Count the number of bytes")
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
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
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
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return count
}

func main() {
	flag.Parse()

	var bytes []byte
	var err error

	if len(flag.Args()) == 0 {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(flag.Args()[0])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "File couldn't be opened.", err)
	}
	str := string(bytes)
	if !(count_words && count_lines && count_bytes) {
		fmt.Printf("%d\t%d\t%d\n", lines(&str), words(&str), len(bytes))
		fmt.Println("yeah")
	} else {
		if count_lines {
			fmt.Printf("%d\t", lines(&str))
		}
		if count_words {
			fmt.Printf("%d\t", words(&str))
		}
		if count_bytes {
			fmt.Printf("%d\t", len(bytes))
		}
		fmt.Println()
	}
}
