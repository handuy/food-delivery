package appctx

import "gorm.io/gorm"

type ctx struct {
	mainDB *gorm.DB
	tokenSecret string
}

func New(db *gorm.DB,tokenSecret string ) *ctx {
	return &ctx{mainDB: db, tokenSecret: tokenSecret}
}

func(c *ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}

func(c *ctx) GetTokenSecret() string {
	return c.tokenSecret
}