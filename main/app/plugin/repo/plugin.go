/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:34:31
 * @Desc:
 */
package repo

import (
	"fmt"
	"selfx/init/db"

	"gorm.io/gorm/clause"
)

var Plugin = new(PluginRepo)

func init() {
	if err := db.DB.AutoMigrate(&PluginTable{}); err != nil {
		fmt.Println("migrate plugin table error: ", err.Error())
	}
}

type PluginTable struct {
	ID      string `gorm:"primaryKey;type:varchar(100);" json:"id"`
	Info    string `gorm:"type:string;"                  json:"info"`
	Options string `gorm:"type:string;"                  json:"options"`
}

func (PluginTable) TableName() string {
	return "plugin"
}

type PluginRepo struct{}

func (*PluginRepo) SaveOptions(id string, options []byte) error {
	return db.DB.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"options"})}).
		Create(&PluginTable{ID: id, Options: string(options)}).Error
}

func (*PluginRepo) SaveInfo(id string, info []byte) error {
	return db.DB.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"info"})}).
		Create(&PluginTable{ID: id, Info: string(info)}).Error
}

func (*PluginRepo) GetInfo(id string) ([]byte, error) {
	var info string
	err := db.DB.Model(&PluginTable{}).Where("id = ?", id).Pluck("info", &info).Error
	return []byte(info), err
}

func (*PluginRepo) GetOptions(id string) ([]byte, error) {
	var options string
	err := db.DB.Model(&PluginTable{}).Where("id = ?", id).Pluck("options", &options).Error
	return []byte(options), err
}
