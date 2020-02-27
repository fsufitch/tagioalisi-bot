package memelink

import "github.com/bwmarrin/discordgo"

const memeEditorACL = "boarbot.modules.memelink::edit" // TODO: rename to tagi namespace once ACL is more easily editable

// TODO: change to (bool, err) return type to bubble errors up
func (m Module) isMemeEditor(s *discordgo.Session, userID string, guildID string) bool {
	if allowed, err := m.ACLDAO.CheckUserACL(userID, memeEditorACL); err != nil {
		m.Log.Errorf("memelink: could not check meme editor user ACL: %v", err)
		return false
	} else if allowed {
		return true
	}

	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		m.Log.Errorf("could not obtain member from user: %v", err)
		return false
	}

	allowed, err := m.ACLDAO.CheckMultiRoleACL(member.Roles, memeEditorACL)
	if err != nil {
		m.Log.Errorf("could not check meme editor roles ACL: %v", err)
		return false
	}
	return allowed
}
