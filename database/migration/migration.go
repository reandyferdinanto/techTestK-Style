package migration

import (
	"fmt"
	"log"
	"techTestK-Style/database"
	"techTestK-Style/model/entity"
)

func RunMigration(){
	err := database.DB.AutoMigrate(&entity.Member{} )
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}

