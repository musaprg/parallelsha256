package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"
)

type Result struct {
	order int
	value string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: parallelsha256 <path_to_file>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
    defer f.Close()

	r := bufio.NewReader(f)

	lines := []string{}

	for {
		if l, err := r.ReadString('\n'); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			lines = append(lines, l)
		}
	}

	c := make(chan Result, len(lines))
	wg := &sync.WaitGroup{}

	for i, l := range lines {
		wg.Add(1)
		go func(i int, s string) {
			defer wg.Done()
			h := sha256.Sum256([]byte(s))
			c <- Result{order: i, value: hex.EncodeToString(h[:])}
		}(i, l)
	}
	wg.Wait()

	results := make([]string, len(lines), len(lines))

    for i := 0; i < len(lines); i++ {
		r := <-c
		results[r.order] = r.value
	}

	for _,s := range results {
		fmt.Println(s)
	}
}
