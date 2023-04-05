package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/util"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FullName string    `gorm:"not null" json:"full_name" validate:"required-Full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Password string    `gorm:"not null" json:"password" validate:"required-Password is required,MinStringLength(6)-Password has to have a minimum length of 6 characters"`
	Products []Product `json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = util.HashPass(u.Password)
	err = nil
	return
}