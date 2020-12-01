package service

import (
	"Go-000/Week02/dao"
	"database/sql"
	"log"
)

func service() {
	id := 110
	userinfo, err := dao.Dao(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("query userinfo no rows:%+v\n", err)
		}
		log.Printf("query userinfo detail failed: %+v\n", err)
	}

	log.Printf("userinfo: %+v\n", userinfo)
}
