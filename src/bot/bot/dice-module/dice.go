package dice

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/dice-module/calc"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a bot module that responds to "!roll" commands
type Module struct {
	Log        *log.Logger
	Calculator calc.DiceCalculator
	AppID      config.ApplicationID
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "dice" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	if err := m.RegisterApplicationCommand(ctx, session); err != nil {
		return err
	}
	cancel := session.AddHandler(m.handleCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("dice module context done")
		cancel()
	}()
	return nil
}

func (m *Module) roll(cmdCtx commandContext, verbose bool, query string) error {
	result, err := m.Calculator.Calculate(query)
	if err != nil {
		return util.DiscordMessageSendRawBlock(cmdCtx.session, cmdCtx.channelID, fmt.Sprintf("calculator error: %v", err))
	}

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Dice roll: `%s`\n\n", query)

	if verbose {
		var strTokens []string
		strTokens = []string{}
		for _, t := range result.InfixTokens {
			strTokens = append(strTokens, t.String)
		}
		fmt.Fprintf(buf, "Infix interpretation: `%s`\n", strings.Join(strTokens, " "))

		strTokens = []string{}
		for _, t := range result.PostfixTokens {
			strTokens = append(strTokens, t.String)
		}
		fmt.Fprintf(buf, "Postfix interpretation: `%s`\n", strings.Join(strTokens, " "))
	}

	fmt.Fprintf(buf, "**Result: %d**\n\n", result.Value)

	if len(result.Rolls) > 0 {
		fmt.Fprintf(buf, "Actual dice rolled:\n")
		for _, r := range result.Rolls {
			if len(r.Results) > 50 {
				fmt.Fprintf(buf, "- %dd%d ⮕ (first 50 of %d) %v\n", r.Count, r.Sides, len(r.Results), r.Results[:50])

			} else {
				fmt.Fprintf(buf, "- %dd%d ⮕ %v\n", r.Count, r.Sides, r.Results)
			}
		}
	}

	_, err = cmdCtx.session.ChannelMessageSend(cmdCtx.channelID, buf.String())
	return err
}
