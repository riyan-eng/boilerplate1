package migration

import "time"

type Users struct {
	ID        string    `gorm:"primaryKey, default:gen_random_uuid()"`
	UserName  string    `gorm:"column:username"`
	Email     string    `gorm:"email"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

type UserDatas struct {
	ID         string `gorm:"primaryKey, default:gen_random_uuid()"`
	Name       string
	Status     string
	BirthDate  time.Time
	BirthPlace string
	Address    string
}
