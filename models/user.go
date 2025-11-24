package models

import "example.com/event-booking/db"

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (user *User) Save() error {
	query := `
	INSERT INTO users
	(name, email, password)
	VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil{
		return err
	}
	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func GetAllUsers()([]User, error){
	query := `
	Select id, name, email from users;
	`
	users := []User{}
	rows, err := db.DB.Query(query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var user User 
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func NewUser()*User{
	return &User{}
}