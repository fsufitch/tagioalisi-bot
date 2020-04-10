package wiki

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/urfave/cli/v2"
)

func (m *Module) cliApp(ctx commandContext) (app *cli.App, stdout, stderr *bytes.Buffer) {
	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)

	app = &cli.App{
		Name:        "!wiki",
		Usage:       "Look up Wikimedia entries",
		Writer:      stdout,
		ErrWriter:   stderr,
		HideVersion: true,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "en",
				Usage:   "language to search using (if supported)",
			},
			&cli.StringFlag{
				Name:    "wiki",
				Aliases: []string{"w"},
				Value:   "wp",
				Usage:   "which wiki to query",
			},
			&cli.BoolFlag{
				Name:  "options",
				Usage: "print options for wikis to query",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("options") {
				return m.showOptions(ctx)
			}
			return nil
		},
	}
	m.Log.Debugf("groups: created urfave/cli command for message %v", ctx.messageCreate.ID)

	return
}

type commandContext struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

func (m Module) handleCommand(s *discordgo.Session, event *discordgo.MessageCreate) {
	fields := strings.Fields(event.Message.Content)
	if len(fields) < 1 || fields[0] != "!wiki" {
		return
	}

	m.Log.Debugf("wiki: message %v triggers !wiki: %s", event.ID, event.Content)

	if event.Message.Author == nil || event.Message.Author.Bot {
		m.Log.Debugf("wiki: message %v ignored due to nil/bot author", event.ID)
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.Log.Errorf("wiki: message %v error while running cli: %v", event.ID, err)
	}

	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		m.Log.Errorf("wiki: message %v error output while executing command: %s", event.ID, string(errData))
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData)); err != nil {
			m.Log.Errorf("wiki: message %v error while sending stdout: %v", event.ID, err)
		}

	}
	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData)); err != nil {
			m.Log.Errorf("wiki: message %v error while sending stderr: %v", event.ID, err)
		}
	}
}
