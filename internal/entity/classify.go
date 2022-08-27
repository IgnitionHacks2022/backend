package entity

type Classify struct {
	Contents string `json:"contents"`
}

type ClassifyResponse struct {
	Type  string `json:"type"`
	Audio string `json:"audio"`
}

// GOOGLE VISION API
// Request

type ImageContent struct {
	Content string `json:"content"`
}

type FeatureContent struct {
	Type       string `json:"type"`
	MaxResults int    `json:"maxResults"`
}

type ImageContextContent struct {
	Content string `json:"content"`
}

type ImageType struct {
	Image    ImageContent     `json:"image"`
	Features []FeatureContent `json:"features"`
	// ImageContext ImageContextContent `json:"imageContext"`
}

type GoogleRequest struct {
	Requests []ImageType `json:"requests"`
}

// Response

type LabelAnnotation struct {
	Description string `json:"description"`
}

type LocalizedObjectAnnotation struct {
	Name string `json:"name"`
}

type ResponseContent struct {
	LabelAnnotations           []LabelAnnotation           `json:"labelAnnotations"`
	LocalizedObjectAnnotations []LocalizedObjectAnnotation `json:"localizedObjectAnnotations"`
}

type GoogleResponse struct {
	Responses []ResponseContent `json:"responses"`
}

// GOOGLE Text to Speech API
// Request

type AudioConfig struct {
	AudioEncoding string  `json:"audioEncoding"`
	Pitch         float32 `json:"pitch"`
	SpeakingRate  float32 `json:"speakingRate"`
}

type Voice struct {
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type Input struct {
	Text string `json:"text"`
}

type SpeechRequest struct {
	AudioConfig AudioConfig `json:"audioConfig"`
	Input       Input       `json:"input"`
	Voice       Voice       `json:"voice"`
}

// Response

type SpeechResponse struct {
	AudioContent string `json:"audioContent"`
}
