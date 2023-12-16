package main

import (
	"github.com/Kurichi/plesio-monorepo/services/item/infra"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
)

func main() {
	cfg := config.NewDBConfig()
	db := database.New(cfg)
	if err := infra.Migrate(db); err != nil {
		panic(err)
	}
}
