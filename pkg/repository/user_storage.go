package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tidy/pkg/model"
)

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) UserStorage {
	return &userStorage{
		db: db,
	}
}

func (c *userStorage) Create(m *model.User) (int, error) {
	result, err := c.db.Exec(`INSERT INTO user (name, sername) VALUES (?, ?) RETURNING id`, m.Name, m.Surname)

	if err != nil {
		return 0, errors.New("insert error")
	}
	res, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

func (c *userStorage) ReadOne(id int) (model.User, error) {
	rows, err := c.db.Query(`SELECT user.id, user.name, user.sername
	FROM user WHERE user.id=?`, id)
	var a model.User
	if err != nil {
		log.Println(err)
		return a, err
	}

	for rows.Next() {
		err = rows.Scan(&a.ID, &a.Name, &a.Surname)
		if err != nil {
			log.Println(err)
			return a, err
		}

	}

	if len(a.Name) == 0 {
		return a, errors.New("bad req")
	}

	return a, nil
}

func (c *userStorage) Update(m *model.UpdateU, id int) error {

	_, err := c.db.Exec(`UPDATE user SET name=?, sername=? WHERE user.id = ?`, m.Name, m.Surname, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *userStorage) Delete(id int) error {
	_, err := c.db.Exec(`DELETE FROM user WHERE user.id=?`, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
