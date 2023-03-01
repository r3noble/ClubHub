package models

// struct to represent the table in SQL database
type User struct {
	Name string `gorm:"type:varchar(255);not null"`
	//constraint to ensure email doesn't already exist and sets it as primary key for each row
	Email     string `gorm:"uniqueIndex;primary_key"`
	Password  string `gorm:"type:varchar(255);not null"`
	userLevel int    `gorm:"not null"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
