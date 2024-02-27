package groupscommand

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNoSuchGroup = errors.New("no such group")

type Prefixer struct {
	Prefix string
}

func (p Prefixer) RemoveGroupRolePrefix(roleName string) (string, error) {
	if !strings.HasPrefix(roleName, p.Prefix) {
		return "", fmt.Errorf("%w: role=%+v", ErrNoSuchGroup, roleName)
	}
	return roleName[len(p.Prefix):], nil
}

func (p Prefixer) AddGroupRolePrefix(groupName string) string {
	return p.Prefix + groupName
}

func ProvideDefaultPrefixer() *Prefixer {
	return &Prefixer{Prefix: "g-"}
}
