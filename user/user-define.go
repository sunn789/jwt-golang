package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
	Name     string `json:"name,omitempty" form:"name"`
}

func LoadTestUser() *User {
	// Just for demonstration purpose, we create a user with the encrypted "test" password.
	// In real-world applications, you might load the user from the database by specific parameters (email, username, etc.)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{Password: string(hashedPassword), Name: "Test user"}
}
