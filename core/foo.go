package core

import (
	"context"
	"github.com/goliajp/gormx"
	"github.com/goliajp/http-api-gin/data/rdm"
	"github.com/goliajp/http-api-gin/utils/strx"
	"github.com/goliajp/http-api-gin/utils/tpx"
	"gorm.io/gorm"
	"strconv"
)

type Foo rdm.Foo

func preloadFoo(rd *gorm.DB) *gorm.DB {
	return rd
}

func GetFooById(_ context.Context, rd *gorm.DB, id int) (*Foo, error) {
	var foo Foo
	if err := rd.Where("id = ?", id).
		Scopes(preloadFoo).
		First(&foo).Error; err != nil {
		return nil, err
	}
	return &foo, nil
}

func ListFoo(_ context.Context, rd *gorm.DB, keyword *string, page, pageSize int) ([]Foo, int, error) {
	var list []Foo
	var total int64

	// copy rd for count and filter
	rdCopy := rd.Model(&Foo{})
	if err := rdCopy.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// filter by keyword
	if keyword != nil {
		rdCopy = rdCopy.Where("name LIKE ?", "%"+*keyword+"%")
	}

	// ind foo list
	if err := rdCopy.
		Scopes(gormx.ScopeOrderByCreatedAtDesc, gormx.ScopePagination(page, pageSize), preloadFoo).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, int(total), nil
}

type CreateFooParams struct {
	Count       int    `json:"count"`
	Description string `json:"description"`
	UserId      int    `json:"userId" binding:"required"`
}

func CreateFoo(ctx context.Context, rd *gorm.DB, params *CreateFooParams) (*Foo, error) {

	// build new foo
	foo := &Foo{
		Count:       params.Count,
		Description: params.Description,
		UserId:      params.UserId,
	}

	// create foo
	if err := rd.Create(foo).Error; err != nil {
		return nil, err
	}
	return GetFooById(ctx, rd, foo.Id)
}

func UpdateFooById(ctx context.Context, rd *gorm.DB, id int, updates []tpx.Kv) (*Foo, error) {
	var err error

	// build update payload
	payload := make(map[string]interface{})
	for _, keyValue := range updates {
		k := strx.SnakeCase(keyValue.Key)
		v := keyValue.Value

		// ignore id
		if k == "id" {
			continue
		}

		// convert count, userId to int
		if k == "userId" || k == "count" {
			v, err = strconv.Atoi(v.(string))
			if err != nil {
				return nil, err
			}
		}
		payload[k] = v
	}

	// update user
	if err := rd.Model(&Foo{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		return nil, err
	}
	return GetFooById(ctx, rd, id)
}

func DeleteFooById(_ context.Context, rd *gorm.DB, id int) error {
	return rd.Delete(&Foo{}, id).Error
}
