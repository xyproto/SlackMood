package emojigo

import (
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func OpenDB() error {
	var err error
	db, err = bolt.Open(Config.Db, 0600, nil)
	if err != nil {
		return err
	}

	// IGNORE ALL THE ERRORS \o/
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("emoji"))
		return nil
	})

	return nil
}
