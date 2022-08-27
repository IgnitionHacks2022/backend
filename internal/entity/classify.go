package entity

type Classify struct {
	Contents string `json:"contents"`
}

type ClassifyResponse struct {
	Type string `json:"type"`
}

// Response

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
