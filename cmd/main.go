package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NullHypothesis/cues"
)

func estimate(r io.Reader, c *cues.Counter[string]) uint64 {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c.Feed(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return c.Estimate()
}

func main() {
	var (
		path    string
		r       io.Reader
		err     error
		bufSize int
	)

	flag.IntVar(&bufSize, "buf", 1024, "Buffer size. Larger buffers yield more accurate estimates.")
	flag.StringVar(&path, "path", "", "File containing newline-separated words.")
	flag.Parse()

	if path != "" {
		r, err = os.Open(path)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer r.(*os.File).Close()
	} else {
		r = os.Stdin
		log.Println("Reading from stdin.")
	}

	fmt.Printf("Approximately %d elements in stream.\n",
		estimate(r, cues.New[string](bufSize)))
}
