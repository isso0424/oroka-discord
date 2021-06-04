package main

import (
	"isso0424/oroka-discord/handler"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()
	discord.Token = os.Getenv("TOKEN")
	if err != nil {
		panic(err)
	}
	discord.AddHandler(handler.OnMessageHandler)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}
