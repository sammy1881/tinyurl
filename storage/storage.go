package storage

import (
	"fmt"

	bolt "github.com/boltdb/bolt"
	c "github.com/sammy1881/tinyurl/config"
)

// Storage fuctions Interface
type StorageService interface {
	Put(id string, url string) error
	Get(id string) (string, error)
	Count() (int64, error)
	GetAllRecords() (string, error)
}

// BoltTransport implements the TransportInterface using the Bolt database.
type BoltStorage struct {
	db         *bolt.DB
	bucketName string
}

func NewStrorageService() (StorageService, error) {
	return NewBoltStorage()
}

// NewBoltTransport create a new BoltStorage instange.
func NewBoltStorage() (*BoltStorage, error) {

	config := c.GetConfig()
	var err error

	bucketName := config.Bucket

	path := config.DB

	u := fmt.Sprintf(path + "/" + bucketName)

	if path == "" {
		return nil, fmt.Errorf("DB Path missing: %s", u)
	}

	db, err := bolt.Open(path, 0o600, nil)
	if err != nil {
		return nil, fmt.Errorf(`%s: %s`, u, err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf(`%s: bucket could not be created: %w`, u, err)
	}

	return &BoltStorage{
		db:         db,
		bucketName: bucketName,
	}, nil
}

func (b *BoltStorage) Put(id string, url string) error {
	defer b.db.Close()
	return b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		err := b.Put([]byte(id), []byte(url))
		return err
	})
}

func (b *BoltStorage) Get(id string) (string, error) {
	var url string
	defer b.db.Close()
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		url = string(b.Get([]byte(id)))
		return nil
	})

	return url, err
}

func (b *BoltStorage) GetAllRecords() (resp string, err error) {
	defer b.db.Close()
	err = b.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(b.bucketName))
		cu := b.Cursor()
		resp = "<div class=\"col-md-12\"><table><tr><td>ShortLink</td><td>Link</td></tr>"

		for k, v := cu.First(); k != nil; k, v = cu.Next() {
			key := string(k)
			value := string(v)
			resp += fmt.Sprintf("<tr><td>"+c.GetConfig().ShortenerHostname+":"+c.GetConfig().Port+"/%s</td><td><a href=\"%s\">%s</a></td></tr>", key, value, value)
		}
		resp += "</table></div>"
		return nil
	})
	return resp, err
}

func (b *BoltStorage) Count() (int64, error) {
	var count int64
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		count = int64(b.Stats().KeyN)
		return nil
	})
	return count, err
}
