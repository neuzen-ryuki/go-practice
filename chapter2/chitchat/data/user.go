package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	Id        int
	UserId    int
	Uuid      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	uuid := createUUID()
	// Create a new session
	statement := "INSERT INTO sessions (user_id, uuid, email) VALUES (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil { return }
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, uuid, user.Email)
	if err != nil { return }
	defer stmt.Close()

	// scan the returned id into the User struct
	statement = "SELECT id, user_id, uuid, email, created_at FROM sessions WHERE uuid = ?"
	stmt, err = Db.Prepare(statement)
	if err != nil { return }
	defer stmt.Close()
	err = stmt.QueryRow(uuid).Scan(&session.Id, &session.UserId, &session.Uuid, &session.Email, &session.CreatedAt)
	if err != nil { return }
	defer stmt.Close()

	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "DELETE FROM sessions WHERE uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "DELETE FROM sessions"
	_, err = Db.Exec(statement)
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// create a new user
	statement := "INSERT INTO users (uuid, name, email, password) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil { return }
	defer stmt.Close()

	uuid := createUUID()

	_, err = stmt.Exec(uuid, user.Name, user.Email, Encrypt(user.Password))
	if err != nil { return }
	defer stmt.Close()

	// scan the returned id into the User struct
	statement = "SELECT id, uuid, created_at FROM users WHERE uuid = ?"
	stmt, err = Db.Prepare(statement)
	if err != nil { return }
	defer stmt.Close()

	err = stmt.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	if err != nil { return }
	defer stmt.Close()

	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "UPDATE users SET name = $2, email = $3 WHERE id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "DELETE FROM users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
