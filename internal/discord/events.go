package discord

import (
	"log"
	"regexp"

	dg "github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessage(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.ChannelID != b.config.ChannelID {
		match, _ := regexp.MatchString("^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$", m.Content)
		if match {
			s.ChannelMessageSend(b.config.ChannelID, "It's a URL! "+m.Content)
		}
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" && m.ChannelID != b.config.ChannelID {
		s.ChannelMessageSend(b.config.ChannelID, "Ping!")
	}
}

func (b *Bot) onReady(s *dg.Session, _ *dg.Ready) {
	log.Printf("Ready as %s (version %s)\n", s.State.User.Username, b.version)
}
