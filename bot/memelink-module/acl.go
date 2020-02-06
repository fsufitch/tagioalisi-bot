package memelink

import "github.com/bwmarrin/discordgo"

const memeEditorACL = "boarbot.modules.memelink::edit"

func (m Module) isMemeEditor(s *discordgo.Session, userID string, guildID string) bool {
	if allowed, err := m.ACLDAO.CheckUserACL(userID, memeEditorACL); err != nil {
		m.Log.Errorf("could not check meme editor user ACL: %v", err)
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
