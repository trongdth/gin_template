package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/trongdth/gin_template/config"
	"github.com/trongdth/gin_template/daos"
	"github.com/trongdth/gin_template/models"
	"golang.org/x/crypto/bcrypt"
)

// User : struct
type User struct {
	ud   *daos.User
	conf *config.Config
}

// NewUserService : user dao, config
func NewUserService(ud *daos.User, conf *config.Config) *User {
	return &User{
		ud:   ud,
		conf: conf,
	}
}

func (u *User) validate(firstName, lastName, email, password, confirmPassword string) error {
	if email == "" {
		return ErrInvalidEmail
	}
	if password == "" || confirmPassword == "" {
		return ErrInvalidPassword
	}
	if password != confirmPassword {
		return ErrPasswordMismatch
	}
	return nil
}

// FindByID : user id
func (u *User) FindByID(id uint) (*models.User, error) {
	user, err := u.ud.FindByID(id)
	if err != nil {
		return nil, ErrorWithMessage(ErrSystemError, "u.userDAO.FindByID")
	}

	return user, nil
}

// Authenticate : email, password
func (u *User) Authenticate(email, password string) (*models.User, error) {
	user, err := u.ud.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.userDAO.FindByEmail")
	}
	if user == nil {
		return nil, ErrEmailNotExists
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

// Register : firstName, lastName, email, password, confirmPassword
func (u *User) Register(firstName, lastName,
	email, password, confirmPassword string) (*models.User, error) {

	if err := u.validate(firstName, lastName, email, password, confirmPassword); err != nil {
		return nil, errors.Wrap(err, "u.validate")
	}

	user, err := u.ud.FindByEmail(email)
	if user != nil {
		return nil, ErrEmailAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrorWithMessage(ErrSystemError, "bcrypt.GenerateFromPassword")
	}

	newUser := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  firstName + " " + lastName,
		Email:     email,
		Password:  string(hashed),
	}

	err = daos.WithTransaction(func(tx *gorm.DB) error {
		err = u.ud.Create(tx, newUser)
		if err != nil {
			return ErrorWithMessage(ErrSystemError, err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// SaveSubscribeEmail : email, name, company
func (u *User) SaveSubscribeEmail(email string, name string, company string) (*models.UserSubscribe, error) {

	var errTnx error
	us, _ := u.ud.FindSubscribedUserByEmail(email)
	if us == nil {
		us = &models.UserSubscribe{
			Email:   email,
			Name:    name,
			Company: company,
		}

	} else {
		us.Name = name
		us.Company = company

	}

	errTnx = daos.WithTransaction(func(tx *gorm.DB) error {
		err := u.ud.SaveSubscribedUser(tx, us)
		if err != nil {
			return errors.Wrap(err, "u.UserDAO.SaveSubscribedUser")
		}
		return nil
	})

	if errTnx != nil {
		return nil, errors.Wrap(errTnx, "WithTransaction")
	}

	return us, nil
}
