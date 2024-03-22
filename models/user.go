package models

import (
	"errors"

	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username harus diisi"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email harus diisi,email~Email tidak valid"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password harus diisi,minstringlength(6)~Password minimal 6 karakter"`
	Age          uint8         `gorm:"not null" json:"age" form:"age" valid:"required~Age harus diisi"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Socialmedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	//validate govalidator
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	// validate age
	if u.Age < 8 {
		return errors.New("UMUR TIDAK BOLEH KURANG DARI 8 TAHUN")
	}

	// hash password
	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	// validate age
	if u.Age < 8 {
		return errors.New("UMUR TIDAK BOLEH KURANG DARI 8 TAHUN")
	}

	if u.Email == "" {
		return errors.New("EMAIL TIDAK BOLEH KOSONG")
	}

	if u.Username == "" {
		return errors.New("USERNAME TIDAK BOLEH KOSONG")
	}

	err = nil
	return
}
