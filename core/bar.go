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

type Bar rdm.Bar

func preloadBar(rd *gorm.DB) *gorm.DB {
	return rd
}

func GetBarById(_ context.Context, rd *gorm.DB, id int) (*Bar, error) {
	var bar Bar
	if err := rd.Where("id = ?", id).
		Scopes(preloadBar).
		First(&bar).Error; err != nil {
		return nil, err
	}
	return &bar, nil
}

func ListBar(_ context.Context, rd *gorm.DB, keyword *string, page, pageSize int) ([]Bar, int, error) {
	var list []Bar
	var total int64

	// copy rd for count and filter
	rdCopy := rd.Model(&Bar{})
	if err := rdCopy.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// filter by keyword
	if keyword != nil {
		rdCopy = rdCopy.Where("name LIKE ?", "%"+*keyword+"%")
	}

	// ind bar list
	if err := rdCopy.
		Scopes(gormx.ScopeOrderByCreatedAtDesc, gormx.ScopePagination(page, pageSize), preloadBar).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, int(total), nil
}

type CreateBarParams struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	UserId      int    `json:"userId" binding:"required"`
}

func CreateBar(ctx context.Context, rd *gorm.DB, params *CreateBarParams) (*Bar, error) {

	// build new bar
	bar := &Bar{
		Name:        params.Name,
		Description: params.Description,
		UserId:      params.UserId,
	}

	// create bar
	if err := rd.Create(bar).Error; err != nil {
		return nil, err
	}
	return GetBarById(ctx, rd, bar.Id)
}

func UpdateBarById(ctx context.Context, rd *gorm.DB, id int, updates []tpx.Kv) (*Bar, error) {
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

		// convert userId to int
		if k == "userId" {
			v, err = strconv.Atoi(v.(string))
			if err != nil {
				return nil, err
			}
		}
		payload[k] = v
	}

	// update user
	if err := rd.Model(&Bar{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		return nil, err
	}
	return GetBarById(ctx, rd, id)
}

func DeleteBarById(_ context.Context, rd *gorm.DB, id int) error {
	return rd.Delete(&Bar{}, id).Error
}
