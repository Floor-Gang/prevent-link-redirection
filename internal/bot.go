package internal

import (
	"log"

	auth "github.com/Floor-Gang/authclient"
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
)

// Bot structure
type Bot struct {
	Auth   *auth.AuthClient
	Client *dg.Session
	Config *Config
}

// Start starts discord client, configuration and database
func Start() {
	var err error

	// get Config.yml
	config := GetConfig()

	// setup authentication server
	// you can use this to get the bot's access token
	// and authenticate each user using a command.
	authClient, err := auth.GetClient(config.Auth)

	if err != nil {
		log.Fatalln("Failed to connect to authentication server", err)
	}

	register, err := authClient.Register(
		auth.Feature{
			Name:        "Stop Redirect",
			Description: "Looks for links to check if they are redirecting to another web site.",
			Commands: []auth.SubCommand{
				{
					Name:        "help",
					Description: "This command describes what the bot does.",
					Example:     []string{""},
				},
			},
			CommandPrefix: config.Prefix,
		},
	)

	if err != nil {
		log.Fatalln("Failed to register with authentication server", err)
	}

	client, err := dg.New(register.Token)

	if err != nil {
		panic(err)
	}

	bot := Bot{
		Auth:   &authClient,
		Client: client,
		Config: &config,
	}

	client.AddHandlerOnce(bot.onReady) // This will call onReady only once
	client.AddHandler(bot.onMessage)   // This will catch all the new messages that the bot can see

	if err = client.Open(); err != nil {
		util.Report("Was an authentication token provided?", err)
	}
}
