package exec

import (
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/joho/godotenv"
)

func Run() {
	loadConf()
	// Run command
}

func loadConf() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}

	logger.InitZap()
}
