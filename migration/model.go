package migration

import "time"

type UserTypes struct {
	ID        string `gorm:"primaryKey; default:gen_random_uuid()"`
	Code      string `gorm:"unique"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}

type UserDatas struct {
	ID         string `gorm:"primaryKey; default:gen_random_uuid()"`
	Name       string
	Status     string
	BirthDate  time.Time
	BirthPlace string
	Address    string
	CreatedAt  time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}

type Companies struct {
	ID        string    `gorm:"primaryKey; default:gen_random_uuid()"`
	Name      string    `gorm:"unique"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}

type Users struct {
	ID           string    `gorm:"primaryKey; default:gen_random_uuid()"`
	UserName     string    `gorm:"column:username; unique"`
	UserTypeCode string    `gorm:"column:user_type_code"`
	UserType     UserTypes `gorm:"references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	PhoneNumber  string    `gorm:"column:phone_number"`
	CompanyID    string    `gorm:"column:company_id"`
	Company      Companies `gorm:"references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserDataID   string    `gorm:"column:user_data_id"`
	UserData     UserDatas `gorm:"references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}
