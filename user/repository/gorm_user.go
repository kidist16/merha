package repository

import (
	"github.com//kidist16/sport-betting/entity"
	"github.com/jinzhu/gorm"
	"github.com/kidist16/sport-betting/user"
)

// UserGormRepo implements the user.UserRepository interface
type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo will create a new UserGormRepo object
func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

// Users returns all users stored in the database
func (uRepo *UserGormRepo) Users() ([]entity.User, []error) {
	usrs := []entity.User{}
	errs := uRepo.conn.Find(&usrs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// User retrieves a user by its id from the database
func (uRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	usr := entity.User{}
	errs := uRepo.conn.First(&usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &usr, errs
}

// UpdateUser updates a given user in the database
func (uRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := uRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user from the database
func (uRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := uRepo.User(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = uRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a given user in the database
func (uRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := uRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UserByFullNameAndPassport retrieves a user by its full name and passport from the database
func (uRepo *UserGormRepo) UserByFullNameAndPassport(fullName string, passport string) (*entity.User, []error) {
	usr := entity.User{}
	errs := uRepo.conn.Where("full_name = ? AND passport = ?", fullName, passport).Find(&usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &usr, errs
}

// PassportExists check if a given phone number is found
func (uRepo *UserGormRepo) PassportExists(passport string) bool {
	user := entity.User{}
	errs := uRepo.conn.Find(&user, "passport=?", passport).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// EmailExists check if a given email is found
func (uRepo *UserGormRepo) EmailExists(email string) bool {
	user := entity.User{}
	errs := uRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// UserRoles returns list of application roles that a given user has
func (uRepo *UserGormRepo) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles := []entity.Role{}
	errs := uRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
