package acl

import "database/sql"

// UserACL encapsulates a single permission granted to a user
type UserACL struct {
	RowID   int
	ACLID   string
	UserID  string
	Details string
}

// CheckUserACL checks whether the user is granted the permission
func (dao DAO) CheckUserACL(userID string, aclID string) (bool, error) {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(`
		SELECT COUNT(1) 
		FROM user_acl
		WHERE user_id=$1 AND acl_id=$2
	`, userID, aclID)

	var result int
	err = row.Scan(&result)
	return result > 0, err
}

// GrantUserACL grants a permission to a user
func (dao DAO) GrantUserACL(userID string, aclID string, details string) error {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO user_acl (user_id, acl_id, details)
		VALUES ($1, $2, $3)
	`, userID, aclID, details)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RevokeUserACL revokes a permission from a user
func (dao DAO) RevokeUserACL(userID string, aclID string) error {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		DELETE FROM user_acl
		WHERE user_id=$1 AND acl_id=$2
	`, userID, aclID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
