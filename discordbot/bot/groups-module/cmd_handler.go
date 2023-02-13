package groups

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
		Name:  "!groups",
		Usage: "Manipulate user groups",
		Commands: []*cli.Command{
			{
				Name:      "create",
				Usage:     "create a new user group (admins only)",
				ArgsUsage: "group_name",
				Action: func(cliCtx *cli.Context) error {
					return m.groupCreate(ctx.session, ctx.messageCreate, cliCtx.Args().Get(0))
				},
			},
			{
				Name:      "delete",
				Usage:     "delete a user group (admins only)",
				ArgsUsage: "group_name",
				Action: func(cliCtx *cli.Context) error {
					return m.groupDelete(ctx.session, ctx.messageCreate, cliCtx.Args().Get(0))
				},
			},
			{
				Name:  "list",
				Usage: "list all the registered groups, receiving them in a private message",
				Action: func(cliCtx *cli.Context) error {
					return m.groupList(ctx.session, ctx.messageCreate)
				},
			},
			{
				Name:      "join",
				Usage:     "join either yourself or someone else (admin only) to a group",
				ArgsUsage: "group_name [user_mentions...]",
				Action: func(cliCtx *cli.Context) error {
					return m.groupJoin(ctx.session, ctx.messageCreate, cliCtx.Args().Get(0))
				},
			},
			{
				Name:      "leave",
				Usage:     "remove either yourself or someone else (admin only) from a group",
				ArgsUsage: "group_name [user_mentions...]",
				Action: func(cliCtx *cli.Context) error {
					return m.groupLeave(ctx.session, ctx.messageCreate, cliCtx.Args().Get(0))
				},
			},
		},
		Writer:      stdout,
		ErrWriter:   stderr,
		HideVersion: true,
		CommandNotFound: func(context *cli.Context, command string) {
			fmt.Fprintf(stderr, "Unknown command for `%s`: %s `%s`\n", context.App.Name, context.Command.Name, command)
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
	if len(fields) < 1 || fields[0] != "!groups" {
		return
	}

	m.Log.Debugf("groups: message %v triggers !groups: %s", event.ID, event.Content)

	if event.Message.Author == nil || event.Message.Author.Bot {
		m.Log.Debugf("groups: message %v ignored due to nil/bot author", event.ID)
		return
	}

	cmd, stdout, stderr := m.cliApp(commandContext{s, event})
	if err := cmd.Run(fields); err != nil {
		m.Log.Errorf("groups: message %v error while running cli: %v", event.ID, err)
	}

	if errData, _ := ioutil.ReadAll(stderr); len(errData) > 0 {
		m.Log.Errorf("groups: message %v error output while executing groups command: %s", event.ID, string(errData))
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(errData)); err != nil {
			m.Log.Errorf("groups: message %v error while sending stdout: %v", event.ID, err)
		}

	}
	if stdData, _ := ioutil.ReadAll(stdout); len(stdData) > 0 {
		if err := util.DiscordMessageSendRawBlock(s, event.Message.ChannelID, string(stdData)); err != nil {
			m.Log.Errorf("groups: message %v error while sending stderr: %v", event.ID, err)
		}
	}
}
