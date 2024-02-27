package groups

import (
	"errors"
	"fmt"
	"slices"
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

func userIsInGroup(s *discordgo.Session, guildID string, userID string, group string) (bool, error) {
	groups, err := getMemberGroups(s, guildID, userID)
	fmt.Printf("user %s groups: [%s]\n", userID, strings.Join(groups, ", "))
	if err != nil {
		return false, err
	}
	return slices.Contains(groups, group), nil
}

// func rolesToGroups[T any](roles []T) []string {
// 	groups := []string{}
// 	for _, role := range roles {
// 		var roleName string
// 		var untypedRole any = role
// 		fmt.Printf("name: %+v utrole:%+v\n", roleName, untypedRole)
// 		switch typedRole := untypedRole.(type) {
// 		case string:
// 			fmt.Println("was a string")
// 			roleName = typedRole
// 		case *discordgo.Role:
// 			fmt.Println("was *role")
// 			roleName = typedRole.Name
// 		case discordgo.Role:
// 			fmt.Println("was role")
// 			roleName = typedRole.Name
// 		default:
// 			continue
// 		}

// 		if isGroup(roleName) {
// 			groups = append(groups, roleToGroup(roleName))
// 		}
// 	}
// 	return groups
// }

func getMemberGroups(s *discordgo.Session, guildID string, memberID string) ([]string, error) {
	member, err := s.GuildMember(guildID, memberID)
	if err != nil {
		return nil, err
	}
	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}

	groups := []string{}

	for _, role := range guildRoles {
		if !slices.Contains(member.Roles, role.ID) {
			continue
		}
		groups = append(groups, roleToGroup(role.Name))
	}
	return groups, nil
}

func getGuildGroups(s *discordgo.Session, guildID string) ([]string, error) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	groups := []string{}
	for _, role := range roles {
		group := roleToGroup(role.Name)
		fmt.Printf("guild %s has role %s (id: %s), which is group %s\n", guildID, role.Name, role.ID, group)
		if group == "" {
			continue
		}
		groups = append(groups, group)
	}
	return groups, nil
}

var errNoSuchGroup = errors.New("no such group")
