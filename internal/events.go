package internal

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func (bot *Bot) onMessage(session *dg.Session, msg *dg.MessageCreate) {
	if msg.Author.Bot { return; }
	// check if they provided a prefix and they're not a bot
	if msg.Author.Bot || !strings.HasPrefix(msg.Content, bot.Config.Prefix) {
		// Check if the message is a URL
		URLRegex := regexp.MustCompile("(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?")
		match := URLRegex.FindStringSubmatch(msg.Content)
		if len(match)>0 && (strings.HasPrefix(match[0], "http://") || strings.HasPrefix(match[0], "https://")) {
			// Send GET request to the URL
			resp, err := http.Get(match[0])
			if err != nil {
				log.Fatalf("http.Get => %v", err.Error())
			}
			finalURL := resp.Request.URL.String()
			msgAuthor:= msg.Author.ID
			botMessage := ""
			// Check if URL has "redirect" in it
			if strings.Contains(match[0], "redirect") {
				botMessage = "sent this <" + match[0] + "> which contains redirect";
			// Check if the response URL is the same as request url
			} else if finalURL != match[0] {
				botMessage = "sent this <" + match[0] + "> which redirects to <" + finalURL + ">";
			} 
			util.Mention(session, msgAuthor, bot.Config.NotificationChannel, botMessage)
		}
	}

	// we can ask the authentication server if they're an admin of the bot
	// isAdmin defaults to "false" if err is not nil
	isAdmin, err := bot.Auth.CheckMember(msg.Author.ID)

	// let's split their message up into arguments
	// args = [prefix, sub-command name]
	args := strings.Fields(msg.Content)

	if len(args) < 2 { // this would mean args is [prefix] which at that point ignore them
		return
	}

	// we can now find out what command they were calling
	switch args[1] {
	case "help":
		// if all your commands are admin-relate then just wrap the whole switch statement
		// with this if statement

		// you can also just do "if isAdmin", the error doesn't matter that much if you want cleaner
		// code.
		if err != nil {
			util.Reply(bot.Client, msg.Message, "Failed to contact auth server")
		} else if isAdmin {
			bot.helpMessage(msg.Message)
		} else {
			util.Reply(bot.Client, msg.Message, "You must be an admin to run this command.")
		}
	}
}

func (bot *Bot) onReady(_ *dg.Session, ready *dg.Ready) {
	log.Printf("client bot - ready as %s#%s", ready.User.Username, ready.User.Discriminator)
}
