package model

type Base struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt int  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int  `gorm:"column:deleted_at" json:"deleted_at"`
}
