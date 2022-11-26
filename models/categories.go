package models

import (
	"time"
)

/******sql******
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `business_id` int unsigned NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `title` char(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类标题',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类描述',
  `sort` mediumint unsigned NOT NULL COMMENT '排序',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类信息表'
******sql******/
// Categories 分类信息表
type Categories struct {
	ID          uint64     `gorm:"primaryKey;column:id" json:"id"`        // 主键
	BusinessID  uint       `gorm:"column:business_id" json:"businessId"`  // 店铺ID
	Title       string     `gorm:"column:title" json:"title"`             // 分类标题
	Description string     `gorm:"column:description" json:"description"` // 分类描述
	Sort        uint32     `gorm:"column:sort" json:"sort"`               // 排序
	CreateAt    time.Time  `gorm:"column:create_at" json:"createAt"`      // 创建时间
	UpdateAt    *time.Time `gorm:"column:update_at" json:"updateAt"`      // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Categories) TableName() string {
	return "categories"
}

// CategoriesColumns get sql column name.获取数据库列名
var CategoriesColumns = struct {
	ID          string
	BusinessID  string
	Title       string
	Description string
	Sort        string
	CreateAt    string
	UpdateAt    string
}{
	ID:          "id",
	BusinessID:  "business_id",
	Title:       "title",
	Description: "description",
	Sort:        "sort",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}
