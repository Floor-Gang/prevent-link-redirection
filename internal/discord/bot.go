package discord

import (
	"github.com/Floor-Gang/prevent-link-redirection/internal"
	"log"

	auth "github.com/Floor-Gang/authclient"
	"github.com/bwmarrin/discordgo"
)

// Bot structure
type Bot struct {
	config  Config
	client  *discordgo.Session
	confLoc string
	auth    auth.AuthClient
}

// Start starts discord client, configuration and database
func Start(config Config, configLocation string) {

	// Setup Authentication client
	authClient, err := auth.GetClient(config.Auth)

	if err != nil {
		log.Fatalln("Failed to connect to authentication server because \n" + err.Error())
	}

	register, err := authClient.Register(
		auth.Feature{
			Name:          "PrevetLinkRedirection",
			Description:   "This is responsible preventing links to include redirection",
			Commands:      []auth.SubCommand{}, //TODO add commands if necessary
			CommandPrefix: config.Prefix,
		},
	)

	if err != nil {
		log.Fatalln("Failed to register with the authenticaiton server\n" + err.Error())
	}

	// Setup Discord
	client, _ := discordgo.New(register.Token)

	intents := discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	client.Identify.Intents = intents

	bot := Bot{
		config:  botConfig,
		client:  client,
		confLoc: configPath,
		auth:    authClient,
	}

	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onMessage)

	if err = client.Open(); err != nil {
		log.Fatalln(
			"Failed to connect to Discord, was an access token provided?\n" +
				err.Error(),
		)
	}
}
