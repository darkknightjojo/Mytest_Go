package main

import (
	"answers/pkg/config"
	_ "answers/pkg/database"
	"answers/pkg/logger"
	"answers/routes"
)

var err error

func main() {

	r := routes.RegisterRouters()

	port := config.Config.Port
	logger.Info("http server started, listened on port %v", port)
	r.Run(":" + port)
}
