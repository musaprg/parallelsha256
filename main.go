package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"crypto/sha256"
)

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

	r := bufio.NewReader(f)

	for {
		if l, err := r.ReadString('\n'); err == io.EOF {
            break
		} else if err != nil {
			panic(err)
        } else {
			fmt.Print(l)
			//h := sha256.New()
			//h.Write([]byte("hello, world\n"))
			//fmt.Printf("%x", h.Sum(nil))
		}
	}

}
