/*
 * @Author: coller
 * @Date: 2023-12-25 13:22:03
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:22:14
 * @Desc:
 */
package date

import "time"

func TodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
