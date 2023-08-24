package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/domain"
)

func (s *Server) CreateLink() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		uri := ctx.PostForm("uri")
		link := domain.NewLink(name, uri)
		s.linkService.CreateLink(link)
		ctx.Redirect(302, "/")
	}
}

func (s *Server) Portal() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uri, exists := ctx.GetQuery("uri")
		if !exists {
			ctx.JSON(404, gin.H{
				"message": "not found",
			})
			ctx.Abort()
			return
		}
		fmt.Printf("redirecting to %s\n", uri)
		ctx.Redirect(http.StatusFound, uri)
	}
}
