package emojimood

// UserEmojis represents emoji usage per user
type UserEmojis struct {
	User          string
	Positive      float32
	Negative      float32
	PositiveCount int32
	NegativeCount int32
}
