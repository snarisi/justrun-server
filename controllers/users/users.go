package users

import (
  "net/http"
  "fmt"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "justrun-server/db"
)

func FindAll(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  err, users := db.GetAllUsers()
  if err != nil { fmt.Println(err) }

  res, _ := json.Marshal(users)
  fmt.Println(users)
  fmt.Println(string(res))
  rw.Write(res)
}

func Create(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  decoder := json.NewDecoder(r.Body)
  var user db.User
  err := decoder.Decode(&user)
  if err != nil { fmt.Println(err) }
  err = db.CreateUser(user)
  if err != nil { fmt.Println(err) }
  fmt.Println(user)
}

func FindByEmail(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  err, user := db.GetUserByEmail(p.ByName("email"))
  if err != nil { fmt.Println(err) }

  res, _ := json.Marshal(user)
  rw.Write(res)
}
