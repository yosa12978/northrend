package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/services"
)

type Server struct {
	router          *gin.Engine
	postService     services.PostService
	linkService     services.LinkService
	announceService services.AnnounceService
}

func NewServer() *Server {
	s := new(Server)
	s.linkService = services.NewLinkService()
	s.postService = services.NewPostService()
	s.announceService = services.NewAnnounceService()
	return s
}

func (s *Server) SetupRouter() {
	s.router = gin.New()
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())
	s.router.Static("/assets", "./static")
	s.router.LoadHTMLGlob("templates/*")

	s.router.GET("/", s.Home())
	s.router.GET("/post/:id", s.GetPost())
	s.router.POST("/post", s.CreatePost())
	s.router.POST("/link", s.CreateLink())
	s.router.GET("/portal", s.Portal())
	s.router.POST("/announce", s.CreateAnnounce())
	s.router.DELETE("/announce", s.RemoveAnnounce())
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
