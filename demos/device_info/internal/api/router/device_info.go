package router

import (
	"git.sofunny.io/data-analysis/device_info/internal/api/controller"
	"git.sofunny.io/data-analysis/device_info/internal/database"
	"git.sofunny.io/data-analysis/device_info/internal/service/device_info"
	"github.com/gin-gonic/gin"
)

func NewDeviceInfoRouter(group *gin.RouterGroup, db *database.DB) {
	svc := device_info.New(db)
	ctrl := controller.NewDeviceInfoController(svc)

	deviceInfoGroup := group.Group("device-infos")
	deviceInfoGroup.GET("", ctrl.List)
	deviceInfoGroup.POST("", ctrl.Create)
}
