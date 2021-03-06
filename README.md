# Tiny Webserver (tserver)
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

### download and installation on linux

```shell script
sudo curl -L https://github.com/mahmoud-eskandari/tiny-webserver/releases/download/1.0.0/linux-amd64 -o /usr/local/bin/tserver
sudo chmod +x /usr/local/bin/tserver
sudo ln -s /usr/local/bin/tserver /usr/bin/tserver
tserver -v
```
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
  -random-proxy
    	  set reverse proxy algorithm to random (default: round-robin)
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
  -v int
       Print current version and exit
```

### It's made for ease of use
#### Just for fun run
