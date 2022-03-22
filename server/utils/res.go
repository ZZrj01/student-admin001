/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:25:48
 * @Description:
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

/**
 * @description:返回成功
 * @param {*gin.Context} c
 * @param {interface{}} data
 * @return {*}
 */
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Res{
		Code:    200,
		Data:    data,
		Message: "success",
	})
}

/**
 * @description:
 * @param {*gin.Context} c
 * @param {int} code
 * @param {string} message
 * @return {*}
 */
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Res{
		Code:    code,
		Data:    nil,
		Message: message,
	})
}
