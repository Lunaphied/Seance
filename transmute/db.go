package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Collective struct {
	gorm.Model

	// Mattermost user ID of base Mattermost account.
	UserID string

	// Members that are a part of this collective.
	Members []Member

	// TODO: Autoproxy.
	// TODO: What other data might be useful to keep on a collective level?
}

type Member struct {
	gorm.Model

	// User ID for the Mattermost user that is used to proxy this member's messages.
	UserID string

	// Regex representing this member's proxy pattern to match in incoming messages.
	Regex string

	// Parent collective this member is part of.
	CollectiveID uint

	// Pronouns for this member.
	Pronouns string

	// TODO: Figure out how to chain a deletion of a Collective to deleting all it's members.
}

// TODO: wrap this in a higher level struct to shield the rest of the logic a bit.
func (p *Plugin) openDB() (*gorm.DB, error) {

	// We make our database inside the plugin's bundle path which is a writable state directory.
	// Get the directory for the bundle.
	plugin_dir, err := p.API.GetBundlePath()
	if err != nil {
		return nil, err
	};

	// Actually open the database.
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/transmute.db", plugin_dir)), &gorm.Config{})

	// Errors are passed upwards.
	if err != nil {
		return nil, err
	}

	// Register all our models automatically.
	db.AutoMigrate(&Collective{})
	db.AutoMigrate(&Member{})

	return db, nil
}
