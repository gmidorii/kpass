package main

import (
	"fmt"
	"os"

	"github.com/midorigreen/kpass/gate"
)

var gaters = []gate.Gater{
	gate.Minimumer{},
}

func check(elems []string) error {
	result := make(chan gate.Result, 10)
	counter := 0
	for _, g := range gaters {
		c := gate.Condition{
			Elements: elems,
			Num:      3,
		}
		go g.Examine(c, elems[2], result)
		counter++
	}

	for {
		if counter == 0 {
			break
		}
		select {
		case r := <-result:
			if !r.OK {
				fmt.Fprintf(os.Stderr, "failed check: %v", r.Message)
			} else {
			}
			counter--
		}
	}

	return nil
}
