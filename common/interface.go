package common

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetTokenSecret() string
}