package db

import (
	"database/sql"
	"fmt"

	"github.com/JulzDiverse/eirinidotcf/api/api"
	_ "github.com/go-sql-driver/mysql"
)

type Photobase struct {
	db *sql.DB
}

func NewSQLPhotobase(db *sql.DB) Photobase {
	return Photobase{db: db}
}

func (p *Photobase) List() ([]api.Photo, error) {
	rows, err := p.db.Query("SELECT title, author, hero, data FROM photos ORDER BY timestamp DESC;")
	if err != nil {
		fmt.Println("Failed to get photos", err.Error())
		return []api.Photo{}, err
	}
	defer rows.Close()

	var (
		title, author, hero string
		data                []byte
	)

	photos := []api.Photo{}
	for rows.Next() {
		err = rows.Scan(&title, &author, &hero, &data)
		if err != nil {
			fmt.Println("Failed to get a photo", err.Error())
			continue
		}
		defer rows.Close()
		photos = append(photos, api.Photo{
			Title:  title,
			Hero:   hero,
			Author: author,
			Data:   data,
		})
	}

	return photos, nil
}

func (p *Photobase) Add(photo api.Photo) error {
	statement, err := p.db.Prepare("INSERT INTO photos(title, author, hero, data, timestamp) VALUES (?, ?, ?, ?, NOW())")
	if err != nil {
		fmt.Println("Failed to prepare statement", err.Error())
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(photo.Title, photo.Author, photo.Hero, photo.Data)
	if err != nil {
		fmt.Println("Failed to insert photo", err.Error())
		return err
	}
	return nil
}

func (p *Photobase) Get(id string) (api.Photo, error) {
	statement, err := p.db.Prepare("SELECT title, author, hero, data FROM photos WHERE hero=?;")
	if err != nil {
		fmt.Println("Failed to prepare statement", err.Error())
		return api.Photo{}, nil
	}
	defer statement.Close()

	rows, err := statement.Query(id)
	if err != nil {
		fmt.Println("Failed to get photo", err.Error())
		return api.Photo{}, nil
	}
	rows.Next()
	var (
		title, author, hero string
		data                []byte
	)

	err = rows.Scan(&title, &author, &hero, &data)
	if err != nil {
		fmt.Println("Failed to get a photo", err.Error())
		return api.Photo{}, nil
	}
	defer rows.Close()
	return api.Photo{
		Title:  title,
		Hero:   hero,
		Author: author,
		Data:   data,
	}, nil
}
