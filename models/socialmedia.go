package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Nama social media harus diisi"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~URL social media harus diisi"`
	UserID         uint   `gorm:"not null" json:"user_id"`
}

func (p *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return err
}

func (p *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	err = nil
	return err
}
