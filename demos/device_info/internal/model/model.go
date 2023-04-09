package model

import "git.sofunny.io/data-analysis/device_info/internal/database"

func Migrate(db *database.DB) error {
	return db.AutoMigrate(&DeviceInfo{})
}
