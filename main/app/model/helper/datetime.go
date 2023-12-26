/*
 * @Author: Coller
 * @Date: 2022-05-18 15:47:37
 * @LastEditTime: 2023-12-25 16:54:13
 * @Desc: 自定义日期，接收字符串，存日期
 */
package helper

import (
	"database/sql/driver"
	"fmt"
	"selfx/utils/date"
	"strings"
	"time"
)

type DateTime time.Time

// 写入数据库转换类型
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// 取出做类型转换
func (t *DateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DateTime(value)
		return nil
	}
	return nil
}

func (t *DateTime) String() string {
	// 如果时间 null 那么我们需要把返回的值进行修改
	if t == nil {
		return ""
	}
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%s", time.Time(*t).Format(date.FmtDateTime))
}

func (t *DateTime) IsZero() bool {
	if t == nil {
		return true
	}
	return time.Time(*t).IsZero()
}

func (t *DateTime) Unix() int64 {
	if t == nil {
		return 0
	}
	return time.Time(*t).Unix()
}

func (t *DateTime) Format(fmts string) string {
	if t == nil {
		return ""
	}
	return time.Time(*t).Format(fmts)
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.ParseInLocation(date.FmtDateTime, timeStr, time.Local)
	*t = DateTime(t1)
	return err
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	// 如果时间值是空或者0值 返回为null 如果写空字符串会报错
	if &t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format(date.FmtDateTime))), nil
}
