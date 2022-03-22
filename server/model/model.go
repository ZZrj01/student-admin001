/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:24:42
 * @Description:
 */
/**
 * @description: 父model 对象 拥有公有的字段
 * @param {*}
 * @return {*}
 */
package model

/**
 * @description: 公有的父对象
 * @param {*}
 * @return {*}
 */
type BaseModel struct {
	ID   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}
