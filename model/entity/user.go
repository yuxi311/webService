package entity

import (
	"time"
)

type User struct {
	ID          uint64    `gorm:"primarykey;"`
	Name        string    `gorm:"type:string;size:255;not null;"`
	Username    string    `gorm:"type:string;size:255;not null;unique"`
	Password    string    `gorm:"type:string;size:255;not null;"`
	Role        int32     
	Description string    `gorm:"type:string;size:255;"`
	CreateAt    time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
}
