package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Judul foto harus diisi"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~URL foto harus diisi"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil

	return err
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {

	if p.Title == "" {
		return errors.New("JUDUL FOTO TIDAK BOLEH KOSONG")
	}

	if p.PhotoURL == "" {
		return errors.New("URL FOTO TIDAK BOLEH KOSONG")
	}

	err = nil

	return err
}
