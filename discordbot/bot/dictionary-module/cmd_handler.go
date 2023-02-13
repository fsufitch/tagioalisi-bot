package dictionary

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
		Name:        "!define",
		Usage:       "Look up a definition in the Merriam-Webster dictionary",
		Writer:      stdout,
		ErrWriter:   stderr,
		HideVersion: true,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
		},
		ArgsUsage: "[word]",
		Commands:  []*cli.Command{},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ChannelID, "No word specified. Try: !define --help")
			}

			return m.define(ctx, c.Args().First())
		},
	}
	m.Log.Debugf("dictionary: created urfave/cli command for message %v", ctx.messageCreate.ID)

	return
}

type commandContext struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

func (m Module) handleCommand(s *discordgo.Session, event *discordgo.MessageCreate) {
	fields := strings.Fields(event.Message.Content)
	if len(fields) < 1 || fields[0] != "!define" {
		return
	}

	m.Log.Debugf("dictionary: message %v triggers !define: %s", event.ID, event.Content)
	m.Log.Debugf("dictionary: channel %v", event.ChannelID)

	if event.Message.Author == nil || event.Message.Author.Bot {
		m.Log.Debugf("dictionary: message %v ignored due to nil/bot author", event.ID)
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.Log.Errorf("dictionary: message %v error while running cli: %v", event.ID, err)
	}

	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		m.Log.Errorf("dictionary: message %v error output while executing command: %s", event.ID, string(errData))
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData)); err != nil {
			m.Log.Errorf("dictionary: message %v error while sending stdout: %v", event.ID, err)
		}

	}
	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData)); err != nil {
			m.Log.Errorf("dictionary: message %v error while sending stderr: %v", event.ID, err)
		}
	}
}
