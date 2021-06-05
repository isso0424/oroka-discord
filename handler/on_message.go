package handler

import (
	"fmt"
	"isso0424/oroka-discord/config"

	"github.com/bwmarrin/discordgo"
)

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID || event.Content == "" {
		return
	}

	for _, pattern := range config.ReactionPatterns {
		if event.Content == pattern.Message {
			for _, emoji := range pattern.Emojis {
				err := session.MessageReactionAdd(event.ChannelID, event.Message.ID, emoji)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
