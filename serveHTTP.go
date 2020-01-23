package main

import (
	"fmt"
	"log"
)

func serveHTTP() {
	if *ssl {
		return
	}
	if *debug {
		e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", *host, *httpPort)))
	} else {
		err := e.Start(fmt.Sprintf("%s:%d", *host, *httpPort))
		if err != nil {
			log.Fatal(err)
		}
	}
}
