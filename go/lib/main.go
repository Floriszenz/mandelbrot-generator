package mandelbrot

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
        ImageWidth: 800,
        ImageHeight: 600,
        MaxIterations: 10_000,
        ThreadCount: 1000,
        RealCenter: -0.5,
        ImagCenter: 0,
        RealWidth: 2.5,
    }

    if err := GenerateMandelbrot(f, c); err != nil {
        f.Close()
        log.Fatal(err)
    }

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

