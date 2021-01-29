package appctx

import "gorm.io/gorm"

type ctx struct {
	mainDB *gorm.DB
}

func New(db *gorm.DB) *ctx {
	return &ctx{mainDB: db}
}

func(c *ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}