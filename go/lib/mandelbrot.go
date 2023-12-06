package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"sync"
)

type MandelbrotConfig struct {
    imageWidth, imageHeight, maxIterations, threadCount int
    realCenter, imagCenter, realWidth float64
}

const ESCAPE_RADIUS float64 = 2.0

func inferRealBounds(c *MandelbrotConfig) (realStart float64, realEnd float64) {
    realStart = c.realCenter - 0.5 * c.realWidth
    realEnd = c.realCenter + 0.5 * c.realWidth

    return
}

func inferImagBounds(c *MandelbrotConfig) (imagStart float64, imagEnd float64) {
    imagHeight := (c.realWidth * float64(c.imageHeight)) / float64(c.imageWidth)
    imagStart = c.imagCenter - 0.5 * imagHeight
    imagEnd = c.imagCenter + 0.5 * imagHeight

    return
}

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

func generateMandelbrotChunk(img *image.RGBA, c *MandelbrotConfig, rows int, offset int) {
    realStart, realEnd := inferRealBounds(c)
    imagStart, imagEnd := inferImagBounds(c)

    for y := offset; y < offset + rows; y++ {
        y0 := float64(y) * ((imagEnd - imagStart) / float64(img.Rect.Dy())) + imagStart

        for x := 0; x < img.Rect.Dx(); x++ {
            x0 := float64(x) * ((realEnd - realStart) / float64(img.Rect.Dx())) + realStart
            pix := computePixel(x0, y0, c.maxIterations)

            img.Set(x, img.Rect.Dy() - y, pix)
        }
    }
}

func generateMandelbrotSequentially(w io.Writer, c MandelbrotConfig) error {
    img := image.NewRGBA(image.Rect(0, 0, c.imageWidth, c.imageHeight))
    realStart, realEnd := inferRealBounds(&c)
    imagStart, imagEnd := inferImagBounds(&c)

    for y := 0; y < img.Rect.Dy(); y++ {
        y0 := float64(y) * ((imagEnd - imagStart) / float64(img.Rect.Dy())) + imagStart

        for x := 0; x < img.Rect.Dx(); x++ {
            x0 := float64(x) * ((realEnd - realStart) / float64(img.Rect.Dx())) + realStart
            pix := computePixel(x0, y0, c.maxIterations)

            img.Set(x, img.Rect.Dy() - y, pix)
        }
    }

    return png.Encode(w, img)
}

func generateMandelbrotConcurrently(w io.Writer, c MandelbrotConfig) error {
    var wg sync.WaitGroup
    img := image.NewRGBA(image.Rect(0, 0, c.imageWidth, c.imageHeight))
    threadCount := int(math.Min(float64(c.threadCount), float64(img.Rect.Dy())))

    if threadCount < c.threadCount {
        fmt.Printf("warning: number of thread was clamped to %d\n", threadCount)
    }

    chunkSize := img.Rect.Dy() / threadCount

    for i := 0; i < threadCount; i++ {
        wg.Add(1)
        i := i

        go func() {
            defer wg.Done()
            generateMandelbrotChunk(img, &c, chunkSize, i * chunkSize)
        }()
    }

    wg.Wait()

    return png.Encode(w, img)
}

func GenerateMandelbrot(w io.Writer, c MandelbrotConfig) error {
    if c.threadCount == 1 {
        return generateMandelbrotSequentially(w, c)
    } else {
        return generateMandelbrotConcurrently(w, c)
    }
}


