package memes

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/fsufitch/discord-boar-bot/db/connection"
)

// DAO is a database abstraction around the meme feature set
type DAO struct {
	dbConn *connection.DatabaseConnection
}

// Meme encapsulates the data about an individual meme
type Meme struct {
	ID    int
	URLs  []MemeURL
	Names []MemeName
}

// MemeURL encapsulates the data about a specific URL for a meme
type MemeURL struct {
	ID        int
	URL       string
	Timestamp time.Time
	Author    string
}

// MemeName encapsulates the data about the name for a meme
type MemeName struct {
	ID        int
	Name      string
	Timestamp time.Time
	Author    string
}

// SearchByName finds a meme given the filename prefix
func (dao DAO) SearchByName(name string) (*Meme, error) {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		SELECT m.id, m_n.id, m_n.name, m_n.timestamp, m_n.author, m_u.id, m_u.url, m_u.timestamp, m_u.author
		FROM meme_names m_n
		INNER JOIN memes m ON m_n.meme_id=m.id
		INNER JOIN meme_urls m_u ON m.id=m_u.meme_id
		WHERE m_n.name=$1
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(strings.ToLower(name))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	meme := Meme{
		Names: []MemeName{},
		URLs:  []MemeURL{},
	}
	var (
		mID         int
		mnID        int
		mnName      string
		mnTimestamp time.Time
		mnAuthor    string
		muID        int
		muURL       string
		muTimestamp time.Time
		muAuthor    string
	)
	for rows.Next() {
		if err = rows.Scan(&mID, &mnID, &mnName, &mnTimestamp, &mnAuthor, &muID, &muURL, &muTimestamp, &muAuthor); err != nil {
			return nil, err
		}

		if meme.ID > 0 {
			if meme.ID != mID {
				return nil, fmt.Errorf("Multiple IDs found for search %s: %d, %d", name, meme.ID, mID)
			}
		} else {
			meme.ID = mID
		}

		if mnID > 0 {
			meme.Names = append(meme.Names, MemeName{
				ID:        mnID,
				Name:      mnName,
				Timestamp: mnTimestamp,
				Author:    mnAuthor,
			})
		}
		if muID > 0 {
			meme.URLs = append(meme.URLs, MemeURL{
				ID:        muID,
				URL:       muURL,
				Timestamp: muTimestamp,
				Author:    muAuthor,
			})
		}
	}
	if meme.ID == 0 {
		return nil, nil
	}

	return &meme, nil
}

// AddName adds a name alias to a meme
func (dao DAO) AddName(memeID int, name string, author string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO meme_names (name, timestamp, author, meme_id)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(name, time.Now(), author, memeID); err != nil {
		return err
	}

	return tx.Commit()
}

// AddURL adds a url to a meme
func (dao DAO) AddURL(memeID int, url string, author string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO meme_urls (url, timestamp, author, meme_id)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(url, time.Now(), author, memeID); err != nil {
		return err
	}

	return tx.Commit()
}

// Add creates a new meme with the given name and URL
func (dao DAO) Add(name string, url string, author string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var memeID int

	if stmt, err := tx.Prepare(`
		INSERT INTO memes DEFAULT VALUES RETURNING id
	`); err != nil {
		return err
	} else if err := stmt.QueryRow().Scan(&memeID); err != nil {
		return err
	}

	timestamp := time.Now()

	if stmt, err := tx.Prepare(`
		INSERT INTO meme_names(name, timestamp, author, meme_id)
		VALUES($1, $2, $3, $4)
	`); err != nil {
		return err
	} else if _, err := stmt.Exec(name, timestamp, author, memeID); err != nil {
		return err
	}

	if stmt, err := tx.Prepare(`
		INSERT INTO meme_urls(url, timestamp, author, meme_id)
		VALUES($1, $2, $3, $4)
	`); err != nil {
		return err
	} else if _, err := stmt.Exec(url, timestamp, author, memeID); err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteName deletes a meme name
func (dao DAO) DeleteName(memeID int, name string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if stmt, err := tx.Prepare(`
		DELETE FROM meme_names WHERE meme_id=$1 AND name=$2
	`); err != nil {
		return err
	} else if _, err := stmt.Exec(memeID, name); err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteURL deletes a meme URL
func (dao DAO) DeleteURL(memeID int, url string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if stmt, err := tx.Prepare(`
		DELETE FROM meme_urls WHERE meme_id=$1 AND url=$2
	`); err != nil {
		return err
	} else if _, err := stmt.Exec(memeID, url); err != nil {
		return err
	}

	return tx.Commit()
}

// Delete deletes a meme
func (dao DAO) Delete(memeID int) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if stmt, err := tx.Prepare(`
		DELETE FROM memes WHERE id=$1
	`); err != nil {
		return err
	} else if _, err := stmt.Exec(memeID); err != nil {
		return err
	}

	return tx.Commit()
}

// NewMemeDAO creates a new MemeDAO
func NewMemeDAO(dbConn *connection.DatabaseConnection) *DAO {
	return &DAO{dbConn: dbConn}
}
