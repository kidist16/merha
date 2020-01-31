package user

import "github.com/kidist16/sport-betting/entity"

// UserService specifies user services
type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreLogin(user *entity.Login) (*entity.Login, []error)
	EmailExists(email string) bool
	UserRoles(user *entity.User) ([]entity.Role, []error)
}

type LoginService interface {
	Logins() ([]entity.Login, []error)
	Login(id uint) (*entity.Login, []error)
	UpdateLogin(user *entity.Login) (*entity.Login, []error)
	DeleteLogin(id uint) (*entity.Login, []error)
	StoreLogin(login *entity.Login) (*entity.Login, []error)
	LoginByUsername(username string) (*entity.Login, []error)
	LoginByUserId(usrId uint) (*entity.Login, []error)
	UsernameExists(username string) bool
}

// RoleService speifies application user role related services
type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}

