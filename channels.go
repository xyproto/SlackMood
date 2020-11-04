package emojimood

import (
	api "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

var channelList []api.Channel

// Get a list of all the channels
func channels(s *Slack) {
	log.Debug("Fetching channels")

	chn, err := s.API.GetChannels(false)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warning("Could not fetch channels from Slack")
	} else {
		// TODO: Why not "channelList = chn" ?
		var allChannels []api.Channel
		allChannels = append(allChannels, chn...)
		channelList = allChannels
	}

	log.WithFields(log.Fields{
		"channels": len(channelList),
	}).Debug("Updated channel list")
}
