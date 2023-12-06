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

    c := MandelbrotConfig {
        imageWidth: 800,
        imageHeight: 600,
        maxIterations: 1_000,
        threadCount: 1,
        realCenter: -0.5,
        imagCenter: 0,
        realWidth: 2.5,
    }

    if err := GenerateMandelbrot(f, c); err != nil {
        f.Close()
        log.Fatal(err)
    }

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

