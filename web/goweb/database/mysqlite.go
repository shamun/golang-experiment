package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./mysqlite")
    checkErr(err)

    stmt, err := db.Prepare("insert into userinfo(username, departname, created) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec("kekeke", "Research", "2013-01-01")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("kekeke", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    rows, err := db.Query("select * from userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("%d/%s/%s/%s", uid, username, department, created)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
