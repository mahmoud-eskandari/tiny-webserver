package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
)

// Redirect SSL
func RedirectHttpToHttps() {
	if !*forceSSL {
		return
	}
	r := echo.New()
	if *sslPort == 443 {
		r.Any("*", func(c echo.Context) error {
			r := c.Request()
			return c.Redirect(301, "https://"+r.Host+r.URL.Path)
		})
	} else {
		r.Any("*", func(c echo.Context) error {
			r := c.Request()
			return c.Redirect(301, fmt.Sprintf("https://%s:%d%s", r.Host, *sslPort, r.URL.Path))
		})
	}

	if *debug {
		r.Logger.Fatal(r.Start(fmt.Sprintf("%s:%d", *host, *httpPort)))
	} else {
		err := r.Start(fmt.Sprintf("%s:%d", *host, *httpPort))
		if err != nil {
			log.Fatal(err)
		}
	}
}
