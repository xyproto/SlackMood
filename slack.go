package happyteam

import (
	"fmt"

	api "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// Slack represents a connection to a Slack server
type Slack struct {
	API *api.Client
}

// Connect tries to connect to Slack
func Connect(config *Config) (*Slack, error) {
	s := Slack{}

	s.API = api.New(config.SlackToken)
	auth, err := s.API.AuthTest()
	if err != nil {
		return &s, fmt.Errorf("error authenticating with Slack: %s", err)
	}

	log.WithFields(log.Fields{
		"teamName": auth.Team,
		"userId":   auth.UserID,
		"teamUrl":  auth.URL,
	}).Info("Authenticated with Slack")

	return &s, nil
}
