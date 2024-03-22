package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Nama social media harus diisi"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~URL social media harus diisi"`
	UserID         uint   `gorm:"not null" json:"user_id"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return err
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {

	if p.SocialMediaURL == "" {
		return errors.New("URL SOCIAL MEDIA TIDAK BOLEH KOSONG")
	}

	if p.Name == "" {
		return errors.New("NAMA SOCIAL MEDIA TIDAK BOLEH KOSONG")
	}

	err = nil
	return err
}
