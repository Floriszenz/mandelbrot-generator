package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	mandelbrot "github.com/Floriszenz/mandelbrot-generator/go/lib"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleOptionsRequest(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
}

func generateMandelbrot(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	rawBody, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var body mandelbrot.MandelbrotConfig

	if err := json.Unmarshal(rawBody, &body); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if err := mandelbrot.GenerateMandelbrot(w, body); err != nil {
		http.Error(w, "Error while generating the mandelbrot image", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

}

func main() {
	http.HandleFunc("POST /generateMandelbrot", generateMandelbrot)
	http.HandleFunc("OPTIONS /generateMandelbrot", handleOptionsRequest)

	http.ListenAndServe(":42069", nil)
}
