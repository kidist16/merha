package service

import (
		"github.com//kidist16/sport-betting/entity"
		"github.com/kidist16/sport-betting/user"
)

// LoginService implements user.LoginService interface
type LoginService struct {
	loginRepo user.LoginRepository
}

// NewLoginService will create new LoginService object
func NewLoginService(lRepo login.LoginRepository) login.LoginService {
	return &LoginService{loginRepo: lRepo}
}

// Logins returns list of logins
func (lServ *LoginService) Logins() ([]entity.Login, []error) {

	lgns, errs := lServ.loginRepo.Logins()

	if len(errs) > 0 {
		return nil, errs
	}

	return lgns, nil
}

// Login retrieves a login by its id
func (lServ *LoginService) Login(id uint) (*entity.Login, []error) {
	lgn, errs := lServ.loginRepo.Login(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// UpdateLogin updates a given login
func (lServ *LoginService) UpdateLogin(login *entity.Login) (*entity.Login, []error) {
	lgn, errs := lServ.loginRepo.UpdateLogin(login)
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// DeleteLogin deletes a given login
func (lServ *LoginService) DeleteLogin(id uint) (*entity.Login, []error) {
	lgn, errs := lServ.loginRepo.DeleteLogin(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// StoreLogin persists new login information
func (lServ *LoginService) StoreLogin(login *entity.Login) (*entity.Login, []error) {

	lgn, errs := lServ.loginRepo.StoreLogin(login)

	if len(errs) > 0 {
		return nil, errs
	}

	return lgn, nil
}

// LoginByUsername retrieves a login by its username
func (lServ *LoginService) LoginByUsername(username string) (*entity.Login, []error) {
	lgn, errs := lServ.loginRepo.LoginByUsername(username)
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// LoginByUserId retrieves a login by its user id
func (lServ *LoginService) LoginByUserId(usrId string) (*entity.Login, []error) {
	lgn, errs := lServ.loginRepo.LoginByUserId(usrId)
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// UsernameExists check if there is a login with a given username
func (us *LoginService) UsernameExists(username string) bool {
	exists := us.loginRepo.UsernameExists(username)
	return exists
}
