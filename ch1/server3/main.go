package main

import (
	"fmt"
	"gopl-solutions/ch1/lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hej!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.ParseFloat(r.URL.Query().Get("cycles"), 64)
		if err != nil {
			log.Println("Could not parse cycles param, ", err)
		}
		lissajous.Lissajous(w, cycles)
		fmt.Println(r.URL.Query().Get("cycles"))
	})
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
}
