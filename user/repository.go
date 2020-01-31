package user

import (
"github.com/kidist16/sport-beting/entity"

)

// UserRepository specifies  user database operations
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	UserByFullNameAndPassport(fullName string, passport string) (*entity.User, []error)
	PassportExists(passport string) bool
	EmailExists(email string) bool
	UserRoles(user *entity.User) ([]entity.Role, []error)
}

// LoginRepository specifies flight user database operations
type LoginRepository interface {
	Logins() ([]entity.Login, []error)
	Login(id uint) (*entity.Login, []error)
	UpdateLogin(user *entity.Login) (*entity.Login, []error)
	DeleteLogin(id uint) (*entity.Login, []error)
	StoreLogin(login *entity.Login) (*entity.Login, []error)
	LoginByUsername(username string) (*entity.Login, []error)
	LoginByUserId(usrId uint) (*entity.Login, []error)
	UsernameExists(username string) bool
}

// RoleRepository speifies application user role related database operations
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}
