package service

import (
	"github.com//kidist16/sport-betting/entity"
	"github.com/kidist16/sport-betting/user"
)

// UserService implements user.UserService interface
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService will create new UserService object
func NewUserService(uRepo user.UserRepository) user.UserService {
	return &UserService{userRepo: uRepo}
}

// Users returns list of users
func (uServ *UserService) Users() ([]entity.User, []error) {

	usrs, errs := uServ.userRepo.Users()

	if len(errs) > 0 {
		return nil, errs
	}

	return usrs, nil
}

// User retrieves a user by its id
func (uServ *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := uServ.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateUser updates a given user
func (uServ *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := uServ.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user
func (uServ *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := uServ.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser persists new user information
func (uServ *UserService) StoreUser(user *entity.User) (*entity.User, []error) {

	usr, errs := uServ.userRepo.StoreUser(user)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// UserByFullNameAndPassport retrieves a user by its full name and passport
func (uServ *UserService) UserByFullNameAndPassport(fullName string, passport string) (*entity.User, []error) {
	usr, errs := uServ.userRepo.UserByFullNameAndPassport(fullName, passport)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// PassportExists check if there is a user with a given passport
func (us *UserService) PassportExists(passport string) bool {
	exists := us.userRepo.PassportExists(passport)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (us *UserService) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}

// UserRoles returns list of roles a user has
func (us *UserService) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles, errs := us.userRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
