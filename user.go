package codegen

//go:generate repogen

//repogen:entity
type User struct {
	ID           uint `gorm:"primary_key"`
	Email        string
	PasswordHash string
}
