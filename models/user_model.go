package models

import "time"

type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Name     string    `grom:"type:varchar(128);not null" json:"name"`
	Email    string    `gorm:"type:varchar(128);not null;unique" json:"email"`
	Password string    `gorm:"type:varchar(128);not null" json:"password"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
}
