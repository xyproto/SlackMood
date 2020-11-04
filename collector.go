package happyteam

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// StartEmojiCollector connects to Slack and
// starts streaming all public channels
func (config *Config) StartEmojiCollector() bool {
	s, err := Connect(config)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Could not connect to Slack!")
		return false
	}

	customEmoji(s)
	go collector(s)

	return true
}

func collector(s *Slack) {
	oldest := time.Now().Add(-30 * 24 * time.Hour)
	channels(s)
	updateChannels(s, oldest)
	oldest = time.Now()

	ticker := time.NewTicker(15 * time.Minute)
	for now := range ticker.C {
		log.WithFields(log.Fields{
			"now": now,
		}).Info("Running collector")

		channels(s)
		updateChannels(s, oldest)
		oldest = time.Now()
	}

	ticker.Stop()
}
