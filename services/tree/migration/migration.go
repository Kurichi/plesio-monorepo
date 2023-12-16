package main

import (
	"github.com/Kurichi/plesio-monorepo/services/tree/infra"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/database"
)

func main() {
	db := database.New(config.NewDBConfig())
	if err := infra.Migrate(db); err != nil {
		panic(err)
	}
}
