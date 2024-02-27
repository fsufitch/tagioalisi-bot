package groupscommand

import "github.com/bwmarrin/discordgo"

var cmdLeaveGroup = &discordgo.ApplicationCommandOption{
	Name:        "leave",
	Description: "leave a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "group",
			Description:  "group to leave",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     true,
			Autocomplete: true,
		},
	},
}
