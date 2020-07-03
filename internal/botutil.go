package internal

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Util struct methods

func (b *Bot) reply(event *dg.MessageCreate, message string) {
	_, err := b.session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("<@%s> %s", event.Author.ID, message))
	if err != nil {
		Report(err)
	}
}

func (b *Bot) checkChannel(commandMessage *dg.Message) bool {
	return commandMessage.ChannelID == b.config.ChannelID
}

// helpEmbed makes an embed with the mentionMessage
func (b *Bot) helpEmbed() (*dg.Message, error) {
	embed := dg.MessageEmbed{
		Author: &dg.MessageEmbedAuthor{},
		Color:  0xff0000,
		Fields: []*dg.MessageEmbedField{
			{
				Name:   "Add a mention",
				Value:  "`.mention add <regex> <action> <description>`",
				Inline: false,
			},
			{
				Name:   "Remove a mention",
				Value:  "`.mention remove <id>`",
				Inline: false,
			},
			{
				Name:   "Display all mentions",
				Value:  "`.mention mentions`",
				Inline: false,
			},
			{
				Name:   "Display a singular mention",
				Value:  "`.mention mention <id>`",
				Inline: false,
			},
			{
				Name:   "Change what happens on mention",
				Value:  "`.mention change_action <id> <type (filter/remove)>`",
				Inline: false,
			},
			{
				Name:   "Change regex of mention",
				Value:  "`.mention change_regex <id> <regex>`",
				Inline: false,
			},
			{
				Name:   "Change description of mention",
				Value:  "`.mention change_description <id> <description>`",
				Inline: false,
			},
			{
				Name:   "Display this message",
				Value:  "`.mention help`",
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     ".mention help",
	}

	msg, err := b.session.ChannelMessageSendEmbed(b.config.ChannelID, &embed)

	if err != nil {
		Report(err)
		return nil, err
	}

	return msg, nil
}

func (b *Bot) checkRoles(member *dg.Member) bool {
	return StringInSlice(b.config.LeadDevID, member.Roles) ||
		StringInSlice(b.config.AdminID, member.Roles)
}

func checkAction(action string) bool {
	if action == "remove" || action == "filter" {
		return true
	}
	return false
}
