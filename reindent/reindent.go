package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var write = flag.Bool("w", false, "Write directly to file")

func main() {
	flag.Parse()
	files := flag.Args()
	var processWG sync.WaitGroup
	processWG.Add(len(files))
	for _, f := range files {
		go func(file string) {
			must(lint(file))
			processWG.Done()
		}(f)
	}

	processWG.Wait()
	fmt.Fprintln(os.Stderr, "👍")
}

func lint(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	var buf bytes.Buffer

	scanner := bufio.NewScanner(f)
	currInc := 0
	tempInc := 0
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		switch {
		case hasAnyPrefix(t, ")", "}", "]"):
			currInc--
		}
		buf.WriteString(strings.Repeat(" ", (currInc+tempInc)*2))
		buf.WriteString(t)
		buf.WriteRune('\n')
		tempInc = 0
		switch {
		case hasAnySuffix(t, "{", "[", "("):
			currInc++
		case hasAnySuffix(t, ",", ".", ":"):
			tempInc++
		}
	}
	f.Close()
	if scanner.Err() != nil {
		return scanner.Err()
	}
	if *write {
		ioutil.WriteFile(filename, buf.Bytes(), 0666)
	} else {
		_, err = buf.WriteTo(os.Stdout)
	}
	return err
}

func hasAnySuffix(s string, sxs ...string) bool {
	for _, sx := range sxs {
		if strings.HasSuffix(s, sx) {
			return true
		}
	}
	return false
}

func hasAnyPrefix(s string, sxs ...string) bool {
	for _, sx := range sxs {
		if strings.HasPrefix(s, sx) {
			return true
		}
	}
	return false
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
