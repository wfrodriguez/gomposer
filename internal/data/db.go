package data

import (
	"database/sql"
	"time"

	"github.com/wfrodriguez/gomposer/cfg"
	"github.com/wfrodriguez/mimir/util"
	_ "modernc.org/sqlite"
)

type SPost map[string]string

var (
	db *sql.DB
)

func NewMemDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite", ":memory:")
	// db, err = sql.Open("sqlite", "/tmp/test.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreatePostsTable(db *sql.DB) error {
	_, err := db.Exec(cfg.SQLCreateTable)
	return err
}

func SaveTag(db *sql.DB, tag string) error {
	_, err := db.Exec("insert or ignore into tag(tag) values(?)", tag)
	return err
}

func GetTags(db *sql.DB) ([]string, error) {
	rows, err := db.Query("select tag from tag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []string
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func GetPosts(db *sql.DB) ([]SPost, error) {
	rows, err := db.Query(
		"select p.title, p.slug, iif(p.desc = '', '-Sin descripción-', p.desc) desc, strftime('%Y-%m-%d', p.date, " +
			"'unixepoch') as fecha from post p order by p.date desc",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []SPost
	for rows.Next() {
		title := ""
		slug := ""
		date := ""
		desc := ""
		err = rows.Scan(&title, &slug, &desc, &date)
		if err != nil {
			return nil, err
		}
		posts = append(posts, SPost{
			"title": title,
			"slug":  slug,
			"date":  date,
			"desc":  desc,
		})
	}

	return posts, nil
}

func GetPostsByTag(db *sql.DB, tag string) ([]SPost, error) {
	rows, err := db.Query(
		"select p.title, p.slug, iif(p.desc = '', '-Sin descripción-', p.desc) desc, "+
			"strftime('%Y-%m-%d', p.date, 'unixepoch') as pdate from post p, post_tag pt where pt.post = p.id and "+
			"pt.tag = ? order by p.date desc",
		tag,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []SPost
	for rows.Next() {
		title := ""
		slug := ""
		desc := ""
		var pdate sql.NullString
		err = rows.Scan(&title, &slug, &desc, &pdate)
		if err != nil {
			return nil, err
		}
		posts = append(posts, SPost{
			"title": title,
			"slug":  slug,
			"desc":  desc,
			"date":  util.TernaryIf(pdate.Valid, pdate.String, "-Sin fecha-"),
		})
	}

	return posts, nil
}

func SavePost(db *sql.DB, title, slug, desc string, date int64, tags []string) error {
	id := 0
	if date == 0 {
		date = time.Now().Unix()
	}
	err := db.QueryRow(
		"insert into post(title, slug, date, desc) values(?, ?, ?, ?) returning id",
		title,
		slug,
		date,
		desc,
	).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		for _, tag := range tags {
			_, err = db.Exec("insert or ignore into post_tag(post, tag) values(?, ?)", id, tag)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
