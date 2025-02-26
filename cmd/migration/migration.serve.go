package main

import (
	"flag"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"log"
)

func main() {
	var migrationType string
	// flag to determine migration type
	flag.StringVar(&migrationType, "type", "run", "Type of migration: run or rollback")
	flag.Parse()

	// setup user config
	cfg, err := configs.UserConfigInit()
	if err != nil {
		panic(err)
	}

	// setup postgres connection
	log.Println("Starting create connection  ...")
	migration := configs.NewMigration(cfg.Postgres)
	err = migration.MigrationInit()
	if err != nil {
		panic(err)
	}

	// run or rollback migration
	log.Println("Starting migration ...")
	switch migrationType {
	case "run":
		migration.Run()
	case "rollback":
		migration.Rollback()
	default:
		log.Fatalf("Invalid migration type: %s. Must be 'run' or 'rollback'", migrationType)
	}
	log.Println("Migration done ...")

}
