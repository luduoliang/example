package sqllite

import "example/model"

//表记录数据结构
type SqlliteModel struct {
	Id   int    `gorm:"primary_key;unsigned" json:"id"`
	Name string `gorm:"type:VARCHAR(255);not null;default:''" json:"name"` //姓名
}

//添加
func (info *SqlliteModel) Create() (*SqlliteModel, error) {
	err := model.SqlliteDb.Model(&SqlliteModel{}).Create(&info).Error
	return info, err
}

//查询
func (info *SqlliteModel) GetInfo(id int) *SqlliteModel {
	model.SqlliteDb.Model(&SqlliteModel{}).Where("id = ?", id).First(&info)
	return info
}

//更新，用结构体更新，只更新非0字段
func (info *SqlliteModel) Save() error {
	return model.SqlliteDb.Model(&SqlliteModel{}).Save(&info).Error
}

//更新部分字段，用结构体更新，只更新非0字段
func (info *SqlliteModel) Update(fields map[string]interface{}) error {
	return model.SqlliteDb.Model(&SqlliteModel{}).Where("id = ?", info.Id).Update(fields).Error
}

//删除，info中要包含id字段
func (info *SqlliteModel) Delete() error {
	return model.SqlliteDb.Model(&SqlliteModel{}).Delete(info).Error
}

//列表,filterFields为筛选的字段和值
func (info *SqlliteModel) List(filterFields map[string]interface{}, order string) []*SqlliteModel {
	var list []*SqlliteModel = make([]*SqlliteModel, 0)
	query := model.SqlliteDb.Model(&SqlliteModel{})

	if len(filterFields) > 0 {
		for k, v := range filterFields {
			query = query.Where(k+" = ?", v)
		}
	}

	query.Order(order).Find(&list)
	return list
}

//列表,filterFields为筛选的字段和值
func (info *SqlliteModel) ListPage(filterFields map[string]interface{}, order string, page int, pageSize int) []*SqlliteModel {
	offSet := (page - 1) * pageSize
	var list []*SqlliteModel = make([]*SqlliteModel, 0)
	query := model.SqlliteDb.Model(&SqlliteModel{})

	if len(filterFields) > 0 {
		for k, v := range filterFields {
			query = query.Where(k+" = ?", v)
		}
	}

	query.Order(order).Offset(offSet).Limit(pageSize).Find(&list)
	return list
}
