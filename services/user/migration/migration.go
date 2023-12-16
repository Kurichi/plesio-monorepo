package main

import (
	"github.com/Kurichi/plesio-monorepo/services/user/infra"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/database"
)

func main() {
	db := database.New(config.NewDBConfig())
	if err := infra.Migrate(db); err != nil {
		panic(err)
	}
}
