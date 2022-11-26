package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SalesMgr struct {
	*_BaseMgr
}

// SalesMgr open func
func SalesMgr(db *gorm.DB) *_SalesMgr {
	if db == nil {
		panic(fmt.Errorf("SalesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SalesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sales"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SalesMgr) GetTableName() string {
	return "sales"
}

// Reset 重置gorm会话
func (obj *_SalesMgr) Reset() *_SalesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SalesMgr) Get() (result Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SalesMgr) Gets() (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SalesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Sales{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SalesMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBusinessID business_id获取 商户id
func (obj *_SalesMgr) WithBusinessID(businessID uint) Option {
	return optionFunc(func(o *options) { o.query["business_id"] = businessID })
}

// WithProductID product_id获取 商品id
func (obj *_SalesMgr) WithProductID(productID uint64) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithNumber number获取 销售数量
func (obj *_SalesMgr) WithNumber(number uint32) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithPrice price获取 商品价格
func (obj *_SalesMgr) WithPrice(price int) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithProfit profit获取 总计盈利
func (obj *_SalesMgr) WithProfit(profit uint) Option {
	return optionFunc(func(o *options) { o.query["profit"] = profit })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SalesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SalesMgr) GetByOption(opts ...Option) (result Sales, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SalesMgr) GetByOptions(opts ...Option) (results []*Sales, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SalesMgr) GetFromID(id uint64) (result Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SalesMgr) GetBatchFromID(ids []uint64) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBusinessID 通过business_id获取内容 商户id
func (obj *_SalesMgr) GetFromBusinessID(businessID uint) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`business_id` = ?", businessID).Find(&results).Error

	return
}

// GetBatchFromBusinessID 批量查找 商户id
func (obj *_SalesMgr) GetBatchFromBusinessID(businessIDs []uint) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`business_id` IN (?)", businessIDs).Find(&results).Error

	return
}

// GetFromProductID 通过product_id获取内容 商品id
func (obj *_SalesMgr) GetFromProductID(productID uint64) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// GetBatchFromProductID 批量查找 商品id
func (obj *_SalesMgr) GetBatchFromProductID(productIDs []uint64) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`product_id` IN (?)", productIDs).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 销售数量
func (obj *_SalesMgr) GetFromNumber(number uint32) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`number` = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量查找 销售数量
func (obj *_SalesMgr) GetBatchFromNumber(numbers []uint32) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *_SalesMgr) GetFromPrice(price int) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *_SalesMgr) GetBatchFromPrice(prices []int) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromProfit 通过profit获取内容 总计盈利
func (obj *_SalesMgr) GetFromProfit(profit uint) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`profit` = ?", profit).Find(&results).Error

	return
}

// GetBatchFromProfit 批量查找 总计盈利
func (obj *_SalesMgr) GetBatchFromProfit(profits []uint) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`profit` IN (?)", profits).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SalesMgr) GetFromCreateAt(createAt time.Time) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_SalesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SalesMgr) FetchByPrimaryKey(id uint64) (result Sales, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Sales{}).Where("`id` = ?", id).First(&result).Error

	return
}
