package model

type Mapping struct {
	ID        int `gorm:"type:int;size:32;primaryKey;autoIncrement"          json:"id"`
	ArticleID int `gorm:"type:int;size:32;uniqueIndex:idx_mapping"       json:"article_id"`
	TagID     int `gorm:"type:int;size:32;uniqueIndex:idx_mapping;index" json:"tag_id"`
}
