/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-02 18:19:46
 * @Description:
 */
package model

type Class struct {
	ID       string `json:"id" gorm:"id"`
	Name     string `json:"name" gorm:"name"`
	CreateAt string `json:"createAt" gorm:"create_at"`
}

func (Class) TableName() string {
	return "tab_class"
}
