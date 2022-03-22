/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-12-14 18:55:28
 * @Description:
 */
package model

// User 对象表tb_user 字段
// 首字母大写表示public 公有的，小写是私有的 private
type User struct {
	BaseModel
	Passwd   string `gorm:"passwd" json:"passwd"`      // 密码
	Phone    string `gorm:"phone" json:"phone"`        // 电话
	Sex      int    `gorm:"sex" json:"sex"`            // 性别 0 女孩，1 男孩
	CreateAt int    `gorm:"create_at" json:"createAt"` // 创建时间
	Ip       string `gorm:"ip" json:"ip"`
}

/**
 * @description: 实现TableName方法，返回表名
 * @param {*}
 * @return {*}
 */
func (User) TableName() string {
	return "tab_user"
}
