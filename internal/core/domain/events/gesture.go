package events

type Gesture struct {
	TweetID     string `json:"tweet_id"`
	GestureType string `json:"gesture_type"`
	Username  string   `json:"username"`
}