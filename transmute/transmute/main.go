package main

import (
	"github.com/mattermost/mattermost/server/public/plugin"
)

func main() {
	_, err := openDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	plugin.ClientMain(&Plugin{})
}
