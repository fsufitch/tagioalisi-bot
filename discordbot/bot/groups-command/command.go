package groupscommand

import (
	"github.com/bwmarrin/discordgo"

	guildcache "github.com/fsufitch/tagioalisi-bot/bot/guild-cache"
	"github.com/fsufitch/tagioalisi-bot/bot/util/interactions"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

type GroupsCommandModule struct {
	*interactions.ApplicationCommandWrapper

	*Prefixer
	*log.Logger
	*guildcache.GuildCacheManager
}

func ProvideApplicationCommand(logger *log.Logger, prefixer *Prefixer, appID config.ApplicationID, gcm *guildcache.GuildCacheManager) *GroupsCommandModule {
	cmd := &GroupsCommandModule{
		ApplicationCommandWrapper: &interactions.ApplicationCommandWrapper{
			Logger: logger,
			Command: &discordgo.ApplicationCommand{
				ApplicationID: string(appID),
				Name:          "groups2",
				Description:   "interact with interest groups",
				Options: []*discordgo.ApplicationCommandOption{
					cmdJoinGroup,
					cmdLeaveGroup,
				},
			},
		},
		Logger:            logger,
		Prefixer:          prefixer,
		GuildCacheManager: gcm,
	}

	cmd.ApplicationCommandWrapper.Handlers = []interactions.HandlerFunc{
		cmd.subcommandJoin,
	}
	cmd.ApplicationCommandWrapper.AutocompleteHandlers = []interactions.AutocompleteHandlerFunc{
		cmd.subcommandJoinAutocomplete,
	}

	return cmd
}

func (mod *GroupsCommandModule) CommandWrapper() *interactions.ApplicationCommandWrapper {
	return mod.ApplicationCommandWrapper
}
