package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var hostname string
var help bool

func init() {
	flag.BoolVar(&help, "h", false, "show usage")
	flag.StringVar(&hostname, "s", "http://localhost:12345", "website to fetch, eg: https://google.com")
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	}

	res, err := http.Get(hostname)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var buf bytes.Buffer

	io.Copy(&buf, res.Body)
	pageStats(buf.Bytes())
}

func pageStats(buf []byte) {

	fmt.Printf("Stats for Page from '%v' has:\n", hostname)
	fmt.Printf("\tPage size: %v bytes\n", len(buf))
	fmt.Printf("\tAmount of <button>: %v\n", bytes.Count(buf, []byte("<button")))
	fmt.Printf("\tAmount of <div>: %v\n", bytes.Count(buf, []byte("<div")))
	fmt.Printf("\tAmount of <form>: %v\n", bytes.Count(buf, []byte("<form")))
	fmt.Printf("\tAmount of <img>: %v\n", bytes.Count(buf, []byte("<img")))
	fmt.Printf("\tAmount of <input>: %v\n", bytes.Count(buf, []byte("<input")))
	fmt.Printf("\tAmount of <a>: %v\n", bytes.Count(buf, []byte("<a")))
	fmt.Printf("\tAmount of <p>: %v\n", bytes.Count(buf, []byte("<p")))
}
