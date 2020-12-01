package service

import (
	"Go-000/Week02/dao"
	"log"
)

func service() {
	id := 110
	userinfo, err := dao.Dao(id)
	if err != nil {
		log.Printf("query userinfo detail failed: %+v\n", err)
	}

	log.Printf("userinfo: %+v\n", userinfo)
}
