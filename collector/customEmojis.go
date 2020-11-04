package collector

import (
	log "github.com/sirupsen/logrus"
	"github.com/xyproto/emojimood/slack"
)

// Loads the custom emoji for the team
func customEmoji(s *slack.Slack) {
	log.Debug("Fetching custom emoji")

	emoji, err := s.Api.GetEmoji()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warning("Could not fetch emoji from Slack")
	} else {
		var customEmoji []string
		for emoji := range emoji {
			customEmoji = append(customEmoji, emoji)
		}

		log.WithFields(log.Fields{
			"emojiCount": len(customEmoji),
		}).Debug("Loaded custom emoji")
	}
}
