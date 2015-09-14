package main

import (
	"flag"
	"log"
	"net/http"
)

var address = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	log.Println("Start Connection Hub routine")
	go h.run()
	log.Println("Add Web socket handler")
	http.HandleFunc("/", serveWs)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
