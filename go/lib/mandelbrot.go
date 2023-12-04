package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
)

const THREAD_COUNT int = 1
const ESCAPE_RADIUS float64 = 2.0
const MAX_ITERATIONS = 1_000
const IMAGE_WIDTH, IMAGE_HEIGHT = 800, 600

// const CX, CY, WIDTH float64 = -0.5, 0.0, 2.5
const CX, CY, WIDTH float64 = -0.7436447860, 0.1318252536, 0.0000029336
// const CX, CY, WIDTH float64 = -0.743643135, 0.131825963, 0.000014628
// const CX, CY, WIDTH float64 = -0.743643900055, 0.131825890901, 0.000000049304

const HEIGHT float64 = (WIDTH * float64(IMAGE_HEIGHT)) / float64(IMAGE_WIDTH)
const RE_START float64 = CX - 0.5 * WIDTH
const RE_END float64 = CX + 0.5 * WIDTH
const IM_START float64 = CY - 0.5 * HEIGHT
const IM_END float64 = CY + 0.5 * HEIGHT


func computePixel(x0, y0 float64, maxIterations int) color.RGBA {
    var x, y float64 = 0.0, 0.0
    i := 0

    for {
        if i == maxIterations {
            return color.RGBA{0, 0, 0, 255}
        }

        if x * x + y * y > ESCAPE_RADIUS * ESCAPE_RADIUS {
            break
        }

        x, y = x * x - y * y + x0, 2.0 * x * y + y0
        i++
    }

    h := 360.0 - (360.0 * float64(i) / float64(maxIterations))
    r, g, b := hsvToRgb(h, 0.8, 0.8)

    return color.RGBA{r, g, b, 255}
}

func generateMandelbrotSequentially(w io.Writer) error {
    img := image.NewRGBA(image.Rect(0, 0, IMAGE_WIDTH, IMAGE_HEIGHT))

    for y := 0; y < IMAGE_HEIGHT; y++ {
        y0 := float64(y) * ((IM_END - IM_START) / IMAGE_HEIGHT) + IM_START

        for x := 0; x < IMAGE_WIDTH; x++ {
            x0 := float64(x) * ((RE_END - RE_START) / IMAGE_WIDTH) + RE_START
            c := computePixel(x0, y0, MAX_ITERATIONS)

            img.Set(x, IMAGE_HEIGHT - y, c)
        }
    }

    return png.Encode(w, img)
}

func GenerateMandelbrot(w io.Writer) error {
    if THREAD_COUNT == 1 {
        return generateMandelbrotSequentially(w)
    } else {
        log.Fatal("unimplemented")
    }

    return nil
}


