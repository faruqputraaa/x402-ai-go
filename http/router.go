package http

import (
	"net/http"

	"caption-backend/http/middleware"
)

func Router(handler http.Handler, to string, price int64) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/generate", middleware.X402(middleware.X402Config{
		PriceWei: price,
		To:       to,
	})(handler))
	return mux
}
