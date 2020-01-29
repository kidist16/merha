package entity

type userPayment struct {
	payment_id     uint
	user_id        uint
	
}

type userTransaction struct {
	Transaction_id      uint
	user_id             uint
	
	
}


type userBet struct {
	userBet_id  uint
	user_id     uint
	bet_id      uint
}

type User struct {
	user_id      uint
	FName        string     `gorm:"type:varchar(255);not null"`
	LName        string     `gorm:"type:varchar(255);not null; unique"`
	cellphone     uint     
}

type userTeam struct {
	userTeam_id   uint
	Sport_id      uint
	
}
type sport struct {
	Sport_id    uint
	sport_name   string

}
type leagues struct {
	leagues_id uint
	sport_id   uint
}
type teams struct{
	teams_id     uint
	leagues_id   uint

}
type bets struct{
	bet_id     uint
	match_id   uint
	decription
	startDate
	endDate
	odds

}
type matchs struct{
	match_id  uint
	sport_id   uint
	home_team_id
	away_team_id
	data

}
type matchStat struct{
  matchStat_id     uint
  match_id         uint
	home_team_result
	away_team_result
}
type Login struct {
	ID       uint
	Username string `gorm:"type:varchar(255);not null; unique"`
	Password string `gorm:"type:varchar(255);not null"`
	UserID 	 uint
}




type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}























