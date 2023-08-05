package rdm

import "github.com/goliajp/gormx"

type UserGender = string

const (
	UserGenderUndefined UserGender = ""
	UserGenderMale      UserGender = "male"
	UserGenderFemale    UserGender = "female"
)

type User struct {
	gormx.Model

	// attr
	Name     string     `json:"name"`
	Password string     `json:"-"`
	Email    string     `json:"email"`
	Gender   UserGender `json:"gender"`

	// related

	// relation
	Foos []Foo `json:"foos,omitempty"`
	Bars []Bar `json:"bars,omitempty"`
}
