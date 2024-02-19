package groups

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const gPrefix = "g-"

func isGroup(role string) bool {
	return strings.HasPrefix(role, gPrefix)
}

func groupToRole(group string) string {
	return gPrefix + group
}

func roleToGroup(role string) string {
	if !isGroup(role) {
		return ""
	}
	return role[len(gPrefix):]
}

func getRoleByName(s *discordgo.Session, guildID string, roleName string) (*discordgo.Role, error) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	for _, r := range roles {
		if r.Name == roleName {
			return r, nil
		}
	}
	return nil, errNoSuchGroup
}

func rolesToGroups[T any](roles []T) []string {
	groups := []string{}
	for _, role := range roles {
		var roleName string
		var untypedRole any = role
		switch typedRole := untypedRole.(type) {
		case string:
			roleName = typedRole
		case *discordgo.Role:
			roleName = typedRole.Name
		case discordgo.Role:
			roleName = typedRole.Name
		default:
			continue
		}

		if isGroup(roleName) {
			groups = append(groups, roleToGroup(roleName))
		}
	}
	return groups
}

// func getMemberGroups(s *discordgo.Session, guildID string, memberID string) ([]string, error) {
// 	member, err := s.GuildMember(guildID, memberID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return rolesToGroups(member.Roles), nil
// }

func allGroups(s *discordgo.Session, guildID string) ([]string, error) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	return rolesToGroups(roles), nil
}

var errNoSuchGroup = errors.New("no such group")
