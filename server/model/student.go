/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-17 22:54:18
 * @Last Modified by: WUMIANHUA
 * @Last Modified time: 2021-10-17 22:54:40
 */
package model

type Student struct {
	ID           string `json:"id" gorm:"id"`
	Name         string `json:"name" gorm:"name"`
	CreateAt     string `json:"createAt" gorm:"create_at"`
	Sex          string `json:"sex" gorm:"sex"`
	Professional string `json:"professional" gorm:"professional"`
	Phone        string `json:"phone" gorm:"phone"`
	Group        string `json:"group" gorm:"group"`
	Seat         string `json:"seat" gorm:"seat"`
}

func (Student) TableName() string {
	return "tab_student"
}
