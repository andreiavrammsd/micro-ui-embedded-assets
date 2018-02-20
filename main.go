package main

import (
	"flag"
	"log"
	"net/http"
)

const defaultAddr = ":1234"
const maxLimit = 100
const defaultLimit = 50
const timeFormat = "2006-01-02 15:04:05"

var endpoints map[string]string

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	endpoints = map[string]string{
		"users": "https://users.api.company.tld/",
	}

	addr := flag.String("addr", defaultAddr, "-addr host:port")
	flag.Parse()
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.HandleFunc("/template/static/", StaticHandler)
	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
