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
		return // Ignore bots
	}

	if strings.HasPrefix(strings.TrimSpace(event.Message.Content), "!") {
		return // Ignore bot commands
	}

	uniqueMemes := map[string]string{}
	for _, fileName := range memeRegex.FindAllString(event.Message.Content, -1) {
		if parts := strings.SplitN(strings.ToLower(fileName), ".", 2); len(parts) < 1 {
			m.Log.Errorf("memelink: somehow found meme without a dot: %v", fileName)
			continue
		} else {
			uniqueMemes[parts[0]] = fileName
		}
	}

	memeCount := 0
	for memeName, memeFileName := range uniqueMemes {
		meme, err := m.MemeDAO.SearchByName(memeName)
		if err != nil {
			m.Log.Errorf("error searching for meme `%s`: %v", memeName, err)
			continue
		}
		if meme == nil {
			continue
		}
		m.Log.Debugf("memelink: found meme with name `%v`", memeName)

		url := meme.URLs[rand.Intn(len(meme.URLs))]

		// TODO: figure out embeds; they seem to only work on direct link/video images
		message := fmt.Sprintf("**%s:** %s", memeFileName, url.URL)
		_, err = s.ChannelMessageSend(event.Message.ChannelID, message)
		if err != nil {
			m.Log.Errorf("error sending meme: %v", err)
			return
		}
		memeCount++
		if memeCount >= 3 {
			m.Log.Debugf("memelink: capping memes per proc at 3")
			break
		}
	}
	if memeCount >= 3 {
		if _, err := s.ChannelMessageSend(event.Message.ChannelID, "Too many memes in one message! Chill out!"); err != nil {
			m.Log.Errorf("error sending throttle message: %v", err)
		}
	}
	return
}
