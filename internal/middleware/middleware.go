package middleware

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.HandlerFunc, ms ...middleware) http.HandlerFunc {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}
	return h
}
