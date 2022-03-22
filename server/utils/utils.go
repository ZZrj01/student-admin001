/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:25:53
 * @Description:
 */
package utils

import (
	"strings"

	"github.com/google/uuid"
)

/**
 * @description: 创建uuid
 * @param {*}
 * @return {*}
 */
func GetUUID() (id string) {
	return strings.ReplaceAll(string(uuid.New().String()), "-", "")
}
