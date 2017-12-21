package malcolm

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

// Chain is a struct to store middleware
type Chain struct {
	middleware []middleware
}

// Then takes a middleware and returns a chain
func (c *Chain) Then(m middleware) *Chain {

	c.middleware = append([]middleware{m}, c.middleware...)

	return c
}

// Create returns a handler
func (c *Chain) Create() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return c.finally(h)
	}
}

// Then takes a middleware and returns a chain
func Then(m middleware) *Chain {
	return &Chain{[]middleware{m}}
}

func (c *Chain) finally(h http.Handler) http.Handler {

	for _, handler := range c.middleware {
		h = handler(h)
	}

	return h
}
