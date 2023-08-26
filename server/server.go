package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/server/middleware"
	"github.com/yosa12978/northrend/services"
)

type Server struct {
	router          *gin.Engine
	postService     services.PostService
	linkService     services.LinkService
	announceService services.AnnounceService
	userService     services.UserService
	commentService  services.CommentService
}

func NewServer() *Server {
	s := new(Server)
	s.linkService = services.NewLinkService()
	s.postService = services.NewPostService()
	s.announceService = services.NewAnnounceService()
	s.userService = services.NewUserService()
	s.commentService = services.NewCommentService()
	s.userService.Seed()
	return s
}

func (s *Server) SetupRouter() {
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.New()
	s.router.Use(gin.Recovery())
	s.router.Use(middleware.Logger())
	s.router.Static("/assets", "./static")
	s.router.LoadHTMLGlob("templates/*")

	s.router.GET("/", s.Home())
	s.router.GET("/post/:id", s.GetPost())
	s.router.GET("/portal", s.Portal())

	s.router.GET("/login", s.LoginGet())
	s.router.POST("/login", s.Login())

	s.router.POST("/comment/:postId", s.AddComment())

	sec := s.router.Group("/")
	sec.Use(middleware.Admin())
	sec.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(200, "admin.html", nil)
	})
	sec.POST("/post", s.CreatePost())
	sec.POST("/link", s.CreateLink())
	sec.POST("/announce", s.CreateAnnounce())
	sec.POST("/announce/remove", s.RemoveAnnounce())
	sec.GET("/logout", s.Logout())

	sec.POST("/deletepost", s.DeletePost())
	sec.POST("/deletelink", s.DeleteLink())
}

func (server *Server) Listen(listener net.Listener) {
	server.SetupRouter()
	server.router.RunListener(listener)
}

func (s *Server) Home() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"links":    s.linkService.GetLinks(),
			"posts":    s.postService.GetPosts(),
			"announce": s.announceService.Get(),
		})
	}
}
