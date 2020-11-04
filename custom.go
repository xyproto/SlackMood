package emojimood

import (
	log "github.com/sirupsen/logrus"
)

// Loads the custom emoji for the team
func customEmoji(s *Slack) {
	log.Debug("Fetching custom emoji")

	emoji, err := s.API.GetEmoji()
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
