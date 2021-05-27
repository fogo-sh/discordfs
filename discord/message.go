package discord

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fogo-sh/discordfs/fs"
)

var messageCreatePath = "/on_message_create"
var messageCreateChan = make(chan []byte)

func initMessageCreate(rootdir string) {
	var path = rootdir + messageCreatePath
	fs.WriteFifo(path, messageCreateChan)
}

type CordMessageCreate struct {
	ID string `json:"id"`

	// The ID of the channel in which the message was sent.
	ChannelID string `json:"channel_id"`

	// The content of the message.
	Content string `json:"content"`

	// The author of the message. This is not guaranteed to be a
	// valid user (webhook-sent messages do not possess a full author).
	AuthorID string `json:"author_id"`
}

func serializeCreate(m *discordgo.MessageCreate) ([]byte, error) {
	return json.Marshal(CordMessageCreate{
		ID:        m.ID,
		ChannelID: m.ChannelID,
		AuthorID:  m.Author.ID,
		Content:   m.Content,
	})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	b, err := serializeCreate(m)
	if err != nil {
		fmt.Println("error:", err)
	}

	messageCreateChan <- append(b, []byte("\n")...)
}
