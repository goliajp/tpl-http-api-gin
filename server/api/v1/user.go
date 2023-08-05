package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/ginx"
	"github.com/goliajp/http-api-gin/core"
	"github.com/goliajp/http-api-gin/data"
	"github.com/goliajp/http-api-gin/utils/tpx"
)

func (userCtrl) GetById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// get user by id
	user, err := core.GetUserById(ctx, rd, id)
	if err != nil {
		return err
	}
	return ctx.Success(user)
}

func (userCtrl) List(ctx *ginx.Context) error {
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

	// get user list
	userList, total, err := core.ListUser(ctx, rd, params.Keyword, params.Page, params.PageSize)
	if err != nil {
		return err
	}
	return ctx.Success(userList, gin.H{"total": total})
}

func (userCtrl) Create(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	var params core.CreateUserParams
	if err := ctx.BindJSON(&params); err != nil {
		return err
	}

	// create user
	user, err := core.CreateUser(ctx, rd, &params)
	if err != nil {
		return err
	}
	return ctx.Success(user)
}

func (userCtrl) UpdateById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")
	var kv struct {
		Updates []tpx.Kv `json:"updates"`
	}
	if err := ctx.BindJSON(&kv); err != nil {
		return err
	}

	// update user
	user, err := core.UpdateUserById(ctx, rd, id, kv.Updates)
	if err != nil {
		return err
	}
	return ctx.Success(user)
}

func (userCtrl) DeleteById(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)

	// get params
	id := ctx.ParamInt("id")

	// delete user
	if err := core.DeleteUserById(ctx, rd, id); err != nil {
		return err
	}
	return ctx.Success(nil)
}

func (userCtrl) Login(ctx *ginx.Context) error {
	rd := data.GetPg(ctx)
	kv := data.GetRedis()

	// get params
	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&params); err != nil {
		return err
	}
	user, token, expires, err := core.UserLogin(ctx, rd, kv, params.Email, params.Password)
	if err != nil {
		return err
	}
	return ctx.Success(user, gin.H{"token": token, "expires": expires})
}

func (userCtrl) Logout(ctx *ginx.Context) error {
	kv := data.GetRedis()

	// get params
	token := ctx.GetString("token")

	// logout
	if err := core.UserLogout(ctx, kv, token); err != nil {
		return err
	}
	return ctx.Success(nil)
}
