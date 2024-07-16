package entity

import (
	"time"
)

type User struct {
	ID          uint64
	Name        string
	Username    string
	Password    string
	Role        int32
	Description string
	CreateAt    time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
}
