package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func createEmojiList(message string) {
}

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID || event.Content == "" {
		return
	}

	if event.Content == "愚か" {
		err := session.MessageReactionAdd(event.ChannelID, event.Message.ID, ":oro:848915057475715104")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = session.MessageReactionAdd(event.ChannelID, event.Message.ID, ":ka:848915069571301396")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
