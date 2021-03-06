package discord

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Run(token string, rootdir string) {
	bot, err := discordgo.New("Bot " + os.Getenv("DISCORDFS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(rootdir, 0700)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	initMessageCreate(rootdir)
	bot.AddHandler(messageCreate)

	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("discordfs is now running\npress CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}
