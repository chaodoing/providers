package models

import (
	"time"
)

/******sql******
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `business_id` int unsigned NOT NULL COMMENT '店铺ID',
  `category_id` bigint unsigned NOT NULL COMMENT '分类ID',
  `title` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品标题',
  `prime_cost` int unsigned NOT NULL COMMENT '进货价格',
  `price` int unsigned NOT NULL COMMENT '销售价格',
  `profit` int unsigned NOT NULL COMMENT '每件盈利',
  `stock` mediumint unsigned NOT NULL COMMENT '库存数量',
  `unit` varchar(2) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '单位',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品信息表'
******sql******/
// Products 商品信息表
type Products struct {
	ID         uint64     `gorm:"primaryKey;column:id" json:"id"`       // 主键
	BusinessID uint       `gorm:"column:business_id" json:"businessId"` // 店铺ID
	CategoryID uint64     `gorm:"column:category_id" json:"categoryId"` // 分类ID
	Title      string     `gorm:"column:title" json:"title"`            // 商品标题
	PrimeCost  uint       `gorm:"column:prime_cost" json:"primeCost"`   // 进货价格
	Price      uint       `gorm:"column:price" json:"price"`            // 销售价格
	Profit     uint       `gorm:"column:profit" json:"profit"`          // 每件盈利
	Stock      uint32     `gorm:"column:stock" json:"stock"`            // 库存数量
	Unit       string     `gorm:"column:unit" json:"unit"`              // 单位
	CreateAt   time.Time  `gorm:"column:create_at" json:"createAt"`     // 创建时间
	UpdateAt   *time.Time `gorm:"column:update_at" json:"updateAt"`     // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Products) TableName() string {
	return "products"
}

// ProductsColumns get sql column name.获取数据库列名
var ProductsColumns = struct {
	ID         string
	BusinessID string
	CategoryID string
	Title      string
	PrimeCost  string
	Price      string
	Profit     string
	Stock      string
	Unit       string
	CreateAt   string
	UpdateAt   string
}{
	ID:         "id",
	BusinessID: "business_id",
	CategoryID: "category_id",
	Title:      "title",
	PrimeCost:  "prime_cost",
	Price:      "price",
	Profit:     "profit",
	Stock:      "stock",
	Unit:       "unit",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}
