package main

import (
	"log"

	"github.com/fogo-sh/discordfs/discord"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	discord.Cli()
}
