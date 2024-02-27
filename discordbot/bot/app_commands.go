package bot

import (
	"reflect"

	"github.com/bwmarrin/discordgo"

	groupscommand "github.com/fsufitch/tagioalisi-bot/bot/groups-command"
	"github.com/fsufitch/tagioalisi-bot/bot/util/interactions"
	"github.com/fsufitch/tagioalisi-bot/log"
)

type ApplicationCommandModule interface {
	CommandWrapper() *interactions.ApplicationCommandWrapper
}

type ApplicationCommandModuleBootstrapper struct {
	Log     *log.Logger
	modules []ApplicationCommandModule
}

func ProvideApplicationCommandModuleBootstrapper(
	logger *log.Logger,
	groupsMod *groupscommand.GroupsCommandModule,
	// ...
) ApplicationCommandModuleBootstrapper {
	return ApplicationCommandModuleBootstrapper{
		Log: logger,
		modules: []ApplicationCommandModule{
			groupsMod,
			// ...
		},
	}
}

func (bs ApplicationCommandModuleBootstrapper) registerToAllGuilds(s *discordgo.Session) (cancel func()) {
	var cases []reflect.SelectCase
	var moduleCancelFuncs []func()
	for _, mod := range bs.modules {
		ch, cancel := mod.CommandWrapper().AutoRegisterToSessionGuilds(s)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
		moduleCancelFuncs = append(moduleCancelFuncs, cancel)
	}

	go func() {
		for {
			idx, val, ok := reflect.Select(cases)
			if idx < 0 || idx >= len(bs.modules) {
				panic("invalid index from select")
			}
			mod := bs.modules[idx]
			commandName := "UNKNOWN"
			if mod.CommandWrapper() != nil && mod.CommandWrapper().Command != nil {
				commandName = mod.CommandWrapper().Command.Name
			}
			if !ok {
				mod.CommandWrapper().Logger.Criticalf("select failed; cmd-id='%s'", commandName)
				continue
			}
			result, ok := val.Interface().(interactions.ApplicationCommandRegistration)
			if !ok {
				bs.Log.Criticalf("registration result was of wrong type; cmd-id='%s'", commandName)
				continue
			}
			if result.RegisterError != nil {
				bs.Log.Errorf("registration error (cmd-id='%s' guild-id='%s'): %v", commandName, result.GuildID, result.RegisterError)
				continue
			}

			bs.Log.Infof("app command registered; cmd-id='%s' guild-id='%s'", commandName, result.GuildID)
		}
	}()

	cancel = func() {
		for _, cancelFunc := range moduleCancelFuncs {
			cancelFunc()
		}
	}

	return cancel
}
