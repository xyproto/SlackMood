package emojimood

import (
	"strconv"
	"sync"
	"time"

	api "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// Fetch each channel in sequence and get the messages
func updateChannels(s *Slack, ts time.Time) {
	channels := channelList
	log.WithFields(log.Fields{
		"channels": len(channels),
	}).Debug("Fetching channel history")

	var wg sync.WaitGroup
	i := 0
	for _, c := range channels {
		if c.IsArchived {
			continue
		}
		i++
		wg.Add(1)

		go getChannelHistory(s, c, &ts, &wg)
		if i%10 == 0 {
			wg.Wait()
		}
	}
}

func getChannelHistory(s *Slack, c api.Channel, ts *time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	hp := api.NewHistoryParameters()
	hp.Oldest = strconv.FormatInt(ts.Unix(), 10)
	hp.Count = 1000
	h, err := s.API.GetChannelHistory(c.ID, hp)

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"channelId": c.ID,
			"channel":   c,
		}).Warning("Could not fetch channel history")
	} else {
		ParseEmoji(h.Messages)

		log.WithFields(log.Fields{
			"channel":   c.Name,
			"channelId": c.ID,
			"messages":  len(h.Messages),
		}).Debug("Got channel history")
	}
}
