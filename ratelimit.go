package main

import (
	"github.com/labstack/echo"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, *rateLimitBurst)
		// Include the current time when creating a new visitor.
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

func RemoveOldVisitorsHistory() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		defer mu.Unlock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > time.Duration(*rateLimitTime)*time.Minute {
				delete(visitors, ip)
			}
		}
	}
}

func RateLimit() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			limiter := getVisitor(c.RealIP())
			if !limiter.Allow() {
				return c.NoContent(http.StatusTooManyRequests)
			}
			return handlerFunc(c)
		}
	}
}
