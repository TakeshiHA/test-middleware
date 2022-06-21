package models

type Joke struct {
	ID    string `json:"id,omitempty"`
	URL   string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
}
