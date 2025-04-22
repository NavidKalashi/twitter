package events

type Feed struct {
	Text      string   `json:"text"`
	MediaType string   `json:"media_type"`
	FileNames []string `json:"file_name"`
	Username  string   `json:"username"`
}