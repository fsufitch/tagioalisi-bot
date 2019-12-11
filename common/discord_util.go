package common

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DiscordMessageSendRawBlock sends an arbitrary raw block of "pre" formatted
// text, split into multiple messages if necessaty
func DiscordMessageSendRawBlock(s *discordgo.Session, channelID string, text string) {
	escapedText := strings.Replace(text, "```", "###", -1)
	chunks := Chunk(escapedText)
	for _, chunk := range chunks {
		s.ChannelMessageSend(channelID, fmt.Sprintf("```%s```", chunk))
	}
	if escapedText != text {
		s.ChannelMessageSend(channelID, "```[[ Message contained triple backticks, which were escaped to '###' ]]]```")
	}
}
