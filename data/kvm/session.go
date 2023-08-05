package kvm

import (
	"time"
)

type Session struct {
	Token       string        `json:"-"`
	UserId      int           `json:"userId"`
	Payload     *string       `json:"payload,omitempty"`
	RefreshedAt *time.Time    `json:"refreshedAt,omitempty"`
	Expires     time.Duration `json:"expires"`
}
