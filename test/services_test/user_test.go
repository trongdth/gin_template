package services_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/trongdth/mroom_backend/config"
	"github.com/trongdth/mroom_backend/daos"
	"github.com/trongdth/mroom_backend/services"

	_ "github.com/go-sql-driver/mysql"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type UserTestSuite struct {
	suite.Suite
	userDAO *daos.User
	userSrv *services.User
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *UserTestSuite) SetupTest() {
	// load config
	conf := config.GetConfig()

	// init daos
	if err := daos.Init(conf); err != nil {
		panic(err)
	}

	suite.userDAO = daos.NewUser()
	suite.userSrv = services.NewUserService(suite.userDAO, conf)
}

func (suite *UserTestSuite) TestUserSrvInitSuccessfully() {
	suite.NotNil(suite.userSrv)
}
