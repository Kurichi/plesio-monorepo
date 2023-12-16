package main

import (
	"github.com/Kurichi/plesio-monorepo/services/mission/infra"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/database"
)

func main() {
	db := database.New(config.NewDBConfig())
	if err := infra.Migrate(db); err != nil {
		panic(err)
	}
}
