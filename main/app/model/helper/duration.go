/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:52:07
 * @Desc:
 */
package helper

import "time"

type DurationUnit string

var (
	DurationSecond DurationUnit = "second"
	DurationMinute DurationUnit = "minute"
	DurationHour   DurationUnit = "hour"
	DurationDay    DurationUnit = "day"
)

type Duration struct {
	Number int          `json:"number"`
	Unit   DurationUnit `json:"unit"`
}

func NewDuration(n int, unit DurationUnit) Duration {
	return Duration{Number: n, Unit: unit}
}

func (d *Duration) Duration() time.Duration {
	switch d.Unit {
	case DurationSecond:
		return time.Duration(d.Number) * time.Second
	case DurationMinute:
		return time.Duration(d.Number) * time.Minute
	case DurationHour:
		return time.Duration(d.Number) * time.Hour
	case DurationDay:
		return time.Duration(d.Number) * time.Hour * 24
	default:
		return 0
	}
}
