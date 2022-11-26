package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductsMgr struct {
	*_BaseMgr
}

// ProductsMgr open func
func ProductsMgr(db *gorm.DB) *_ProductsMgr {
	if db == nil {
		panic(fmt.Errorf("ProductsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("products"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductsMgr) GetTableName() string {
	return "products"
}

// Reset 重置gorm会话
func (obj *_ProductsMgr) Reset() *_ProductsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductsMgr) Get() (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductsMgr) Gets() (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Products{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_ProductsMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBusinessID business_id获取 店铺ID
func (obj *_ProductsMgr) WithBusinessID(businessID uint) Option {
	return optionFunc(func(o *options) { o.query["business_id"] = businessID })
}

// WithCategoryID category_id获取 分类ID
func (obj *_ProductsMgr) WithCategoryID(categoryID uint64) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithTitle title获取 商品标题
func (obj *_ProductsMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithPrimeCost prime_cost获取 进货价格
func (obj *_ProductsMgr) WithPrimeCost(primeCost uint) Option {
	return optionFunc(func(o *options) { o.query["prime_cost"] = primeCost })
}

// WithPrice price获取 销售价格
func (obj *_ProductsMgr) WithPrice(price uint) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithProfit profit获取 每件盈利
func (obj *_ProductsMgr) WithProfit(profit uint) Option {
	return optionFunc(func(o *options) { o.query["profit"] = profit })
}

// WithStock stock获取 库存数量
func (obj *_ProductsMgr) WithStock(stock uint32) Option {
	return optionFunc(func(o *options) { o.query["stock"] = stock })
}

// WithUnit unit获取 单位
func (obj *_ProductsMgr) WithUnit(unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = unit })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ProductsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 修改时间
func (obj *_ProductsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_ProductsMgr) GetByOption(opts ...Option) (result Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductsMgr) GetByOptions(opts ...Option) (results []*Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_ProductsMgr) GetFromID(id uint64) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_ProductsMgr) GetBatchFromID(ids []uint64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBusinessID 通过business_id获取内容 店铺ID
func (obj *_ProductsMgr) GetFromBusinessID(businessID uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`business_id` = ?", businessID).Find(&results).Error

	return
}

// GetBatchFromBusinessID 批量查找 店铺ID
func (obj *_ProductsMgr) GetBatchFromBusinessID(businessIDs []uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`business_id` IN (?)", businessIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类ID
func (obj *_ProductsMgr) GetFromCategoryID(categoryID uint64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类ID
func (obj *_ProductsMgr) GetBatchFromCategoryID(categoryIDs []uint64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 商品标题
func (obj *_ProductsMgr) GetFromTitle(title string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 商品标题
func (obj *_ProductsMgr) GetBatchFromTitle(titles []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromPrimeCost 通过prime_cost获取内容 进货价格
func (obj *_ProductsMgr) GetFromPrimeCost(primeCost uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`prime_cost` = ?", primeCost).Find(&results).Error

	return
}

// GetBatchFromPrimeCost 批量查找 进货价格
func (obj *_ProductsMgr) GetBatchFromPrimeCost(primeCosts []uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`prime_cost` IN (?)", primeCosts).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 销售价格
func (obj *_ProductsMgr) GetFromPrice(price uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 销售价格
func (obj *_ProductsMgr) GetBatchFromPrice(prices []uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromProfit 通过profit获取内容 每件盈利
func (obj *_ProductsMgr) GetFromProfit(profit uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`profit` = ?", profit).Find(&results).Error

	return
}

// GetBatchFromProfit 批量查找 每件盈利
func (obj *_ProductsMgr) GetBatchFromProfit(profits []uint) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`profit` IN (?)", profits).Find(&results).Error

	return
}

// GetFromStock 通过stock获取内容 库存数量
func (obj *_ProductsMgr) GetFromStock(stock uint32) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`stock` = ?", stock).Find(&results).Error

	return
}

// GetBatchFromStock 批量查找 库存数量
func (obj *_ProductsMgr) GetBatchFromStock(stocks []uint32) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`stock` IN (?)", stocks).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 单位
func (obj *_ProductsMgr) GetFromUnit(unit string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`unit` = ?", unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量查找 单位
func (obj *_ProductsMgr) GetBatchFromUnit(units []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`unit` IN (?)", units).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ProductsMgr) GetFromCreateAt(createAt time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_ProductsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 修改时间
func (obj *_ProductsMgr) GetFromUpdateAt(updateAt time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 修改时间
func (obj *_ProductsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductsMgr) FetchByPrimaryKey(id uint64) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Products{}).Where("`id` = ?", id).First(&result).Error

	return
}
