package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()
	discord.Token = ""
	if err != nil {
		panic(err)
	}

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}
