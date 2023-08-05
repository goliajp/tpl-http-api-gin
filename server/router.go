package server

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/ginx"
	v1 "github.com/goliajp/http-api-gin/server/api/v1"
)

func withRouters(e *gin.Engine) *gin.Engine {
	// bind api to endpoint
	ep := e.Group("/v1")

	// foo endpoint
	{
		ctrl := v1.FooCtrl
		p := ep.Group("/foo")
		a := ep.Group("/foo", v1.AuthUser)
		// public
		p.GET("", ginx.W(ctrl.List))
		p.GET(":id", ginx.W(ctrl.GetById))
		// auth
		a.POST("", ginx.W(ctrl.Create))
		a.PATCH("/:id", ginx.W(ctrl.UpdateById))
		a.DELETE("/:id", ginx.W(ctrl.DeleteById))
	}

	// bar endpoint
	{
		ctrl := v1.BarCtrl
		p := ep.Group("/bar")
		a := ep.Group("/bar", v1.AuthUser)
		// public
		p.GET("", ginx.W(ctrl.List))
		p.GET(":id", ginx.W(ctrl.GetById))
		// auth
		a.POST("", ginx.W(ctrl.Create))
		a.PATCH("/:id", ginx.W(ctrl.UpdateById))
		a.DELETE("/:id", ginx.W(ctrl.DeleteById))
	}

	// user endpoint
	{
		ctrl := v1.UserCtrl
		p := ep.Group("/user")
		a := ep.Group("/user", v1.AuthUser)
		// public
		p.POST("/actions/login", ginx.W(ctrl.Login))
		// auth
		a.GET("", ginx.W(ctrl.List))
		a.POST("", ginx.W(ctrl.Create))
		a.GET("/:id", ginx.W(ctrl.GetById))
		a.PATCH("/:id", ginx.W(ctrl.UpdateById))
		a.DELETE("/:id", ginx.W(ctrl.DeleteById))
		a.POST("/actions/logout", ginx.W(ctrl.Logout))
	}
	return e
}
