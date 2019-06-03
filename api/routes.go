package api

import (
	jwt "github.com/appleboy/gin-jwt"
)

// Routes : ...
func (s *Server) Routes(authMw *jwt.GinJWTMiddleware) {
	s.g.GET("/", s.DefaultWelcome)
	api := s.g.Group("/api")
	{
		api.GET("/", s.Welcome)
		api.POST("/system/subscribe", s.Subscribe)

		// auth API group
		auth := api.Group("/auth")
		auth.POST("/register", s.Register)
		auth.POST("/login", authMw.LoginHandler)
		auth.Use(authMw.MiddlewareFunc())
		{

		}
	}
}
