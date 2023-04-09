package model

import (
	"database/sql"

	"git.sofunny.io/data-analysis/device_info/internal/database"
)

type DeviceInfo struct {
	database.Model

	DeviceModel string `gorm:"uniqueIndex"`

	Name    sql.NullString
	RAM     sql.NullInt64
	CPU     sql.NullInt64
	CPUFreq sql.NullInt64
	GPU     sql.NullString

	CreatorID uint64
}
