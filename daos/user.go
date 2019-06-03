package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/trongdth/mroom_backend/models"
)

// User : struct
type User struct {
}

// NewUser :
func NewUser() *User {
	return &User{}
}

// Create : tx, user
func (u *User) Create(tx *gorm.DB, user *models.User) error {
	return errors.Wrap(tx.Create(user).Error, "tx.Create")
}

// Update : tx, user
func (u *User) Update(tx *gorm.DB, user *models.User) error {
	return errors.Wrap(tx.Save(user).Error, "tx.Save")
}

// FindByEmail : email
func (u *User) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.First")
	}
	return &user, nil
}

// FindByID : id
func (u *User) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.First")
	}
	return &user, nil
}

// FindSubscribedUserByEmail : email
func (u *User) FindSubscribedUserByEmail(email string) (*models.UserSubscribe, error) {
	var user models.UserSubscribe
	if err := db.Table("user_subscribes").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.First")
	}
	return &user, nil
}

// CreateSubscribedUser : tx, user
func (u *User) CreateSubscribedUser(tx *gorm.DB, us *models.UserSubscribe) error {
	return errors.Wrap(tx.Table("user_subscribes").Create(us).Error, "tx.Create")
}

// SaveSubscribedUser : tx, user
func (u *User) SaveSubscribedUser(tx *gorm.DB, us *models.UserSubscribe) error {
	return errors.Wrap(tx.Table("user_subscribes").Save(us).Error, "tx.Save")
}
