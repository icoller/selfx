/*
 * @Author: Coller
 * @Date: 2022-05-08 11:59:54
 * @LastEditTime: 2023-12-26 17:28:30
 * @Desc: 评论表
 */
package model

type Comment struct {
	ID             uint   `gorm:"column:id;primaryKey;type:uint;size:30;comment:主键" json:"id"`
	UserId         uint   `gorm:"column:user_id;type:uint;size:30;comment:用户ID" json:"userId"`
	UserName       uint   `gorm:"column:user_name;type:varchar(50);comment:用户名" json:"userName"`
	ArticleId      uint   `gorm:"column:article_id;type:uint;size:30;comment:文章ID" json:"articleId"`
	ReplyCommentId uint   `gorm:"column:reply_comment_id;type:uint;size:30;comment:回复评论ID" json:"replyCommentId"`
	Content        string `gorm:"column:content;type:varchar(255);comment:评论内容;" json:"content"`
	Ip             string `gorm:"column:ip;type:varchar(64);comment:IP地址;" json:"ip"`
	Port           string `gorm:"column:port;type:varchar(30);comment:登录端口;" json:"port"`
	Region         string `gorm:"column:region;type:varchar(128);comment:地区;" json:"region"`
	Browser        string `gorm:"column:browser;type:varchar(64);comment:浏览器;" json:"browser"`
	Os             string `gorm:"column:os;type:varchar(64);comment:操作系统;" json:"os"`
	ModelTime
}
