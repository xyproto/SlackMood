package happyteam

import (
	log "github.com/sirupsen/logrus"
	api "github.com/slack-go/slack"
)

var channelList []api.Channel

// Get a list of all the channels
func channels(s *Slack) {
	log.Debug("Fetching channels")

	params := &api.GetConversationsParameters{}
	channels, _, err := s.API.GetConversations(params)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warning("Could not fetch channels from Slack")
	}

	log.WithFields(log.Fields{
		"channels": len(channels),
	}).Debug("Updated channel list")

	channelList = channels
}
