package main

import (
	"fmt"
	"github.com/labstack/echo/middleware"
	"net/url"
	"strings"
)

func reverseProxy() {
	urls := strings.Split(*reverseProxies, ",")
	fmt.Printf("%+v", urls)
	var targets []*middleware.ProxyTarget
	for _, u := range urls {
		urlObj, err := url.Parse(u)
		if err == nil {
			e.Logger.Infof("(%s) added to balancer \n", u)
			targets = append(targets, &middleware.ProxyTarget{URL: urlObj})
		}
	}

	// Append urls
	if !*randomReverseProxy {
		e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	} else {
		e.Use(middleware.Proxy(middleware.NewRandomBalancer(targets)))
	}
}
