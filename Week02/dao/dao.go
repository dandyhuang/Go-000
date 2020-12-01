package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var (
	dbw *Db
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
}

type Db struct {
	Dsn string
	Db  *sql.DB
}

func init() {
	dbw = &Db{
		Dsn: "root:123456@tcp(localhost:3306)/mysql?charset=utf8",
	}
}

func NewDb(dbw *Db) (*Db, error) {
	dd, err := sql.Open("mysql", dbw.Dsn)
	dbw.Db = dd
	if err != nil {
		return &Db{}, errors.Wrap(err, "Db open error!")

	}
	return dbw, nil
}

func Dao(id int) (*UserInfo, error) {
	var uinfo UserInfo
	db, err := NewDb(dbw)
	if err != nil {
		return nil, err
	}

	err = db.Db.QueryRow("select id,name,age from userinfo where id =? limit 1", id).Scan(&uinfo.Id, &uinfo.Name, &uinfo.Age)
	defer db.Db.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Query userinfo error")
	}

	return &uinfo, nil
}
