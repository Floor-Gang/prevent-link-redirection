package internal

import (
	dg "github.com/bwmarrin/discordgo"
)

// Help function
func (b *Bot) help(event *dg.MessageCreate) {
	_, err := b.helpEmbed()
	if err != nil {
		Report(err)
		b.reply(event, "Something went wrong.")
	}
}

// Possibly for later:
// - Set mention in db to active
// - Set mention in db to inactive
