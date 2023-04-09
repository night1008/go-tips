package device_info

import (
	"git.sofunny.io/data-analysis/device_info/internal/database"
)

type DeviceInfoService struct {
	db *database.DB
}

func New(db *database.DB) *DeviceInfoService {
	return &DeviceInfoService{
		db: db,
	}
}

func (s *DeviceInfoService) Run() error {
	return nil
}
