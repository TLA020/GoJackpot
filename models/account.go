package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token" sql:"-"`
}

func (account *Account) Validate() (err error) {

	if !strings.Contains(account.Email, "@") {
		err = fmt.Errorf("email address is required")
		return
	}

	if account.Username == "" {
		err = fmt.Errorf("username is required")
		return
	}

	if len(account.Password) < 6 {
		err = fmt.Errorf("required and at least 6 chars")
		return
	}

	var occurrences = 0
	// check if mail unique
	err = GetDB().Table("accounts").Where("email = ? OR username = ?", account.Email, account.Username).Count(&occurrences).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("connection error")
		return
	}

	if occurrences > 0 {
		err = fmt.Errorf("mail or username already exists")
		return
	}
	return
}
