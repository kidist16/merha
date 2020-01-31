package repository

import (
	"github.com/kidist16/sport-beting/entity"
	"github.com/kidist16/sport-betting/login"
	"github.com/jinzhu/gorm"
)

// LoginGormRepo implements the login.LoginRepository interface
type LoginGormRepo struct {
	conn *gorm.DB
}

// NewLoginGormRepo will create a new LoginGormRepo object
func NewLoginGormRepo(db *gorm.DB) login.LoginRepository {
	return &LoginGormRepo{conn: db}
}

// Logins returns all logins stored in the database
func (lRepo *LoginGormRepo) Logins() ([]entity.Login, []error) {
	lgns := []entity.Login{}
	errs := lRepo.conn.Find(&lgns).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lgns, errs
}

// Login retrieves a login by its id from the database
func (lRepo *LoginGormRepo) Login(id uint) (*entity.Login, []error) {
	lgn := entity.Login{}
	errs := lRepo.conn.First(&lgn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &lgn, errs
}

// UpdateLogin updates a given login in the database
func (lRepo *LoginGormRepo) UpdateLogin(login *entity.Login) (*entity.Login, []error) {
	lgn := login
	errs := lRepo.conn.Save(lgn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// DeleteLogin deletes a given login from the database
func (lRepo *LoginGormRepo) DeleteLogin(id uint) (*entity.Login, []error) {
	lgn, errs := lRepo.Login(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = lRepo.conn.Delete(lgn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// StoreLogin stores a given login in the database
func (lRepo *LoginGormRepo) StoreLogin(login *entity.Login) (*entity.Login, []error) {
	lgn := login
	errs := lRepo.conn.Create(lgn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lgn, errs
}

// LoginByUsername retrieves a login by its username from the database
func (lRepo *LoginGormRepo) LoginByUsername(username string) (*entity.Login, []error) {
	lgn := entity.Login{}
	errs := lRepo.conn.Where("username = ?", username).Find(&lgn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &lgn, errs
}

// UsernameExists check if a given usename is found
func (lRepo *LoginGormRepo) UsernameExists(Username string) bool {
	login := entity.Login{}
	errs := lRepo.conn.Find(&login, "username=?", username).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// LoginByUserId retrieves a login by its user_id from the database
func (lRepo *LoginGormRepo) LoginByUserId(usrId uint) (*entity.Login, []error) {
	lgn := entity.Login{}
	errs := lRepo.conn.Where(" user_id = ?", usrId).Find(&lgn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &lgn, errs
}
