package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func SQLinsert(sql_action string) (string, error) {
	return "", sql.ErrNoRows
}

type User struct {
	Id     int    `gorm:"column:id; primary_key; auto_increment;not null" json:"id"`
	OpenID string `gorm:"column:open_id; type:varchar(64); index: open_id_idx" json:"open_id"`
}

func CheckAuth(phone, password string) (User, bool, error) {
	var user User
	err := db.Where(User{Phone: phone, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, false, err
	}

	if user.Id > 0 {
		return user, true, nil
	}

	return User{}, false, nil
}

func dao() {
	user, isExist, err := CheckAuth()
	if err != nil {
		//appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		//appG.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}
	fmt.Sprintf("user %+v", user)
}

func main() {
	dao()
}
