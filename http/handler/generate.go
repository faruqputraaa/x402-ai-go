package handler

import (
	"encoding/json"
	"net/http"

	"caption-backend/internal/ai"
	"caption-backend/internal/domain"
	"caption-backend/internal/payment"
)

type GenerateHandler struct {
	Verifier *payment.Verifier
	AI       *ai.Client
}

func (h *GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tx := r.Header.Get("x402-payment")
	if err := h.Verifier.Verify(tx); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var req domain.CaptionRequest
	json.NewDecoder(r.Body).Decode(&req)

	prompt := ai.Prompt(req.Topic, req.Tone, req.Audience)
	caption, err := h.AI.Generate(prompt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(domain.CaptionResponse{Caption: caption})
}
