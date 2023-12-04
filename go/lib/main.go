package main

import (
	"log"
	"os"
)

func main() {
    f, err := os.Create("mandelbrot.png")
    if err != nil {
        log.Fatal(err)
    }

    if err := GenerateMandelbrot(f); err != nil {
        f.Close()
        log.Fatal(err)
    }

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

