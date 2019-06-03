package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/trongdth/mroom_backend/config"
	"github.com/trongdth/mroom_backend/models"
	"github.com/trongdth/mroom_backend/services"
)

type Server struct {
	g       *gin.Engine
	userSvc *services.User
	config  *config.Config
}

func (s *Server) pagingFromContext(c *gin.Context) (int, int) {
	var (
		pageS  = c.DefaultQuery("page", "1")
		limitS = c.DefaultQuery("limit", "10")
		page   int
		limit  int
		err    error
	)

	page, err = strconv.Atoi(pageS)
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(limitS)
	if err != nil {
		limit = 10
	}

	return page, limit
}

func (s *Server) userFromContext(c *gin.Context) (*models.User, error) {
	user, err := s.currentUserFromContext(c)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) currentUserFromContext(c *gin.Context) (*models.User, error) {
	userIDVal, ok := c.Get(userIDKey)
	if !ok {
		return nil, errors.New("failed to get userIDKey from context")
	}

	userID := userIDVal.(float64)
	user, err := s.userSvc.FindByID(uint(userID))
	if err != nil {
		return nil, errors.Wrap(err, "s.userSvc.FindByID")
	}

	return user, nil
}

// NewServer : userSvc
func NewServer(g *gin.Engine, userSvc *services.User, config *config.Config) *Server {
	return &Server{
		g:       g,
		userSvc: userSvc,
		config:  config,
	}
}
