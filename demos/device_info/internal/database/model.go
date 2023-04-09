package database

type Model struct {
	ID        uint64 `gorm:"primarykey"`           // 数据库ID
	CreatedAt int64  `gorm:"autoCreateTime:milli"` // 记录创建时间，单位 milli
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"` // 记录更新时间，单位 milli
}
