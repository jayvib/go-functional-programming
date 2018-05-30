package decorator

import (
	"net/http"
	"time"
	"log"
)

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

type ClientFunc func(r *http.Request) (*http.Response, error)
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

var ratelimitDuration time.Duration

func (f ClientFunc) SetRatelimit(duration time.Duration) error {
	ratelimitDuration = duration
	return nil
}

func (f ClientFunc) GetRateLimit() (time.Duration, error) {
	return ratelimitDuration, nil
}

type Decorator func(Client) Client

func Authorize(token string) Decorator {
	return Header("Authorization", token)
}

func Header(key, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request)(*http.Response, error){
			r.Header.Add(key, value)
			return c.Do(r)
		})
	}
}

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error){
			l.Printf("%s %s", r.Method, r.URL)
			return c.Do(r)
		})
	}
}