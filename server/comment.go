package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/domain"
)

func (s *Server) AddComment() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		postId := ctx.Param("postId")
		post, err := s.postService.GetPost(postId)
		if err != nil {
			ctx.String(404, "post not found")
			ctx.Abort()
			return
		}
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		comment := domain.NewComment(email, name, content, post.Id)
		s.commentService.Create(comment)
		ctx.Redirect(302, fmt.Sprintf("/post/%s", postId))
	}
}
