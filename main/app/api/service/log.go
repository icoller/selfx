/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:07:47
 * @Desc:
 */
package service

import (
	"encoding/json"
	"errors"
	"io"
	"selfx/config"
	"selfx/constant"
	"selfx/utils"
	"strings"
)

func LogDirSize() (uint64, error) {
	return utils.DirSize(constant.LogDir)
}

func LogRead(id string, page, limit int) (any, error) {
	handle := config.Config.Log.GetByID(id)
	if handle == nil {
		return nil, errors.New("log id not found")
	}
	return logRead(handle.FilePath, page, limit)
}

func logRead(filePath string, page, limit int) (any, error) {
	var res []map[string]any
	list, err := utils.ReadFilePagination(filePath, page, limit)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "cannot find") || strings.Contains(errStr, "no such file") {
			return nil, nil
		}
		if errors.Is(err, io.EOF) {
			if page > 1 {
				return nil, errors.New("empty")
			}
			return nil, errors.New("not found")
		}
		return nil, err
	}
	for _, line := range list {
		var item = make(map[string]any)
		if err = json.Unmarshal([]byte(line), &item); err != nil {
			continue
		}
		res = append(res, item)
	}
	return res, nil
}
