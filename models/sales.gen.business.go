package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _BusinessMgr struct {
	*_BaseMgr
}

// BusinessMgr open func
func BusinessMgr(db *gorm.DB) *_BusinessMgr {
	if db == nil {
		panic(fmt.Errorf("BusinessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BusinessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("business"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BusinessMgr) GetTableName() string {
	return "business"
}

// Reset 重置gorm会话
func (obj *_BusinessMgr) Reset() *_BusinessMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BusinessMgr) Get() (result Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BusinessMgr) Gets() (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BusinessMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Business{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_BusinessMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 店铺名称
func (obj *_BusinessMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithAvatar avatar获取 店铺头像
func (obj *_BusinessMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithMobile mobile获取 手机号码
func (obj *_BusinessMgr) WithMobile(mobile uint64) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithPassword password获取 登录密码
func (obj *_BusinessMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithAddress address获取 店铺地址
func (obj *_BusinessMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithCreateAt create_at获取 创建时间
func (obj *_BusinessMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 修改时间
func (obj *_BusinessMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_BusinessMgr) GetByOption(opts ...Option) (result Business, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BusinessMgr) GetByOptions(opts ...Option) (results []*Business, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_BusinessMgr) GetFromID(id uint) (result Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_BusinessMgr) GetBatchFromID(ids []uint) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 店铺名称
func (obj *_BusinessMgr) GetFromTitle(title string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 店铺名称
func (obj *_BusinessMgr) GetBatchFromTitle(titles []string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 店铺头像
func (obj *_BusinessMgr) GetFromAvatar(avatar string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 店铺头像
func (obj *_BusinessMgr) GetBatchFromAvatar(avatars []string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号码
func (obj *_BusinessMgr) GetFromMobile(mobile uint64) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号码
func (obj *_BusinessMgr) GetBatchFromMobile(mobiles []uint64) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 登录密码
func (obj *_BusinessMgr) GetFromPassword(password string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 登录密码
func (obj *_BusinessMgr) GetBatchFromPassword(passwords []string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 店铺地址
func (obj *_BusinessMgr) GetFromAddress(address string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 店铺地址
func (obj *_BusinessMgr) GetBatchFromAddress(addresss []string) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_BusinessMgr) GetFromCreateAt(createAt time.Time) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_BusinessMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 修改时间
func (obj *_BusinessMgr) GetFromUpdateAt(updateAt time.Time) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 修改时间
func (obj *_BusinessMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BusinessMgr) FetchByPrimaryKey(id uint) (result Business, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Business{}).Where("`id` = ?", id).First(&result).Error

	return
}
