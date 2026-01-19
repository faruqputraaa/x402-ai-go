package domain

type CaptionRequest struct {
	Topic    string `json:"topic"`
	Tone     string `json:"tone"`
	Audience string `json:"audience"`
}

type CaptionResponse struct {
	Caption string `json:"caption"`
}
