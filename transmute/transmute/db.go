package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Collective struct {
	gorm.Model

	// System name.
	Name string

	// Mattermost user ID of base Mattermost account.
	UserID string

	// Members that are a part of this collective.
	Members []Member

	// TODO: Autoproxy.
	// TODO: What other data might be useful to keep on a collective level?
}

type Member struct {
	gorm.Model

	// Member name. (Convenience for commands, default display name).
	Name string

	// User ID for the Mattermost user that is used to proxy this member's messages.
	UserID string

	// Member display name (Nickname in Mattermost). May be empty (null).
	DisplayName *string

	// Regex representing this member's proxy pattern to match in incoming messages.
	Regex string

	// Parent collective this member is part of.
	CollectiveID uint

	// Pronouns for this member.
	Pronouns string

	// TODO: Is this required to be stored? I think we can just keep track of it on the Mattermost side.
	// Path to where this Member's avatar is stored (this is only used on upload).
	// AvatarURL string

	// TODO: Figure out how to chain a deletion of a Collective to deleting all it's members.
}

// TODO: wrap this in a higher level struct to shield the rest of the logic a bit.
func openDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("transmute.db"), &gorm.Config{})

	// Errors are passed upwards.
	if err != nil {
		return nil, err
	}

	// Register all our models automatically.
	db.AutoMigrate(&Collective{})
	db.AutoMigrate(&Member{})

	return db, nil
}
