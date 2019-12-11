package memelink

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/urfave/cli/v2"
)

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
			},
			{
				Name:      "search",
				Usage:     "search the meme database, receiving results in a private message",
				ArgsUsage: "query",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "all",
						Usage: "display all memes",
					},
				},
			},
			// TODO: more commands, especially delete ones
		},
		Writer:    stdout,
		ErrWriter: stderr,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
		},
	}

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

	if event.Message.Author == nil || event.Message.Author.Bot {
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.log.Error(fmt.Sprintf("error while running !memes cli (`%s`): %v", event.Message.Content, err))
	}

	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		common.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData))
	}
	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		common.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData))
	}
}

func (m Module) handleAddMeme(s *discordgo.Session, event *discordgo.MessageCreate,
	name string, url string, appendOK bool) error {
	s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Adding meme `%s` -> `%s`", name, url))
	if err := m.memeDAO.Add(name, url, event.Author.String()); err != nil {
		s.ChannelMessageSend(event.Message.ChannelID, fmt.Sprintf("Error adding meme: %v", err))
	}
	return nil
}
