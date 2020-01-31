package entity

type userPayment struct {
	ID     uint
	userid uint
}

type Ticket struct {
	ID       uint
	Status   bool
	FlightID uint
	UserID   uint
}

type User struct {
	userid    uint
	FName     string `gorm:"type:varchar(255);not null"`
	LName     string `gorm:"type:varchar(255);not null; unique"`
	cellphone uint
}

type userTeam struct {
	userTeamid uint
	Sportid    uint
}
type sport struct {
	Sportid   uint
	sportname string
}
type leagues struct {
	leagueid uint
	sportid  uint
}
type teams struct {
	teamsid  uint
	leagueid uint
}
type bets struct {
	betid     uint
	matchid   uint
	startDate uint
}
type matchs struct {
	matchid uint
	sportid uint
}
type matchStat struct {
	matchStatid uint
	matchid     uint
}
type Login struct {
	ID       uint
	Username string `gorm:"type:varchar(255);not null; unique"`
	Password string `gorm:"type:varchar(255);not null"`
	UserID   uint
}

type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

type Session struct {
	ID      uint
	UUID    string `gorm:"type:varchar(255);not null"`
	Expires int64  `gorm:"type:varchar(255);not null"`
}
