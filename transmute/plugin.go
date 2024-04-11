package main

import (
	"sync"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// Initialization during plugin activation.
func (p *Plugin) OnActivate() error {
	// For now all we do here is register our base command. In the future we'll open and store a reference to
	// the database.
	p.openDB()

	// TODO: Maybe use a function that returns this or a const? The TODO plugin generates this with a function.
	return p.API.RegisterCommand(&model.Command{
		Trigger: "transmute",
		DisplayName: "Transmute",
		Description: "Invoke a ritual to channel the unseen",
	})
}

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	// _user, err := &p.API.GetUser("8grsi1tzfibj5c7zjwucng47ur");
	// if err != nil {
	// 	p.API.LogError("GetUser error: {}", err)
	// }

	post.UserId = "8grsi1tzfibj5c7zjwucng47ur"

	return post, ""
}

func (p *Plugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	if oldPost.UserId != "meowmeow" {
		return nil, "failed to edit, not authorized"
	}
	return nil, "meow"
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
