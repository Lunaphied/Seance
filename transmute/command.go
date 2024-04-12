// Handle command logic for Transmute.
package main

import (
	"regexp"
	"strings"
	"fmt"
	"math/rand"
	"bytes"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// TODO: UserFromMember (creates a Mattermost user profile object from a member object).

// Top-level command execution hook called for the commands we register.
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO: Clean up this insanity.
	spaceRegExp := regexp.MustCompile(`\s+`)
	trimmedArgs := spaceRegExp.ReplaceAllString(strings.TrimSpace(args.Command), " ")
	stringArgs := strings.Split(trimmedArgs, " ")
	lengthOfArgs := len(stringArgs)
	restOfArgs := []string{}

	_ = restOfArgs
	_ = lengthOfArgs

	memberName := stringArgs[1]
	//var handler func([]string, *model.CommandArgs) (bool, error)
/*
	if lengthOfArgs == 1 {
		handler = p.runListCommand
	} else {
		command := stringArgs[1]
		if lengthOfArgs > 2 {
			restOfArgs = stringArgs[2:]
		}
		switch command {
		case "new":
			handler = p.runNewCommand
		case "list":
			handler = p.runListCommand
		case "pop":
			handler = p.runPopCommand
		case "send":
			handler = p.runSendCommand
		case "settings":
			handler = p.runSettingsCommand
		default:
			if command == "help" {
				p.trackCommand(args.UserId, command)
			} else {
				p.trackCommand(args.UserId, "not found")
			}
			p.postCommandResponse(args, getHelp())
			return &model.CommandResponse{}, nil
		}
	}

	// Handle the command...
	isUserError, err := handler(restOfArgs, args)
*/

	originUser, _ := p.API.GetUser(args.UserId)

	passwordBytes := make([]byte, 32)
	rand.Read(passwordBytes)
	password := string(bytes.Runes(passwordBytes))

	/* Test of user creation again. */
	/* User creation requires an email and a password. We always set these users to verified as well. */
	// Usernames along with the ID of the user can be automatically generated by Mattermost.
	user := &model.User{
		// Email must be unique per-system * members.
		Email:         fmt.Sprintf("%s@%s.transmute.seance", memberName, args.UserId),
		// TODO: It'd be ideal if there were a more reasonable way to do this,...
		Username: fmt.Sprintf("%s-of-%s", memberName, originUser.Username),
		//FirstName:     "Meow",
		Nickname:   memberName,
		//LastName:      "Proxy",
		Password:      password,
		//Password: "\x00",
		EmailVerified: true,
	}

	_, err := p.API.CreateUser(user)
	if err != nil {
		p.API.LogError("Attempt to create Séance::Transmute user failed", err)
		return nil, &model.AppError{ Message: "Something went wrong" }
	}

	//response := model.CommandResponseFromPlainText(fmt.Sprintf("meow meow %v '%s'", err, password))
	//response.ResponseType = model.CommandResponseTypeInChannel;
	response := &model.CommandResponse{
		ResponseType: model.CommandResponseTypeEphemeral,
		Text: "Hello there!",
	}

	return response, nil
}
