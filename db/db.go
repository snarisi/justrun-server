package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "fmt"
)

type User struct {
  Email string
  Password string
}

var db *sql.DB
var err error

func Connect() {
  db, err = sql.Open("mysql", "sam:password@/justrun?charset=utf8")
  if err != nil {
    fmt.Println("Error connecting to db:")
    log.Fatal(err)
  }
}

func CreateUser(user User) error {
  stmt, err := db.Prepare("INSERT users SET email=?, password=?")
  if err != nil { return err }

  _, err = stmt.Exec(user.Email, user.Password)
  if err != nil { return err }

  return err
}

func GetAllUsers() (error, []User) {
  users := make([]User, 0, 100)
  var user User

  rows, err := db.Query(`
    SELECT email, password FROM users
  `)
  if err != nil { return err, nil }

  defer rows.Close()

  for rows.Next() {
    user = User{}
    rows.Scan(&user.Email, &user.Password)
    users = append(users, user)
  }
  return nil, users
}

func GetUserByEmail(email string) (error, User){
  var user User

  rows, err := db.Query(`
    SELECT * FROM users WHERE email=? LIMIT 1
  `, email)
  if err != nil { log.Fatal(err) }

  defer rows.Close()

  for rows.Next() {
    user = User{}
    rows.Scan(&user.Email, &user.Password)
  }
  return nil, user
}
