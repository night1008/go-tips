package router

import (
	"git.sofunny.io/data-analysis/device_info/internal/database"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine, db *database.DB) {
	publicRouter := engine.Group("")

	NewDeviceInfoRouter(publicRouter, db)
}
