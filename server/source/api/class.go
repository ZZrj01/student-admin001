/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 22:22:37
 * @Description:
 */
package api

import (
	"fmt"
	"server/model"
	"server/serve"
	"server/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 添加班級
 * @param {*gin.Context} c
 * @return {*}
 */
func AddClass(c *gin.Context) {
	var query model.Class
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	// 继续执行
	query.ID = utils.GetUUID()
	t := time.Now().Unix()
	query.CreateAt = time.Unix(t, 0).Format("2006-01-02 03:04")
	err = serve.AddClass(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	// 返回新建的id
	utils.Success(c, query)
}

/**
 * @description: 查询所有的教室
 * @param {*gin.Context} c
 * @return {*}
 */
func GetClass(c *gin.Context) {
	// 调用数据层方法SelectClass
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	list, err := serve.SelectClass()
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, list)
	// c.JSON(200, map[string]interface{}{
	// 	"code": 1008,
	// 	"data": list,
	// })
}

/**
 * @description: 删除班级
 * @param {*gin.Context} c
 * @return {*}
 */
func DelClass(c *gin.Context) {
	var query model.Class
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	err = serve.DelClass(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, nil)

}

/**
 * @description: 更新班级
 * @param {*}
 * @return {*}
 */
func ModifyClass(c *gin.Context) {
	var query model.Class
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	//判断name不能为空
	//服务端必须验证数据，不能完全相信前端数据效验
	//严格写法
	//if strings.Trim(query.Name, "") == "" || len(query.Name) == 0 {
	if len(strings.Trim(query.Name, "")) == 0 {
		utils.Error(c, 1001, "班级名称不能为空")
		return
	}
	//判断一下班级名字是否重复了，重复了就不允许提交
	err = serve.ModifyClass(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, nil)
}
