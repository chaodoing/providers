package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CategoriesMgr struct {
	*_BaseMgr
}

// CategoriesMgr open func
func CategoriesMgr(db *gorm.DB) *_CategoriesMgr {
	if db == nil {
		panic(fmt.Errorf("CategoriesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoriesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("categories"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoriesMgr) GetTableName() string {
	return "categories"
}

// Reset 重置gorm会话
func (obj *_CategoriesMgr) Reset() *_CategoriesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CategoriesMgr) Get() (result Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoriesMgr) Gets() (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CategoriesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Categories{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_CategoriesMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBusinessID business_id获取 店铺ID
func (obj *_CategoriesMgr) WithBusinessID(businessID uint) Option {
	return optionFunc(func(o *options) { o.query["business_id"] = businessID })
}

// WithTitle title获取 分类标题
func (obj *_CategoriesMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取 分类描述
func (obj *_CategoriesMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithSort sort获取 排序
func (obj *_CategoriesMgr) WithSort(sort uint32) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCreateAt create_at获取 创建时间
func (obj *_CategoriesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 修改时间
func (obj *_CategoriesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_CategoriesMgr) GetByOption(opts ...Option) (result Categories, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CategoriesMgr) GetByOptions(opts ...Option) (results []*Categories, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_CategoriesMgr) GetFromID(id uint64) (result Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_CategoriesMgr) GetBatchFromID(ids []uint64) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBusinessID 通过business_id获取内容 店铺ID
func (obj *_CategoriesMgr) GetFromBusinessID(businessID uint) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`business_id` = ?", businessID).Find(&results).Error

	return
}

// GetBatchFromBusinessID 批量查找 店铺ID
func (obj *_CategoriesMgr) GetBatchFromBusinessID(businessIDs []uint) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`business_id` IN (?)", businessIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 分类标题
func (obj *_CategoriesMgr) GetFromTitle(title string) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 分类标题
func (obj *_CategoriesMgr) GetBatchFromTitle(titles []string) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 分类描述
func (obj *_CategoriesMgr) GetFromDescription(description string) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 分类描述
func (obj *_CategoriesMgr) GetBatchFromDescription(descriptions []string) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_CategoriesMgr) GetFromSort(sort uint32) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_CategoriesMgr) GetBatchFromSort(sorts []uint32) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_CategoriesMgr) GetFromCreateAt(createAt time.Time) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_CategoriesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 修改时间
func (obj *_CategoriesMgr) GetFromUpdateAt(updateAt time.Time) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 修改时间
func (obj *_CategoriesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CategoriesMgr) FetchByPrimaryKey(id uint64) (result Categories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Categories{}).Where("`id` = ?", id).First(&result).Error

	return
}
