package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/mysql"
)

var MysqlDb *gorm.DB

//初始化数据库连接
func InitMysql(connectString string) (*gorm.DB, error) {
	var err error
	//打开数据库链接，返回db实例
	MysqlDb, err = gorm.Open("mysql", connectString)
	if err != nil {
		return nil, err
	}
	//取消建表时表名自动添加“s”
	MysqlDb.SingularTable(true)
	return MysqlDb, nil
}

//创建表
func CreateMysqlTable(tableModel interface{}) error {
	if MysqlDb == nil {
		return fmt.Errorf("MysqlDb为nil")
	}
	//创建表
	// Migrate the schema
	//MysqlDb.AutoMigrate(&Statics{})
	if !MysqlDb.HasTable(tableModel) {
		if err := MysqlDb.CreateTable(tableModel).Error; err != nil {
			return err
		}
	}

	return nil
}

/*
mysql合适事务示例：
自动事务
通过db.Transaction函数实现事务，如果闭包函数返回错误，则回滚事务。

db.Transaction(func(tx *gorm.DB) error {
  // 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // 返回任何错误都会回滚事务
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // 返回 nil 提交事务
  return nil
})
手动事务
在开发中经常需要数据库事务来保证多个数据库写操作的原子性。例如电商系统中的扣减库存和保存订单。
gorm事务用法：

// 开启事务
tx := db.Begin()

//在事务中执行数据库操作，使用的是tx变量，不是db。

//库存减一
//等价于: UPDATE `foods` SET `stock` = stock - 1  WHERE `foods`.`id` = '2' and stock > 0
//RowsAffected用于返回sql执行后影响的行数
rowsAffected := tx.Model(&food).Where("stock > 0").Update("stock", gorm.Expr("stock - 1")).RowsAffected
if rowsAffected == 0 {
    //如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
    //这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
    //这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
    tx.Rollback()
    return
}
err := tx.Create(保存订单).Error

//保存订单失败，则回滚事务
if err != nil {
    tx.Rollback()
} else {
    tx.Commit()
}
*/
