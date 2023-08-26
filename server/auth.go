package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/server/middleware"
)

func (s *Server) LoginGet() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	}
}

func (s *Server) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		store, _ := middleware.Store.Get(ctx.Request, "user_store")
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		user, err := s.userService.GetUser(username, password)
		if err != nil {
			ctx.JSON(404, gin.H{"message": "user not found"})
			ctx.Abort()
			return
		}
		store.Values["username"] = username
		store.Values["role"] = user.Role
		store.Values["authenticated"] = true
		store.Save(ctx.Request, ctx.Writer)
		ctx.Redirect(302, "/admin")
	}
}

func (s *Server) Logout() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		store, _ := middleware.Store.Get(ctx.Request, "user_store")
		store.Values["username"] = nil
		store.Values["role"] = nil
		store.Values["authenticated"] = false
		store.Save(ctx.Request, ctx.Writer)
		ctx.Redirect(302, "/")
	}
}
