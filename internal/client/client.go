package client

import (
	"os"
	"strconv"

	"github.com/pusher/pusher-http-go/v5"
)

func NewPusher() pusher.Client {
	secure, err := strconv.ParseBool(os.Getenv("PUSHER_SECURE"))
	if err != nil {
		secure = false
	}

	client := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  secure,
	}

	return client
}
