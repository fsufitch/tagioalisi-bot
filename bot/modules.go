package bot

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagialisi-bot/bot/groups-module"
	"github.com/fsufitch/tagialisi-bot/bot/log-module"
	"github.com/fsufitch/tagialisi-bot/bot/memelink-module"
	"github.com/fsufitch/tagialisi-bot/bot/ping-module"
	"github.com/fsufitch/tagialisi-bot/bot/sockpuppet-module"
)

// Module is a generic interface for a registerable modular piece of bot functionality
type Module interface {
	Name() string
	Register(context.Context, *discordgo.Session) error
}

// Modules is a struct containing all the possible implemented modules
type Modules struct {
	Ping       *ping.Module
	Log        *log.Module
	SockPuppet *sockpuppet.Module
	MemeLink   *memelink.Module
	Groups     *groups.Module
}

// ModuleList is a list containing all the possible implemented modules
type ModuleList []Module

// ProvideModuleList builds ModuleList
func ProvideModuleList(m Modules) ModuleList {
	return ModuleList{
		m.Ping,
		m.Log,
		m.SockPuppet,
		m.MemeLink,
		m.Groups,
	}
}
