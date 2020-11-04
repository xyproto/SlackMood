package happyteam

import (
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	api "github.com/slack-go/slack"
)

// Fetch each channel in sequence and get the messages
func updateConversations(s *Slack, ts time.Time) {
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

		go getConversationHistory(s, c, &ts, &wg)
		if i%10 == 0 {
			wg.Wait()
		}
	}
}

func getConversationHistory(s *Slack, c api.Channel, ts *time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	params := &api.GetConversationHistoryParameters{}
	params.ChannelID = c.ID
	params.Oldest = strconv.FormatInt(ts.Unix(), 10)
	params.Limit = 1000

	h, err := s.API.GetConversationHistory(params)

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
