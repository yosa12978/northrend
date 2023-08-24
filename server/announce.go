package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/domain"
)

func (s *Server) CreateAnnounce() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		content := ctx.PostForm("content")
		//expires := ctx.PostForm("expires")
		announce := domain.NewAnnounce(content, 0)
		s.announceService.Create(announce)
		ctx.Redirect(302, "/")
	}
}

func (s *Server) RemoveAnnounce() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		s.announceService.Remove()
		ctx.Redirect(302, "/")
	}
}
