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
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}

func (account *Account) Validate() (err error) {

	if !strings.Contains(account.Email, "@") {
		err = fmt.Errorf("email address is required")
		return
	}

	if len(account.Password) < 6 {
		err = fmt.Errorf("required and at least 6 chars")
		return
	}

	var occurrences = 0
	// check if mail unique
	err = GetDB().Table("accounts").Where("email = ?", account.Email).Count(&occurrences).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("connection error")
		return
	}

	if occurrences > 0 {
		err = fmt.Errorf("address already in use by another user")
		return
	}
	return
}



