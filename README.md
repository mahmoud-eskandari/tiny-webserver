# Tiny Webserver
A golang echo tiny webserver/loadbalancer with autocert.

## Features
* Serve static files
* Round-robin reverse proxy
* Random reverse proxy
* Autocert any domain
* Max size / rate limit protect

### Installation:
- Download latest build files from github release page
- Add binary file or it's link to your `PATH` environment
- run `chmod +x tiny-webserver` in unix/linux to make binary file executable for first time

### Usage:
```shell script
  -cert_crt_path string
        SSL CRT file path (default "/var/www/ssl/cert.crt")
  -cert_key_path string
        SSL key file path (default "/var/www/ssl/cert.key")
  -debug
        Debug mode and log requests
  -force_ssl
        Force clients to using SSL, SSL redirect (default true)
  -http_port int
        HTTP port (default 80)
  -ip string
        ip bind host (default 0.0.0.0)
  -lets
        Let's Encrypt auto cert
  -lets_path string
        Let's Encrypt data cache (default "/var/www/.cache")
  -max_body string
        Max POST body size (default "1M")
  -p string
        reverse proxies comma delimited(http://host1,http://host2...)
  -path string
        file server public path
  -ssl
        SSL enabled
  -ssl_port int
        SSL default port (default 443)
  -rb int
        Rate Limit Burst (default 3)
  -rt int
        Rate Limit time(Minute) (default 3)
```

### It's made for ease of use
#### Just for fun run