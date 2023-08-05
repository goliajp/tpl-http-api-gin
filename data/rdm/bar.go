package rdm

import "github.com/goliajp/gormx"

type Bar struct {
	gormx.Model

	// attr
	Name        string `json:"name"`
	Description string `json:"description"`

	// related
	UserId int `json:"userId" gorm:"index"`

	// relation
	User *User `json:"user,omitempty"`
}
