package rdm

import "github.com/goliajp/gormx"

type Foo struct {
	gormx.Model

	// attr
	Count       int    `json:"count"`
	Description string `json:"description"`

	// related
	UserId int `json:"userId" gorm:"index"`

	// relation
	User *User `json:"user,omitempty"`
}
