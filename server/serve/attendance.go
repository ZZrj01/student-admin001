/*
 * @Author: WUMIANHUA
 * @Date: 2021-12-07 18:22:28
 * @LastEditTime: 2021-12-11 16:24:47
 * @Description:
 */
package serve

import (
	"server/config"
	"server/model"
)

/**
 * @description: 添加考勤管理
 * @param {model.Attendance} query
 * @return {*} 错误
 */
func AddAttendance(query model.Attendance) (err error) {
	err = config.DB.Create(&query).Error
	return
}

/**
 * @description: 查询所有的考勤信息
 * @param {*}
 * @return {*} 万能对象，错误
 */
func SelectAttendance() (interface{}, error) {
	// model.Attendance 类型的数组
	var list []model.Attendance
	// Model 指定模型 tableName 获取到表名
	// &list 结果扫描到list 指针对象
	err := config.DB.Find(&list).Error
	// err = config.DB.Model(&model.Attendance{}).Find(&list).Error
	var obj interface{}
	if err == nil {
		obj = list
	}
	return obj, err
}

/**
 * @description: 删除考勤信息
 * @param {model.Attendance} query
 * @return {*}
 */
func DelAttendance(query model.Attendance) (err error) {
	err = config.DB.Delete(&query).Error
	return
}

/**
 * @description: 编辑考勤信息
 * @param {model.Attendance} query
 * @return {*}
 */
// func ModifyAttendance(query model.Attendance) (err error) {
// 	//Update只能更新一个属性，语法是.update("key","value")key数据库字段，value变更的值
// 	//Updates修改多个属性，语法是Updates（{key:value,key2:value2...})
// 	err = config.DB.Model(&model.Attendance{}).Where("id=?", query.ID).
// 		Updates(map[string]interface{}{
// 			"create_at":   query.Create_at,
// 			"finish_at":   query.Finish_at,
// 			"session":     query.Session,
// 			"actuallynum": query.Actuallynum,
// 			"leavenum":    query.Leavenum,
// 			"shouldnum":   query.Shouldnum,
// 		}).Error
// 	return
// }
