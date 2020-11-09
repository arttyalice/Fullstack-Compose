package main

import (
	"deeploy-exam/app/config"
	"deeploy-exam/app/database"
	"deeploy-exam/app/routes"
	"flag"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// dbMigrate := flag.Bool("db:migrate", false, "use for db migration")
	flag.Parse()

	cfg := config.ParseConfig()

	database.NewSQLClient(cfg)
	database.MigrateSQLSchema()
	// if dbMigrate != nil && *dbMigrate {
	// 	database.MigrateSQLSchema()
	// 	return
	// }

	app := gin.Default()
	app.Use(cors.Default())
	routes.RegisterRoutes(app.Group("/api"))

	app.Run(":" + cfg.ServePort)
}
