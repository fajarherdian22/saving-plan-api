package main

import (
	"log"

	"github.com/fajarherdian22/saving-plan-api/db"
	"github.com/fajarherdian22/saving-plan-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config :", err)
	}
	db.ConDB(config.DBDriver, config.DBSource)
}
