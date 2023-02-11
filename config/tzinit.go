package config

import (
	"log"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "Asia/Manila")
	location, err := time.LoadLocation("Asia/Manila")

	if err != nil {
		log.Fatal(err)
	}

	time.Local = location
}
