package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trongdth/gin_template/helpers"
	"github.com/trongdth/gin_template/serializers"
	"github.com/trongdth/gin_template/services"
)

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Mroom Software Endpoint")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: "Mroom Software REST API"})
}

// Subscribe : ...
func (s *Server) Subscribe(c *gin.Context) {
	var req serializers.UserSubscribeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: services.ErrInvalidArgument})
		return
	}

	if !helpers.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: services.ErrInvalidEmail})
		return
	}

	us, err := s.userSvc.SaveSubscribeEmail(req.Email, req.Name, req.Company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: us, Error: nil})
}
