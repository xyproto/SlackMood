package emojimood

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Rank struct {
	Name string
	Rank int64
}

var EmojiRanks []Rank

func LoadRankings(filename string) error {
	fp := path.Join("rankings", filename)
	fc, err := os.Open(fp)
	if err != nil {
		log.WithFields(log.Fields{
			"file":  fp,
			"error": err,
		}).Error("Could not load emoji rank file")
		return errors.New("Could not loadd emoji rank file: " + filename)
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
		EmojiRanks = append(EmojiRanks, rank)
	}

	log.WithFields(log.Fields{
		"emoji": len(EmojiRanks),
	}).Info("Loaded emoji rankings")

	return nil
}
