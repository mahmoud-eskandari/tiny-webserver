package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// All configuration flags
var (
	debug           = flag.Bool("debug", false, "Debug mode and log requests")
	letsEncrypt     = flag.Bool("lets", false, "Let's Encrypt auto cert")
	letsEncryptPath = flag.String("lets_path", "/var/www/.cache", "Let's Encrypt data cache")

	ssl         = flag.Bool("ssl", false, "SSL enabled")
	certKeyPath = flag.String("cert_key_path", "/var/www/ssl/cert.key", "SSL key file path")
	certCrtPath = flag.String("cert_crt_path", "/var/www/ssl/cert.crt", "SSL CRT file path")
	forceSSL    = flag.Bool("force_ssl", true, "Force clients to using SSL, SSL redirect")

	sslPort  = flag.Int("ssl_port", 443, "SSL default port")
	httpPort = flag.Int("http_port", 80, "HTTP port")

	fileServerPath     = flag.String("path", "", "file server public path")
	reverseProxies     = flag.String("p", "", "reverse proxies comma delimited(http://host1,http://host2...)")
	randomReverseProxy = flag.Bool("random-proxy", false, "set reverse proxy algorithm to random (default: round-robin)")

	maxBodySize    = flag.String("max_body", "1M", "Max POST body size")
	rateLimitBurst = flag.Int("rb", 3, "Rate Limit Burst")
	rateLimitTime  = flag.Int64("rt", 3, "Rate Limit time(Minute)")

	host = flag.String("ip", "", "ip bind host (default 0.0.0.0)")
	// Print Version
	version = flag.Bool("v", false, "Print tiny-webserver version and exit")
)

var e *echo.Echo

func main() {
	//Parse Items
	flag.Parse()
	if *version {
		fmt.Println("Tiny Webserver version: " + VERSION)
		fmt.Println("Golang Echo version: " + echo.Version)
		return
	}

	//Create New Echo handler
	e = echo.New()

	// Set Middleware
	if *debug {
		e.Use(middleware.Logger())
		e.Debug = true
		fmt.Println("Debug mode enabled!")
	}

	// Recovery middleware
	e.Use(middleware.Recover())

	//Add Protection
	if *maxBodySize != "" {
		e.Use(middleware.BodyLimit(*maxBodySize))
	}
	if *rateLimitBurst > 0 {
		go RemoveOldVisitorsHistory()
		e.Use(RateLimit())
	}

	//Add proxies
	if *reverseProxies != "" {
		reverseProxy()
	} else if *fileServerPath != "" {
		// Add Static files
		e.Static("/", *fileServerPath)
	}

	// Serve Auto Cert
	LetsEncrypt()
	// Serve SSL
	serveSSL()
	// Serve http
	serveHTTP()
}
