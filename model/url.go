package model

import "github.com/jinzhu/gorm"

type Url struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

type UrlRepo interface {
	Save(url Url) (int, error)
	Find(id int) (Url, error)
}

type SQLUrlRepo struct {
	DB *gorm.DB
}

func NewSQLUrlRepo(db *gorm.DB) SQLUrlRepo {
	return SQLUrlRepo{
		DB: db,
	}
}

func (s SQLUrlRepo) Save(url Url) (int, error) {
	err := s.DB.Create(&url).Error
	return url.ID, err
}

func (s SQLUrlRepo) Find(id int) (Url, error) {
	var url Url
	err := s.DB.Where("id = ?", id).Take(&url).Error
	return url, err
}
