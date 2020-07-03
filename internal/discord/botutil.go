package discord

import (
	"fmt"

	"github.com/Floor-Gang/prevent-link-redirection/internal"
	dg "github.com/bwmarrin/discordgo"
)

// Util struct methods

func (b *Bot) reply(event *dg.MessageCreate, message string) {
	_, err := b.session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("<@%s> %s", event.Author.ID, message))
	if err != nil {
		internal.Report(err)
	}
}

func (b *Bot) checkChannel(commandMessage *dg.Message) bool {
	return commandMessage.ChannelID == b.config.ChannelID
}

func (b *Bot) checkRoles(member *dg.Member) bool {
	return internal.StringInSlice(b.config.LeadDevID, member.Roles) ||
		internal.StringInSlice(b.config.AdminID, member.Roles)
}

func checkAction(action string) bool {
	if action == "remove" || action == "filter" {
		return true
	}
	return false
}
