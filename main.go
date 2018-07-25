package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	Type      string
	Delimiter string
}

func read(r io.Reader, size int64) (chan []byte, context.Context) {
	in := make(chan []byte, size)
	ctx, cancel := context.WithCancel(context.Background())
	go func(r io.Reader) {
		defer cancel()
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			in <- sc.Bytes()
		}
		if err := sc.Err(); err != nil {
			log.Fatalf("scan failed: %v", err)
		}
	}(r)
	return in, ctx
}

func run() error {
	in, ctx := read(os.Stdin, 10)
	for {
		select {
		case b := <-in:
			line := string(b)
			// csv
			elems := strings.Split(line, ",")
			err := check(elems)
			if err != nil {
				return fmt.Errorf("failed check: %v", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("kpass error exit: ", err)
	}
}
