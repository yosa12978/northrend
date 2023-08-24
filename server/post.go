package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/domain"
)

func (s *Server) GetPost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		post, err := s.postService.GetPost(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"message": "not found"})
			ctx.Abort()
			return
		}
		ctx.HTML(http.StatusOK, "post.html", gin.H{
			"post": post,
		})
	}
}

func (s *Server) CreatePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		content := ctx.PostForm("content")
		post := domain.NewPost(content)
		s.postService.CreatePost(post)
		ctx.Redirect(302, "/")
	}
}
