package main

import (
	"log"
	"net/http"

	ih "caption-backend/http"
	"caption-backend/http/handler"
	"caption-backend/internal/ai"
	"caption-backend/internal/config"
	"caption-backend/internal/payment"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, fallback to OS env")
	}

	// Load config
	cfg := config.Load()

	// Init payment store & verifier
	store := payment.NewStore()
	verifier, err := payment.NewVerifier(
		cfg.BaseRPC,
		cfg.BackendAddress,
		cfg.PriceWei,
		store,
	)
	if err != nil {
		log.Fatal("payment verifier error:", err)
	}

	// Init Gemini AI client (SDK resmi)
	aiClient, err := ai.New(cfg.GeminiAPIKey)
	if err != nil {
		log.Fatal("gemini client error:", err)
	}
	txt, err := aiClient.Generate("siapa xiu jing ping")
	log.Println("TEST GEMINI:", txt, err)

	// HTTP handler
	h := &handler.GenerateHandler{
		Verifier: verifier,
		AI:       aiClient,
	}

	// Router + x402 middleware
	router := ih.Router(h, cfg.BackendAddress, cfg.PriceWei)

	log.Println("API running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
