package mysql

import "example/model"

//表记录数据结构
type MysqlModel struct {
	Id   int    `gorm:"primary_key;unsigned" json:"id"`
	Name string `gorm:"type:VARCHAR(255);not null;default:''" json:"name"` //姓名
}

//添加
func (info *MysqlModel) Create() (*MysqlModel, error) {
	err := model.MysqlDb.Model(&MysqlModel{}).Create(&info).Error
	return info, err
}

//查询
func (info *MysqlModel) GetInfo(id int) *MysqlModel {
	model.MysqlDb.Model(&MysqlModel{}).Where("id = ?", id).First(&info)
	return info
}

//更新，用结构体更新，只更新非0字段
func (info *MysqlModel) Save() error {
	return model.MysqlDb.Model(&MysqlModel{}).Save(&info).Error
}

//更新部分字段，用结构体更新，只更新非0字段
func (info *MysqlModel) Update(fields map[string]interface{}) error {
	return model.MysqlDb.Model(&MysqlModel{}).Where("id = ?", info.Id).Update(fields).Error
}

//删除，info中要包含id字段
func (info *MysqlModel) Delete() error {
	return model.MysqlDb.Model(&MysqlModel{}).Delete(info).Error
}

//列表,filterFields为筛选的字段和值
func (info *MysqlModel) List(filterFields map[string]interface{}, order string) []*MysqlModel {
	var list []*MysqlModel = make([]*MysqlModel, 0)
	query := model.MysqlDb.Model(&MysqlModel{})

	if len(filterFields) > 0 {
		for k, v := range filterFields {
			query = query.Where(k+" = ?", v)
		}
	}

	query.Order(order).Find(&list)
	return list
}

//列表,filterFields为筛选的字段和值
func (info *MysqlModel) ListPage(filterFields map[string]interface{}, order string, page int, pageSize int) []*MysqlModel {
	offSet := (page - 1) * pageSize
	var list []*MysqlModel = make([]*MysqlModel, 0)
	query := model.MysqlDb.Model(&MysqlModel{})

	if len(filterFields) > 0 {
		for k, v := range filterFields {
			query = query.Where(k+" = ?", v)
		}
	}

	query.Order(order).Offset(offSet).Limit(pageSize).Find(&list)
	return list
}
