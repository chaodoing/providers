package models

import (
	"time"
)

/******sql******
CREATE TABLE `business` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺名称',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺头像',
  `mobile` bigint unsigned NOT NULL COMMENT '手机号码',
  `password` char(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录密码',
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺地址',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='店铺信息表'
******sql******/
// Business 店铺信息表
type Business struct {
	ID       uint       `gorm:"primaryKey;column:id" json:"id"`   // 主键
	Title    string     `gorm:"column:title" json:"title"`        // 店铺名称
	Avatar   string     `gorm:"column:avatar" json:"avatar"`      // 店铺头像
	Mobile   uint64     `gorm:"column:mobile" json:"mobile"`      // 手机号码
	Password string     `gorm:"column:password" json:"password"`  // 登录密码
	Address  string     `gorm:"column:address" json:"address"`    // 店铺地址
	CreateAt time.Time  `gorm:"column:create_at" json:"createAt"` // 创建时间
	UpdateAt *time.Time `gorm:"column:update_at" json:"updateAt"` // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Business) TableName() string {
	return "business"
}

// BusinessColumns get sql column name.获取数据库列名
var BusinessColumns = struct {
	ID       string
	Title    string
	Avatar   string
	Mobile   string
	Password string
	Address  string
	CreateAt string
	UpdateAt string
}{
	ID:       "id",
	Title:    "title",
	Avatar:   "avatar",
	Mobile:   "mobile",
	Password: "password",
	Address:  "address",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}
