package main

import (
	"log"
	"net/http"

	mandelbrot "github.com/Floriszenz/mandelbrot-generator/go/lib"
)

func generateMandelbrot(w http.ResponseWriter, req *http.Request) {
	c := mandelbrot.MandelbrotConfig{
		ImageWidth:    800,
		ImageHeight:   600,
		MaxIterations: 10_000,
		ThreadCount:   1000,
		RealCenter:    -0.5,
		ImagCenter:    0,
		RealWidth:     2.5,
	}

	if err := mandelbrot.GenerateMandelbrot(w, c); err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/generateMandelbrot", generateMandelbrot)

	http.ListenAndServe(":42069", nil)
}
