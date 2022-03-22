/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:24:36
 * @Description:
 */
package config

import (
	"server/model"

	"github.com/jinzhu/gorm"
)

var (
	Conf model.Config
	DB   *gorm.DB
)
