package discord

import (
	"github.com/Floor-Gang/prevent-link-redirection/internal"
	dg "github.com/bwmarrin/discordgo"
)

// Bot structure
type Bot struct {
	version string
	session *dg.Session
	config  internal.Config
}

// Start starts discord client, configuration and database
func Start(configPath string) error {
	var err error
	botConfig := internal.GetConfig(configPath)

	client, err := dg.New("Bot " + botConfig.Token)

	if err != nil {
		panic(err)
	}

	bot := Bot{
		session: client,
		config:  botConfig,
	}

	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onMessage)

	if err = client.Open(); err != nil {
		panic(err)
	}

	return err
}
