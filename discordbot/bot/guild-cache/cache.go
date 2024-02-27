package guildcache

import (
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/log"
)

const defaultTTL time.Duration = 30 * time.Second

type GuildCacheManager struct {
	guildCaches map[string]*GuildCache
	Logger      *log.Logger
}

func ProvideGuildCacheManager(logger *log.Logger) *GuildCacheManager {
	return &GuildCacheManager{guildCaches: map[string]*GuildCache{}, Logger: logger}
}

func (mgr *GuildCacheManager) Cache(guildID string) *GuildCache {
	if mgr.guildCaches == nil {
		mgr.guildCaches = map[string]*GuildCache{}
	}
	if gcache, ok := mgr.guildCaches[guildID]; !ok || gcache == nil {
		mgr.guildCaches[guildID] = &GuildCache{
			guildID: guildID,
			logger:  mgr.Logger,
		}
	}
	return mgr.guildCaches[guildID]
}

type GuildCache struct {
	guildID string
	logger  *log.Logger
	TTL     time.Duration

	rolesFetchedTime time.Time
	roles            []*discordgo.Role
}

func (gc GuildCache) GuildID() string {
	return gc.guildID
}

func (gc GuildCache) Expired() bool {
	ttl := gc.TTL
	if ttl == 0 {
		ttl = defaultTTL
	}
	gc.logger.Debugf("fetched=%d ttl=%d now=%d", gc.rolesFetchedTime, ttl, time.Now())
	gc.logger.Debugf("exp=%d  expired? %+v", gc.rolesFetchedTime.Add(ttl), time.Now().After(gc.rolesFetchedTime.Add(ttl)))
	return time.Now().After(gc.rolesFetchedTime.Add(ttl))
}

func (gc *GuildCache) RefreshRoles(s *discordgo.Session) error {
	roles, err := s.GuildRoles(gc.GuildID())
	if err != nil {
		return err
	}
	gc.rolesFetchedTime = time.Now()
	gc.logger.Debugf("refreshed roles; now have %d", len(roles))
	gc.roles = roles
	return nil
}

func (gc *GuildCache) RefreshRolesIfExpired(s *discordgo.Session) (expired bool, err error) {
	if expired = gc.Expired(); !expired {
		return
	}
	return expired, gc.RefreshRoles(s)
}

func (gc *GuildCache) Roles(s *discordgo.Session, selFunc RoleSelector) ([]*discordgo.Role, error) {
	_, err := gc.RefreshRolesIfExpired(s)
	if err != nil {
		return nil, err
	}
	selectedRoles := []*discordgo.Role{}
	for _, role := range gc.roles {
		if !selFunc(role) {
			continue
		}
		selectedRoles = append(selectedRoles, role)
	}
	return selectedRoles, nil
}

func (gc *GuildCache) Role(s *discordgo.Session, selFunc RoleSelector) (*discordgo.Role, error) {
	_, err := gc.RefreshRolesIfExpired(s)
	if err != nil {
		return nil, err
	}
	gc.logger.Debugf("got roles: guild=%+v roles=%+v", gc.guildID, gc.roles)
	for _, role := range gc.roles {
		gc.logger.Debugf("check role '%s': %v", role.Name, selFunc(role))
		if selFunc(role) {
			return role, nil
		}
	}
	return nil, ErrRoleNotFound
}

type RoleSelector func(*discordgo.Role) bool

func AnyRole() RoleSelector {
	return func(*discordgo.Role) bool { return true }
}

func HasID(roleID string) RoleSelector {
	return func(role *discordgo.Role) bool { return role.ID == roleID }
}

func HasName(roleName string) RoleSelector {
	return func(role *discordgo.Role) bool { return role.Name == roleName }
}

var ErrRoleNotFound = errors.New("role not found")
