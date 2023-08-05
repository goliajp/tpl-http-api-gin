package core

import (
	"context"
	"errors"
	"github.com/goliajp/gormx"
	"github.com/goliajp/http-api-gin/data/rdm"
	"github.com/goliajp/http-api-gin/env"
	"github.com/goliajp/http-api-gin/utils/strx"
	"github.com/goliajp/http-api-gin/utils/tpx"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type User rdm.User

func (u *User) CheckPassword(password string) bool {
	return strx.HashPassword(u.Password, env.CryptCost) == password
}

func preloadUser(rd *gorm.DB) *gorm.DB {
	return rd.Preload("Foos").Preload("Bars")
}

func UserLogin(ctx context.Context, rd *gorm.DB, kv *redis.Client, email, password string) (
	*User, string,
	time.Duration, error,
) {
	// get user
	user, err := getUserByEmail(ctx, rd, email)
	if err != nil {
		return nil, "", 0, err
	}

	// check password
	if !user.CheckPassword(password) {
		return nil, "", 0, errors.New("invalid password")
	}

	// create session
	sx, err := CreateSession(ctx, kv, user.Id, nil)
	if err != nil {
		return nil, "", 0, err
	}
	return user, sx.Token, sx.Expires, nil
}

func UserLogout(ctx context.Context, kv *redis.Client, token string) error {
	return DeleteSession(ctx, kv, token)
}

func GetUserById(_ context.Context, rd *gorm.DB, id int) (*User, error) {
	var user User
	if err := rd.Where("id = ?", id).
		Scopes(preloadUser).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ListUser(_ context.Context, rd *gorm.DB, keyword *string, page, size int) ([]User, int, error) {
	var list []User
	var total int64

	// copy rd for count and filter
	rdCopy := rd.Model(&User{})
	if err := rdCopy.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// filter by keyword
	if keyword != nil {
		rdCopy = rdCopy.Where("name LIKE ?", "%"+*keyword+"%")
	}

	// find user list
	if err := rdCopy.
		Scopes(gormx.ScopeOrderByCreatedAtDesc, gormx.ScopePagination(page, size), preloadUser).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, int(total), nil
}

type CreateUserParams struct {
	Name     string         `json:"name" binding:"required"`
	Email    string         `json:"email" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Gender   rdm.UserGender `json:"gender"`
}

func CreateUser(ctx context.Context, rd *gorm.DB, params *CreateUserParams) (*User, error) {

	// check email
	if _, err := getUserByEmail(ctx, rd, params.Email); err == nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, errors.New("email already exists")
		}
	}

	// build new user
	user := &User{
		Name:     params.Name,
		Password: strx.HashPassword(params.Password, env.CryptCost),
		Email:    params.Email,
		Gender:   params.Gender,
	}

	// create user
	if err := rd.Create(user).Error; err != nil {
		return nil, err
	}
	return GetUserById(ctx, rd, user.Id)
}

func UpdateUserById(ctx context.Context, rd *gorm.DB, id int, updates []tpx.Kv) (*User, error) {

	// build update payload
	payload := make(map[string]interface{})
	for _, keyValue := range updates {
		k := strx.SnakeCase(keyValue.Key)
		v := keyValue.Value

		// ignore id, password
		if k == "id" || k == "password" {
			continue
		}
		payload[k] = v
	}

	// update user
	if err := rd.Model(&User{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		return nil, err
	}
	return GetUserById(ctx, rd, id)
}

func DeleteUserById(_ context.Context, rd *gorm.DB, id int) error {
	return rd.Delete(&User{}, id).Error
}

func getUserByEmail(_ context.Context, rd *gorm.DB, email string) (*User, error) {
	var user User
	if err := rd.Where("email = ?", email).
		Scopes(preloadUser).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
