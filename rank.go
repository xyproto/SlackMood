package happyteam

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Rank represents an emoji together with a rank number
type Rank struct {
	Name string
	Rank int64
}

// EmojiRanks is a collection of Rank structs
type EmojiRanks []Rank

// LoadRanks loads the rank CSV file
func (config *Config) LoadRanks() (*EmojiRanks, error) {
	var emojiRanks EmojiRanks
	fc, err := os.Open(config.RankFile)
	if err != nil {
		log.WithFields(log.Fields{
			"file":  config.RankFile,
			"error": err,
		}).Error("Could not load emoji rank file")
		return nil, errors.New("Could not loadd emoji rank file: " + config.RankFile)
	}

	r := csv.NewReader(bufio.NewReader(fc))
	for {
		rc, err := r.Read()
		if err == io.EOF {
			break
		}

		rank := Rank{}
		rank.Name = strings.TrimLeft(strings.TrimRight(rc[0], ":"), ":")
		rank.Rank, _ = strconv.ParseInt(rc[3], 10, 32)
		emojiRanks = append(emojiRanks, rank)
	}

	log.WithFields(log.Fields{
		"emoji": len(emojiRanks),
	}).Info("Loaded emoji rankings")

	return &emojiRanks, nil
}
