package main

import (
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func LetsEncrypt() {
	if !*letsEncrypt {
		return
	}
	go RedirectHttpToHttps()
	e.AutoTLSManager.Cache = autocert.DirCache(*letsEncryptPath)
	if *debug {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		err := e.StartAutoTLS(":443")
		if err != nil {
			log.Fatal(err)
		}
	}
}
