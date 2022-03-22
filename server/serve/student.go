/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:25:12
 * @Description: 数据层
 */
package serve

import (
	"server/config"
	"server/model"
)

/**
 * @description: 添加学生
 * @param {model.Student} query
 * @return {*} 错误
 */
func AddStudent(query model.Student) (err error) {
	err = config.DB.Create(&query).Error
	return
}

/**
 * @description: 查询所有的学生信息
 * @param {*}
 * @return {*} 万能对象，错误
 */
func SelectStudent() (interface{}, error) {
	// model.Student 类型的数组
	var list []model.Student
	// Model 指定模型 tableName 获取到表名
	// &list 结果扫描到list 指针对象
	err := config.DB.Find(&list).Error
	// err = config.DB.Model(&model.Student{}).Find(&list).Error
	var obj interface{}
	if err == nil {
		obj = list
	}
	return obj, err
}

/**
 * @description: 删除学生
 * @param {model.Student} query
 * @return {*}
 */
func DelStudent(query model.Student) (err error) {
	err = config.DB.Delete(&query).Error
	return
}

/**
 * @description: 编辑学生信息
 * @param {model.Student} query
 * @return {*}
 */
func ModifyStudent(query model.Student) (err error) {
	//Update只能更新一个属性，语法是.update("key","value")key数据库字段，value变更的值
	//Updates修改多个属性，语法是Updates（{key:value,key2:value2...})
	err = config.DB.Model(&model.Student{}).Where("id=?", query.ID).
		Updates(map[string]interface{}{
			"name":         query.Name,
			"sex":          query.Sex,
			"phone":        query.Phone,
			"professional": query.Professional,
			"group":        query.Group,
			"seat":         query.Seat,
		}).Error
	return
}
