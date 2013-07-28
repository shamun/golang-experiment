package main

import (
    "fmt"
    "time"
    "database/sql"
    "github.com/astaxie/beedb"
    _ "github.com/mattn/go-sqlite3"
)

var orm beedb.Model
type Userinfo struct {
    Uid int `beedb:"PK"`
    Username string
    Departname string
    Created string
}

func main() {
    db, err := sql.Open("sqlite3", "./mysqlite")
    if err != nil {
        panic(err)
    }
    orm = beedb.New(db)

    var saveone Userinfo
    saveone.Username = "Test User"
    saveone.Departname = "Test Depart"
    saveone.Created = time.Now().Format("2006-01-01 00:00:00")
    orm.Save(&saveone)
    fmt.Println(saveone)

    var query Userinfo
    orm.Where("uid=?", saveone.Uid).Find(&query)
    fmt.Println(query)
}
