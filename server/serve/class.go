/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:24:52
 * @Description: 数据层
 */
package serve

import (
	"server/config"
	"server/model"
)

/**
 * @description: 添加班级
 * @param {model.Class} query
 * @return {*} 错误
 */
func AddClass(query model.Class) (err error) {
	err = config.DB.Create(&query).Error
	return
}

/**
 * @description: 查询所有的班级信息
 * @param {*}
 * @return {*} 万能对象，错误
 */
func SelectClass() (interface{}, error) {
	// model.Class 类型的数组
	var list []model.Class
	// Model 指定模型 tableName 获取到表名
	// &list 结果扫描到list 指针对象
	err := config.DB.Find(&list).Error
	// err = config.DB.Model(&model.Class{}).Find(&list).Error
	var obj interface{}
	if err == nil {
		obj = list
	}
	return obj, err
}

/**
 * @description: 删除班级
 * @param {model.Class} query
 * @return {*}
 */
func DelClass(query model.Class) (err error) {
	err = config.DB.Delete(&query).Error
	return
}

/**
 * @description: 编辑班级信息
 * @param {model.Class} query
 * @return {*}
 */
func ModifyClass(query model.Class) (err error) {
	//Update只能更新一个属性，语法是.update("key","value")key数据库字段，value变更的值
	//Updates修改多个属性，语法是Updates（{key:value,key2:value2...})
	err = config.DB.Model(&model.Class{}).Where("id=?", query.ID).Update("name", query.Name).Error
	return
}
