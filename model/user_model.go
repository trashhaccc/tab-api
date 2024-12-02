package model

type User struct {
	NIK         string `gorm:"primaryKey;size:25" json:"nik"`
	Username    string `gorm:"unique;not null" json:"username"`
	Password    string `gorm:"size:100;not null" json:"password"`
	FullName    string `gorm:"size:100;not null" json:"full_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `gorm:"type:enum('admin','nasabah');not null" json:"role"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
