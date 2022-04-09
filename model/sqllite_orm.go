package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SqlliteDb *gorm.DB

//初始化数据库连接
func InitSqllite(dbPath string) (*gorm.DB, error) {
	var err error
	SqlliteDb, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	SqlliteDb.SingularTable(true)
	return SqlliteDb, nil
}

//创建表
func CreateSqlliteTable(tableModel interface{}) error {
	if SqlliteDb == nil {
		return fmt.Errorf("SqlliteDb为nil")
	}
	//创建表
	// Migrate the schema
	//SqlliteDb.AutoMigrate(&Statics{})
	if !SqlliteDb.HasTable(tableModel) {
		if err := SqlliteDb.CreateTable(tableModel).Error; err != nil {
			return err
		}
	}

	return nil
}
