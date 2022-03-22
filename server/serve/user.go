/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-12-14 18:46:23
 * @Description:
 */
package serve

import (
	"fmt"
	"server/config"
	"server/model"
)

/**
 * @description: 判断用户名和密码一致的用户是否存在，存在就返回用户信息，不错误就返回错误，
 * @param {model.User} user
 * @return {*}
 */
func IsHasUser(user model.User) (temp model.User, err error) {
	// 查询这个用户名和密码一致的用户映射到temp 对象去
	err = config.DB.Model(&model.User{}).Where("name = ? and passwd = ?", user.Name, user.Passwd).Find(&temp).Error
	if err != nil {
		return
	}

	if temp.ID == "" {
		err = fmt.Errorf("用户名和密码错误！")
		return
	}
	return
}

func ModifyIp(query model.User) (err error) {
	err = config.DB.Model(&model.User{}).Where("name=?", query.Name).Update("ip", query.Ip).Error
	return
}
