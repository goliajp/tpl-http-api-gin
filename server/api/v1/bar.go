package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/ginx"
	"github.com/goliajp/http-api-gin/core"
	"github.com/goliajp/http-api-gin/data"
	"github.com/goliajp/http-api-gin/utils/tpx"
)

func (barCtrl) GetById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// get bar by id
	bar, err := core.GetBarById(ctx, rd, id)
	if err != nil {
		return err
	}
	return ctx.Success(bar)
}

func (barCtrl) List(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	var params struct {
		Keyword  *string `form:"keyword"`
		Page     int     `form:"page"`
		PageSize int     `form:"pageSize"`
	}
	if err := ctx.BindQuery(&params); err != nil {
		return err
	}

	// get bar list
	barList, total, err := core.ListBar(ctx, rd, params.Keyword, params.Page, params.PageSize)
	if err != nil {
		return err
	}
	return ctx.Success(barList, gin.H{"total": total})
}

func (barCtrl) Create(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	var params core.CreateBarParams
	if err := ctx.BindJSON(&params); err != nil {
		return err
	}

	// create bar
	bar, err := core.CreateBar(ctx, rd, &params)
	if err != nil {
		return err
	}
	return ctx.Success(bar)
}

func (barCtrl) UpdateById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")
	var kv struct {
		Updates []tpx.Kv `json:"updates"`
	}
	if err := ctx.BindJSON(&kv); err != nil {
		return err
	}

	// update bar
	bar, err := core.UpdateBarById(ctx, rd, id, kv.Updates)
	if err != nil {
		return err
	}
	return ctx.Success(bar)
}

func (barCtrl) DeleteById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// delete bar
	if err := core.DeleteBarById(ctx, rd, id); err != nil {
		return err
	}
	return ctx.Success(nil)
}
