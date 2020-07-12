package discord

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessage(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.ChannelID != b.config.ChannelID {
		// Check if the message is a URL
		// TODO: add support for <https://URL>
		match, _ := regexp.MatchString("^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$", m.Content)
		if match {
			// Send GET request to the URL
			resp, err := http.Get(m.Content)
			if err != nil {
				log.Fatalf("http.Get => %v", err.Error())
			}
			finalURL := resp.Request.URL.String()

			// Check if the response URL is the same as request url
			if finalURL != m.Content || strings.Contains(m.Content, "redirect") {
				s.ChannelMessageSend(b.config.ChannelID, m.Content+" is redirects to: "+finalURL)
			}
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
