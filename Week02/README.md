#作业
##问题：

 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，
 抛给上层。为什么，应该怎么做请写出代码？
 

 
 ##回答：
 
 
``` golang

//Service 
package service

import (
	"Go-000/Week02/dao"
	"log"
)
func service() {
	id := 110
	userinfo, err := dao.Dao(id)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Printf("query userinfo no rows:%+v\n",err)
		}
		log.Printf("query userinfo detail failed: %+v\n", err)
	}

	log.Printf("userinfo: %+v\n", userinfo)
}


//DAO
package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var (
	dbw *Db
)
type UserInfo struct {
	Id int
	Name string
	Age int
}


type Db struct {
	Dsn      string
	Db       *sql.DB
}


func init() {
	dbw= &Db{
		Dsn:"root:123456@tcp(localhost:3306)/mysql?charset=utf8",
	}
}


func NewDb(dbw *Db) (*Db,error){
	dd,err:=sql.Open("mysql",dbw.Dsn)
	dbw.Db=dd
	if err != nil {
		return &Db{},errors.Wrap(err,"Db open error!")

	}
	return dbw,nil
}

func Dao(id int)(*UserInfo,error){
	var uinfo UserInfo
	db,err:=NewDb(dbw)
	if err != nil {
		return nil,err
	}

	err = db.Db.QueryRow("select id,name,age from userinfo where id =? limit 1",id).Scan(&uinfo.Id,&uinfo.Name,&uinfo.Age)
	defer db.Db.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Query userinfo error")
	}


	return &uinfo,nil
}

 
```



