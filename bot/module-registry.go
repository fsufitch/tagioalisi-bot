package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/bot/ping-module"
	"github.com/fsufitch/discord-boar-bot/bot/sockpuppet-module"
	"github.com/fsufitch/discord-boar-bot/common"
)

// RegisterableModule is a generic interface for a registerable modular piece of bot functionality
type RegisterableModule interface {
	Name() string
	Register(*discordgo.Session) error
}

// ModuleRegistry is an aliased slice of modules, for injection purposes
type ModuleRegistry []RegisterableModule

// InitModuleRegistry initializes the module registry and blacklists appropriate modules
func InitModuleRegistry(
	configuration *common.Configuration,
	log *common.LogDispatcher,
	ping *ping.Module,
	sockpuppet *sockpuppet.Module,
) ModuleRegistry {
	unfilteredModules := []RegisterableModule{
		ping,
		sockpuppet,
	}

	filteredModules := []RegisterableModule{}

	for _, module := range unfilteredModules {
		if v, ok := configuration.BlacklistBotModules[module.Name()]; !ok || !v {
			filteredModules = append(filteredModules, module)
		} else {
			log.Info("Blacklisted bot module: " + module.Name())
		}
	}

	return ModuleRegistry(filteredModules)
}
