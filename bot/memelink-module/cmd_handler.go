package memelink

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagialisi-bot/bot/util"
	"github.com/fsufitch/tagialisi-bot/db/memes-dao"
	"github.com/urfave/cli/v2"
)

var urlRegex = regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

func (m *Module) cliApp(ctx commandContext) (app *cli.App, stdout, stderr *bytes.Buffer) {
	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)

	app = &cli.App{
		Name:  "!memes",
		Usage: "Manipulate the meme database",
		Commands: []*cli.Command{
			{
				Name:      "add",
				Usage:     "add a new meme",
				ArgsUsage: "meme_name meme_url",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "append",
						Aliases: []string{"a"},
						Usage:   "on name clash, add URL to existing meme",
					},
				},
				Action: func(cliCtx *cli.Context) error {
					return m.handleAddMeme(ctx.session, ctx.messageCreate,
						cliCtx.Args().Get(0), cliCtx.Args().Get(1), cliCtx.Bool("append"))
				},
			},
			{
				Name:      "alias",
				Usage:     "add a name to an existing meme",
				ArgsUsage: "new_name meme",
				Action: func(cliCtx *cli.Context) error {
					return m.handleAddAlias(ctx.session, ctx.messageCreate,
						cliCtx.Args().Get(0), cliCtx.Args().Get(1))
				},
			},
			{
				Name:      "search",
				Usage:     "search the meme database, receiving results in a private message",
				ArgsUsage: "query",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "display all memes",
					},
				},
				Action: func(cliCtx *cli.Context) error {
					if cliCtx.Bool("all") {
						return m.handleSearch(ctx.session, ctx.messageCreate, "", true)
					}
					return m.handleSearch(ctx.session, ctx.messageCreate, cliCtx.Args().Get(0), false)
				},
			},
			// TODO: more commands, especially delete ones
		},
		Writer:      stdout,
		ErrWriter:   stderr,
		HideVersion: true,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
		},
	}
	m.Log.Debugf("memelink: created urfave/cli command for message %v", ctx.messageCreate.ID)

	return
}

type commandContext struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

func (m Module) handleCommand(s *discordgo.Session, event *discordgo.MessageCreate) {
	fields := strings.Fields(event.Message.Content)
	if len(fields) < 1 || fields[0] != "!memes" {
		return
	}
	m.Log.Debugf("memelink: message %v triggers !memes: %s", event.ID, event.Content)

	if event.Message.Author == nil || event.Message.Author.Bot {
		m.Log.Debugf("groups: message %v ignored due to nil/bot author", event.ID)
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.Log.Errorf("error while running !memes cli (`%s`): %v", event.Message.Content, err)
	}

	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData)); err != nil {
			m.Log.Errorf("groups: message %v error while sending stdout: %v", event.ID, err)
		}
	}
	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		m.Log.Errorf("Error output while executing memelink command: %s", string(errData))
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData)); err != nil {
			m.Log.Errorf("groups: message %v error while sending stderr: %v", event.ID, err)
		}
	}

}

func (m Module) handleAddMeme(s *discordgo.Session, event *discordgo.MessageCreate,
	name string, url string, appendOK bool) error {
	if !m.isMemeEditor(s, event.Author.ID, event.GuildID) {
		_, err := s.ChannelMessageSend(event.Message.ChannelID, "Sorry, you do not have meme editor permissions")
		return err
	}

	name = strings.TrimSpace(name)
	url = strings.TrimSpace(url)

	if name == "" {
		_, err := s.ChannelMessageSend(event.Message.ChannelID, "Meme name must not be empty")
		return err
	}
	if !urlRegex.MatchString(url) {
		_, err := s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Invalid URL: `%s`", url))
		return err
	}

	s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Adding meme `%s` -> `%s`", name, url))

	var err error
	appended := false

	if appendOK {
		existingMeme, errSearch := m.MemeDAO.SearchByName(name)
		if errSearch != nil {
			return errSearch
		}
		if existingMeme != nil {
			s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Adding URL to meme %d", existingMeme.ID))
			err = m.MemeDAO.AddURL(existingMeme.ID, url, event.Author.String())
			appended = true
		}
	}
	if !appended {
		err = m.MemeDAO.Add(name, url, event.Author.String())
	}

	if err != nil {
		s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Error adding meme: %v", err))
	}

	return nil
}

func (m Module) handleAddAlias(s *discordgo.Session, event *discordgo.MessageCreate,
	newName string, oldName string) error {
	if !m.isMemeEditor(s, event.Author.ID, event.GuildID) {
		_, err := s.ChannelMessageSend(event.Message.ChannelID, "Sorry, you do not have meme editor permissions")
		return err
	}

	if newName == "" || oldName == "" {
		_, err := s.ChannelMessageSend(event.Message.ChannelID, "aliasing requires two arguments for meme names (new name, old name)")
		return err
	}
	s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Adding alias `%s` -> `%s`", newName, oldName))

	if existingMeme, err := m.MemeDAO.SearchByName(oldName); err != nil {
		return err
	} else if existingMeme == nil {
		_, err = s.ChannelMessageSend(event.Message.ChannelID, "No meme found with the old alias")
		return err
	} else {
		return m.MemeDAO.AddName(existingMeme.ID, newName, event.Author.String())
	}

}

func (m Module) handleSearch(s *discordgo.Session, event *discordgo.MessageCreate,
	query string, all bool) error {

	var memeResults []memes.Meme
	var err error
	if all {
		memeResults, err = m.MemeDAO.SearchMany("")
	} else if query != "" {
		memeResults, err = m.MemeDAO.SearchMany(query)
	} else {
		_, err = s.ChannelMessageSend(event.Message.ChannelID, "No query specified. Please specify a query or `--all`/`-a`")
		return err
	}

	if err != nil {
		return err
	}

	if len(memeResults) == 0 {
		_, err = s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("No memes match the search `%s`", query))
		return err
	}

	lines := []string{}
	for _, meme := range memeResults {
		nameStrings := []string{}
		for _, name := range meme.Names {
			nameStrings = append(nameStrings, name.Name)
		}

		lines = append(lines, fmt.Sprintf("=== [%d] - %s", meme.ID, strings.Join(nameStrings, ", ")))
		for _, url := range meme.URLs {
			lines = append(lines, fmt.Sprintf(" - %s", url.URL))
		}
	}

	lineGroups := util.ChunkLines(lines)
	ch, err := s.UserChannelCreate(event.Author.ID)
	if err != nil {
		return err
	}

	for _, lineGroup := range lineGroups {
		message := fmt.Sprintf("```%s```", strings.Join(lineGroup, "\n"))
		if _, err := s.ChannelMessageSend(ch.ID, message); err != nil {
			return err
		}
	}
	_, err = s.ChannelMessageSend(event.Message.ChannelID, "Results sent via PM.")
	return err
}
