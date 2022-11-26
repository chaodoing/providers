package models

import (
	"time"
)

/******sql******
CREATE TABLE `sales` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `business_id` int unsigned NOT NULL COMMENT '商户id',
  `product_id` bigint unsigned NOT NULL COMMENT '商品id',
  `number` mediumint unsigned NOT NULL COMMENT '销售数量',
  `price` int NOT NULL COMMENT '商品价格',
  `profit` int unsigned NOT NULL COMMENT '总计盈利',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='销售日志'
******sql******/
// Sales 销售日志
type Sales struct {
	ID         uint64    `gorm:"primaryKey;column:id" json:"id"`       // 主键
	BusinessID uint      `gorm:"column:business_id" json:"businessId"` // 商户id
	ProductID  uint64    `gorm:"column:product_id" json:"productId"`   // 商品id
	Number     uint32    `gorm:"column:number" json:"number"`          // 销售数量
	Price      int       `gorm:"column:price" json:"price"`            // 商品价格
	Profit     uint      `gorm:"column:profit" json:"profit"`          // 总计盈利
	CreateAt   time.Time `gorm:"column:create_at" json:"createAt"`     // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *Sales) TableName() string {
	return "sales"
}

// SalesColumns get sql column name.获取数据库列名
var SalesColumns = struct {
	ID         string
	BusinessID string
	ProductID  string
	Number     string
	Price      string
	Profit     string
	CreateAt   string
}{
	ID:         "id",
	BusinessID: "business_id",
	ProductID:  "product_id",
	Number:     "number",
	Price:      "price",
	Profit:     "profit",
	CreateAt:   "create_at",
}
