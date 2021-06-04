package handler

import "github.com/bwmarrin/discordgo"

func createEmojiList(message string) {
}

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID || event.Content == "" {
		return
	}

	if event.Content == "愚か" {
		session.MessageReactionAdd(event.ChannelID, event.Message.ID, "848915057475715104")
		session.MessageReactionAdd(event.ChannelID, event.Message.ID, "848915069571301396")
	}
}
