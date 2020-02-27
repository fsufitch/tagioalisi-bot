package util

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DiscordMessageSendRawBlock sends an arbitrary raw block of "pre" formatted
// text, split into multiple messages if necessaty
func DiscordMessageSendRawBlock(s *discordgo.Session, channelID string, text string) error {
	escapedText := strings.Replace(text, "```", "###", -1)
	chunks := Chunk(escapedText)
	for _, chunk := range chunks {
		if _, err := s.ChannelMessageSend(channelID, fmt.Sprintf("```%s```", chunk)); err != nil {
			return err
		}

	}
	if escapedText != text {
		if _, err := s.ChannelMessageSend(channelID, "```[[ Message contained triple backticks, which were escaped to '###' ]]]```"); err != nil {
			return err
		}
	}
	return nil
}
