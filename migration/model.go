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
	UserType     UserTypes `gorm:"foreignKey:UserTypeCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Pin          string    `gorm:"column:pin"`
	PhoneNumber  string    `gorm:"column:phone_number"`
	CompanyID    string    `gorm:"column:company_id"`
	Company      Companies `gorm:"foreignKey:CompanyID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserDataID   string    `gorm:"column:user_data_id"`
	UserData     UserDatas `gorm:"references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}

type Coas struct {
	ID        string    `gorm:"primaryKey; default:gen_random_uuid()"`
	Code      string    `gorm:"column:code; unique"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime; default:current_timestamp"`
}

type Transactions struct {
	ID            string    `gorm:"primaryKey; default:gen_random_uuid()"`
	Code          string    `gorm:"column:code; unique"`
	Description   string    `gorm:"column:description"`
	Amount        float64   `gorm:"column:amount"`
	CreatedAt     time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"autoCreateTime; default:current_timestamp"`
	UserID        string    `gorm:"column:user_id"`
	User          Users     `gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyID     string    `gorm:"column:company_id"`
	Company       Companies `gorm:"foreignKey:CompanyID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedBy     string    `gorm:"column:created_by"`
	CreatedByUser Users     `gorm:"foreignKey:CreatedBy;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Journals struct {
	ID              string       `gorm:"primaryKey; default:gen_random_uuid()"`
	TransactionCode string       `gorm:"column:transaction_code"`
	Transaction     Transactions `gorm:"foreignKey:TransactionCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CoaDebitCode    string       `gorm:"column:coa_debit"`
	CoaDebit        Coas         `gorm:"foreignKey:CoaDebitCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CoaCreditCode   string       `gorm:"column:coa_credit"`
	CoaCredit       Coas         `gorm:"foreignKey:CoaCreditCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount          float64      `gorm:"column:amount"`
	CreatedAt       time.Time    `gorm:"autoCreateTime; default:current_timestamp"`
	UpdatedAt       time.Time    `gorm:"autoCreateTime; default:current_timestamp"`
	UserID          string       `gorm:"column:user_id"`
	User            Users        `gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyID       string       `gorm:"column:company_id"`
	Company         Companies    `gorm:"foreignKey:CompanyID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedBy       string       `gorm:"column:created_by"`
	CreatedByUser   Users        `gorm:"foreignKey:CreatedBy;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
