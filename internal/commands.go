package internal

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	
)

// Example command
func (bot *Bot) helpMessage(msg *dg.Message) {
	util.Reply(bot.Client, msg, "Stop Redirect bot checks for messages that contain a link and reports the ones that are redirecting to another website.")
}

