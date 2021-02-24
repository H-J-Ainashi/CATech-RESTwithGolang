package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Find(out interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
}
