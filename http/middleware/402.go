package middleware

import (
	"net/http"
	"strconv"
)

type X402Config struct {
	PriceWei int64
	To       string
}

func X402(cfg X402Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("x402-payment") == "" {
				w.Header().Set("x402-price", strconv.FormatInt(cfg.PriceWei, 10))
				w.Header().Set("x402-to", cfg.To)
				w.Header().Set("x402-chain", "base")
				w.WriteHeader(http.StatusPaymentRequired)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
