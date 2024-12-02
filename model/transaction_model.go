package model

import "time"

type Transaction struct {
	Id        int       `gorm:"primaryKey;size:20" json:"id"`
	NIK       string    `gorm:"size:25" json:"nik"`
	Income    float64   `gorm:"type:decimal(18,2);not null" json:"income"`
	Outcome   float64   `gorm:"type:decimal(18,2);not null" json:"outcome"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
