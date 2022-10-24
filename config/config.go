package config

import (
	"os"
	"strconv"
)

type Config struct {
	DB                string
	Bucket            string
	ShortenerHostname string
	IdAlphabet        string
	IdLength          int
	Port              string
}

func GetConfig() Config {
	dbPath := os.Getenv("TINYURL_DB")

	if dbPath == "" {
		dbPath = "my.db"
	}

	dbBucket := os.Getenv("TINYURL_BUCKET")

	if dbBucket == "" {
		dbBucket = "tinyurl"
	}

	shortenerHostname := os.Getenv("TINYURL_HOSTNAME")

	if shortenerHostname == "" {
		shortenerHostname = "localhost"
	}

	port := os.Getenv("TINYURL_PORT")
	if port == "" {
		port = "8080"
	}

	idLength, err := strconv.ParseInt(os.Getenv("TINYURL_ID_LENGTH"), 10, 32)

	if idLength == 0 || err != nil {
		idLength = 6
	}

	idAlphabet := os.Getenv("TINYURL_ID_ALPHABET")

	if idAlphabet == "" {
		idAlphabet = "0123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ"
	}

	return Config{
		ShortenerHostname: shortenerHostname,
		IdLength:          int(idLength),
		IdAlphabet:        idAlphabet,
		Port:              port,
		DB:                dbPath,
		Bucket:            dbBucket,
	}
}
