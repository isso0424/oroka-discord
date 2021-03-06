package main

import (
	"isso0424/oroka-discord/config"
	"isso0424/oroka-discord/handler"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	config.Setup("patterns.json")
	discord, err := discordgo.New()
	discord.Token = "Bot " + os.Getenv("TOKEN")
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
