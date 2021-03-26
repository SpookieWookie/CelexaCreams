package handler

import (
	"github.com/morganonbass/celexacreams/celexacreams"

	"github.com/bwmarrin/discordgo"
)

// Snack responds to "snack"
type Snack struct{}

// Handle shows snack
func (h *Snack) Handle(m *discordgo.MessageCreate) (string, error) {
	url, err := celexacreams.GetRandomGIF("cat eating")
	if err != nil {
		return "", &celexacreams.CelexaError{
			"GIF error: " + err.Error(),
		}
	}
	return m.Author.Mention() + " " + url, nil
}
