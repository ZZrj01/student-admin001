/*
 * @Author: WUMIANHUA
 * @Date: 2021-12-07 18:18:41
 * @LastEditTime: 2021-12-14 23:24:29
 * @Description:
 */
package model

type Attendance struct {
	ID        string `json:"id" gorm:"id"`
	Create_at string `json:"create_at" gorm:"create_at"`
	Finish_at string `json:"finish_at" gorm:"finish_at"`
	Session   string `json:"session"   gorm:"session"`
	Classnum  string `json:"classnum"  gorm:"classnum"`
	// AttendanceTime string `json:"attendanceTime"  gorm:"attendanceTime"`
}

func (Attendance) TableName() string {
	return "tab_attendance"
}
