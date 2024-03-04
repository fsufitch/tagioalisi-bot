package acl

import (
	"database/sql"
	"fmt"
	"strings"
)

// RoleACL encapsulates a single permission granted to a role
type RoleACL struct {
	RowID   int
	ACLID   string
	UserID  string
	Details string
}

// CheckRoleACL checks whether the role is granted the permission
func (dao DAO) CheckRoleACL(roleID string, aclID string) (bool, error) {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(`
		SELECT COUNT(1) 
		FROM role_acl
		WHERE role_id=$1 AND acl_id=$2
	`, roleID, aclID)

	var result int
	err = row.Scan(&result)
	return result > 0, err
}

// CheckMultiRoleACL checks whether any of the roles are granted the permission
func (dao DAO) CheckMultiRoleACL(roleIDs []string, aclID string) (bool, error) {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	queryParams := []interface{}{aclID}
	inValueList := []string{}
	for i, roleID := range roleIDs {
		queryParams = append(queryParams, roleID)
		inValueList = append(inValueList, fmt.Sprintf("$%d", i+2))
	}
	inValue := "(" + strings.Join(inValueList, ", ") + ")"
	query := fmt.Sprintf(`
		SELECT COUNT(1) 
		FROM role_acl
		WHERE acl_id=$1 AND role_id IN %s
	`, inValue)

	row := tx.QueryRow(query, queryParams...)

	var result int
	err = row.Scan(&result)
	return result > 0, err
}

// GrantRoleACL grants a permission to a role
func (dao DAO) GrantRoleACL(roleID string, aclID string, details string) error {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO role_acl (role_id, acl_id, details)
		VALUES ($1, $2, $3)
	`, roleID, aclID, details)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RevokeRoleACL revokes a permission from a role
func (dao DAO) RevokeRoleACL(roleID string, aclID string) error {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		DELETE FROM role_acl
		WHERE role_id=$1 AND acl_id=$2
	`, roleID, aclID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
