package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    h := sha256.New()
    h.Write([]byte("hello, world\n"))
    fmt.Printf("%x", h.Sum(nil))
}
