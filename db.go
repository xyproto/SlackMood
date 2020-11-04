package happyteam

import (
	"github.com/boltdb/bolt"

	log "github.com/sirupsen/logrus"
)

var db *bolt.DB

// OpenDB tries to open a Bolt DB file
func (config *Config) OpenDB() error {
	var err error
	db, err = bolt.Open(config.DBPath, 0600, nil)
	if err != nil {
		return err
	}

	// IGNORE ALL THE ERRORS \o/
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("emoji"))
		return nil
	})

	log.WithFields(log.Fields{
		"filename": config.DBPath,
	}).Info("Loaded database")

	return nil
}
