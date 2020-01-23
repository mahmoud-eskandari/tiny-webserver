package main

import (
	"fmt"
	"log"
)

func serveSSL() {
	if !*ssl {
		return
	}
	go RedirectHttpToHttps()
	// Listen | Set Logger
	if *debug {
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf("%s:%d", *host, *sslPort), *certCrtPath, *certKeyPath))
	} else {
		err := e.StartTLS(fmt.Sprintf("%s:%d", *host, *sslPort), *certCrtPath, *certKeyPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
