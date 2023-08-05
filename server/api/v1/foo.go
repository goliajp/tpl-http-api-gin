package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/ginx"
	"github.com/goliajp/http-api-gin/core"
	"github.com/goliajp/http-api-gin/data"
	"github.com/goliajp/http-api-gin/utils/tpx"
	log "github.com/sirupsen/logrus"
)

func (fooCtrl) GetById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// get foo by id
	foo, err := core.GetFooById(ctx, rd, id)
	if err != nil {
		return err
	}
	return ctx.Success(foo)
}

func (fooCtrl) List(ctx *ginx.Context) error {
	log.Info("foo list")
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

	// get foo list
	fooList, total, err := core.ListFoo(ctx, rd, params.Keyword, params.Page, params.PageSize)
	if err != nil {
		return err
	}
	return ctx.Success(fooList, gin.H{"total": total})
}

func (fooCtrl) Create(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	var params core.CreateFooParams
	if err := ctx.BindJSON(&params); err != nil {
		return err
	}

	// create foo
	foo, err := core.CreateFoo(ctx, rd, &params)
	if err != nil {
		return err
	}
	return ctx.Success(foo)
}

func (fooCtrl) UpdateById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")
	var kv struct {
		Updates []tpx.Kv `json:"updates"`
	}
	if err := ctx.BindJSON(&kv); err != nil {
		return err
	}

	// update foo
	foo, err := core.UpdateFooById(ctx, rd, id, kv.Updates)
	if err != nil {
		return err
	}
	return ctx.Success(foo)
}

func (fooCtrl) DeleteById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// delete foo
	if err := core.DeleteFooById(ctx, rd, id); err != nil {
		return err
	}
	return ctx.Success(nil)
}
