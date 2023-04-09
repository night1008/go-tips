package controller

import (
	"net/http"

	"git.sofunny.io/data-analysis/device_info/internal/service/device_info"
	"github.com/gin-gonic/gin"
)

type DeviceInfoController struct {
	deviceInfoSvc *device_info.DeviceInfoService
}

func NewDeviceInfoController(deviceInfoSvc *device_info.DeviceInfoService) *DeviceInfoController {
	return &DeviceInfoController{
		deviceInfoSvc: deviceInfoSvc,
	}
}

func (ctrl *DeviceInfoController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": true})
}

func (ctrl *DeviceInfoController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
