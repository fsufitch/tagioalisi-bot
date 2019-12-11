package memelink

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var memeRegex = regexp.MustCompile("[a-zA-Z0-9_-]+(?:\\.[a-zA-Z0-9_-]+)+")

func (m Module) handleLink(s *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Message.Author == nil || event.Message.Author.Bot {
		m.log.Debug("memelink: Ignoring message from bot")
		return // Ignore bots
	}

	if strings.HasPrefix(strings.TrimSpace(event.Message.Content), "!") {
		return // Ignore bot commands
	}

	uniqueMemes := map[string]string{}
	for _, fileName := range memeRegex.FindAllString(event.Message.Content, -1) {
		if parts := strings.SplitN(strings.ToLower(fileName), ".", 2); len(parts) < 1 {
			m.log.Error("Somehow found meme without a dot: " + fileName)
			continue
		} else {
			uniqueMemes[parts[0]] = fileName
		}
	}

	memeCount := 0
	for memeName, memeFileName := range uniqueMemes {
		meme, err := m.memeDAO.SearchByName(memeName)
		if err != nil {
			m.log.Error(fmt.Sprintf("Error searching for meme `%s`: %v", memeName, err))
			continue
		}
		if meme == nil {
			continue
		}
		fmt.Println(meme)
		fmt.Println(meme.Names)
		fmt.Println(meme.URLs)
		url := meme.URLs[rand.Intn(len(meme.URLs))]
		fmt.Println("got")
		fmt.Println(url)

		_, err = s.ChannelMessageSendEmbed(event.Message.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: url.Author,
			},
			Image: &discordgo.MessageEmbedImage{
				URL: url.URL,
			},
			Title: memeFileName,
		})
		fmt.Println(err)
		memeCount++
		if memeCount >= 3 {
			break
		}
	}
	if memeCount >= 3 {
		s.ChannelMessageSend(event.Message.ChannelID, "Too many memes in one message! Chill out!")
	}

}
