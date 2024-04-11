// Handle command logic for Transmute.
package main

import (
	//"regexp"
	//"strings"
	//"fmt"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// Top-level command execution hook called for the commands we register.
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO: Clean up this insanity.
	//spaceRegExp := regexp.MustCompile(`\s+`)
	//trimmedArgs := spaceRegExp.ReplaceAllString(strings.TrimSpace(args.Command), " ")
	//stringArgs := strings.Split(trimmedArgs, " ")
	//lengthOfArgs := len(stringArgs)
	//restOfArgs := []string{}

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

	/* Test of user creation again. */
	user := &model.User{
		Email:         "meower_proxy@localhost",
		Username:      "meower_proxy",
		FirstName:     "Meow",
		Nickname:   "Iris",
		//LastName:      "Proxy",
		Password:      "Pa$$word11",
		//Password: "\x00",
		EmailVerified: true,
	}

	_, err := p.API.CreateUser(user)
	if err != nil {
		p.API.LogError("CreateUser error: {}", err)
	}

	//fmt.Fprint(w, ruser.Id+"\n")
	// ... and output the error in a way that reflects the type of error (user input mistake or internal exception).
	isUserError := false
	if err != nil {
		if isUserError {
		} else {
			//p.API.LogError(err.Error())
		}
	}

	return &model.CommandResponse{}, nil
}
