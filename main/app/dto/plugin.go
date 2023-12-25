/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:18:43
 * @Desc:
 */
package dto

import (
	"selfx/app/plugin/entity"
)

type PluginList struct {
	entity.PluginInfo
	RunTime     int64  `json:"run_time"` // 时间戳
	RunError    string `json:"run_error"`
	RunCount    int    `json:"run_count"`
	RunDuration int64  `json:"run_duration"` // 毫秒
	NextRunTime int64  `json:"next_run_time"`
}
