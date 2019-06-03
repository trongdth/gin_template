package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/trongdth/mroom_backend/models"
	"github.com/trongdth/mroom_backend/serializers"
	"github.com/trongdth/mroom_backend/services"
)

// Authenticate : login by email and password
func (s *Server) Authenticate(c *gin.Context) (*models.User, error) {
	var req serializers.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	user, err := s.userSvc.Authenticate(req.Email, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "u.svc.Authenticate")
	}

	return user, nil
}

// Register : ...
func (s *Server) Register(c *gin.Context) {
	var req serializers.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: services.ErrInvalidArgument})
		return
	}

	user, err := s.userSvc.Register(req.FirstName, req.LastName,
		req.Email, req.Password,
		req.ConfirmPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: user, Error: nil})
}
