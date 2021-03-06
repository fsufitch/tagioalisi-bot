package news

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
		Name:        "!news",
		Usage:       "Search for recent news articles",
		Writer:      stdout,
		ErrWriter:   stderr,
		HideVersion: true,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
		},
		ArgsUsage: "[search terms]",
		Commands:  []*cli.Command{},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "print full previews",
			},
			&cli.IntFlag{
				Name:    "count",
				Aliases: []string{"c"},
				Usage:   "the maximum number of results to get",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ChannelID, "No search terms specified. Try: !news --help")
			}

			count := c.Int("count")
			if count == 0 {
				count = 3
			}
			verbose := c.Bool("verbose")
			query := strings.Join(c.Args().Slice(), " ")

			return m.DoSearch(c.Context, ctx.session, ctx.messageCreate.ChannelID, query, count, verbose)
		},
	}
	m.Log.Debugf("news: created urfave/cli command for message %v", ctx.messageCreate.ID)

	return
}

type commandContext struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

func (m Module) handleCommand(s *discordgo.Session, event *discordgo.MessageCreate) {
	fields := strings.Fields(event.Message.Content)
	if len(fields) < 1 || fields[0] != "!news" {
		return
	}

	m.Log.Debugf("news: message %v triggers !news: %s", event.ID, event.Content)

	if event.Message.Author == nil || event.Message.Author.Bot {
		m.Log.Debugf("news: message %v ignored due to nil/bot author", event.ID)
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.Log.Errorf("news: message %v error while running cli: %v", event.ID, err)
	}

	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		m.Log.Errorf("news: message %v error output while executing command: %s", event.ID, string(errData))
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData)); err != nil {
			m.Log.Errorf("news: message %v error while sending stdout: %v", event.ID, err)
		}

	}
	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData)); err != nil {
			m.Log.Errorf("news: message %v error while sending stderr: %v", event.ID, err)
		}
	}
}
