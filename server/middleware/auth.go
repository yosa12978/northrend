package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

var Store *sessions.CookieStore = sessions.NewCookieStore([]byte("config.Config.Api.CookieSecret"))

func Admin() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		s, err := Store.Get(ctx.Request, "user_store")
		if err != nil {
			ctx.String(403, err.Error())
			ctx.Abort()
			return
		}
		username := s.Values["username"]
		if username == nil {
			ctx.String(403, "forbidden")
			ctx.Abort()
			return
		}
		user, err := repos.NewUserRepo().GetUserByUsername(username.(string))
		if err != nil || user.Role != domain.ROLE_ADMIN {
			ctx.String(403, "wtf bro")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
