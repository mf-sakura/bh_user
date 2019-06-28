package db

import ()

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func InsertUser(user User) (int64, error) {
	q := "INSERT INTO `user` (`name`) VALUES (?)"

	res, err := db.Exec(q, user.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUser(id int64) (*User, error) {

	q := "SELECT * FROM `user` WHERE id = ?"
	var user User
	if err := db.Get(&user, q, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func SelectUserAll() ([]*User, error) {
	q := "SELECT * FROM `user` "
	var hotels []*User
	if err := db.Select(hotels, q); err != nil {
		return nil, err
	}
	return hotels, nil
}
