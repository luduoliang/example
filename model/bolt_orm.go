package model

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var BoltDb *bolt.DB

//初始化BoltDb
func InitBolt(dbPath string) (*bolt.DB, error) {
	var err error
	BoltDb, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	return BoltDb, nil
}

//创建表
func CreateBoltTable(tableName string) error {
	var err error
	if BoltDb == nil {
		return fmt.Errorf("BoltDb为nil")
	}
	//创建表
	err = BoltDb.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(tableName))
		return err
	})

	return err
}

//插入或更新一条记录
func UpdateBolt(tableName string, key []byte, value []byte) error {
	var err error
	if BoltDb == nil {
		return fmt.Errorf("BoltDb为nil")
	}
	err = BoltDb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		err = b.Put(key, value)
		return err
	})

	return err
}

//查找一条记录
func FindBolt(tableName string, key []byte) (value []byte, err error) {
	if BoltDb == nil {
		return nil, fmt.Errorf("BoltDb为nil")
	}
	//查询一条数据
	err = BoltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		value = b.Get(key)
		return err
	})

	return value, err
}

//查找所有记录
func FindAllBolt(tableName string) (values [][]byte, err error) {
	if BoltDb == nil {
		return nil, fmt.Errorf("BoltDb为nil")
	}
	//遍历表数据
	err = BoltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		//遍历表
		b.ForEach(func(k, v []byte) error {
			values = append(values, v)
			return nil
		})
		return err
	})

	return values, err
}

//删除一条记录
func DeleteBolt(tableName string, key []byte) error {
	var err error
	if BoltDb == nil {
		return fmt.Errorf("BoltDb为nil")
	}
	//删除一条数据
	err = BoltDb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		err = b.Delete(key)
		return err
	})

	return err
}
