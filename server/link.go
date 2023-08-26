package server

import (
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
		ctx.Redirect(http.StatusFound, uri)
	}
}

func (s *Server) DeleteLink() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.PostForm("linkId")
		_, err := s.linkService.DeleteLink(id)
		if err != nil {
			ctx.String(404, "link not found")
			ctx.Abort()
			return
		}
		ctx.Redirect(302, "/admin")
	}
}
