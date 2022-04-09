package common

import (
	"encoding/json"
	"errors"

	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

//初始化Redis
func InitRedis(address string, password string, db int) {
	if address != "" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		})
	}
}

//获取key
func Get(key string) (string, error) {
	if RedisClient == nil {
		return "", errors.New("未创建redis连接")
	}
	value, err := RedisClient.Do("GET", key).String()
	if err != nil {
		return "", err
	}

	return value, nil
}

//设置key
func Set(key string, value interface{}, expire int) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if expire > 0 {
		err := RedisClient.Do("SET", key, string(valueStr), "EX", expire).Err()
		if err != nil {
			return err
		}
	} else {
		err := RedisClient.Do("SET", key, string(valueStr)).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

//判断key是否存在
func Exists(key string) (bool, error) {
	if RedisClient == nil {
		return false, errors.New("未创建redis连接")
	}
	ok, err := RedisClient.Do("EXISTS", key).Bool()
	return ok, err
}

//删除key
func Del(key string) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	err := RedisClient.Do("DEL", key).Err()
	return err
}

//redis HGET
func HGet(key, field string) (string, error) {
	if RedisClient == nil {
		return "", errors.New("未创建redis连接")
	}
	value, err := RedisClient.Do("HGET", key, field).String()
	if err != nil {
		return "", err
	}

	return value, nil
}

//redis HSET
func HSet(key string, field string, value interface{}) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = RedisClient.Do("HSET", key, field, string(valueStr)).Err()
	return err
}

//redis HEXISTS
func HExists(key string, field string) (bool, error) {
	if RedisClient == nil {
		return false, errors.New("未创建redis连接")
	}
	ok, err := RedisClient.Do("HEXISTS", key, field).Bool()
	return ok, err
}

//redis HDEL
func HDel(key, field string) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	err := RedisClient.Do("HDEL", key, field).Err()
	return err
}

//redis RPUSH
func RPush(key string, value interface{}) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = RedisClient.Do("RPUSH", key, string(valueStr)).Err()
	if err != nil {
		return err
	}

	return nil
}

//redis RPOP
func RPop(key string) (string, error) {
	if RedisClient == nil {
		return "", errors.New("未创建redis连接")
	}
	value, err := RedisClient.Do("RPOP", key).String()
	if err != nil {
		return "", err
	}

	return value, nil
}

//redis LPUSH
func LPush(key string, value interface{}) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = RedisClient.Do("LPUSH", key, string(valueStr)).Err()
	if err != nil {
		return err
	}

	return nil
}

//redis LPOP
func LPop(key string) (string, error) {
	if RedisClient == nil {
		return "", errors.New("未创建redis连接")
	}
	value, err := RedisClient.Do("LPOP", key).String()
	if err != nil {
		return "", err
	}

	return value, nil
}

//redis SETNX
func SetNX(key string, value interface{}) error {
	if RedisClient == nil {
		return errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	res, err2 := RedisClient.Do("SETNX", key, string(valueStr)).Result()
	if err2 != nil {
		return err2
	}

	if resStr, ok := res.(int64); ok && resStr == 1 {
		return nil
	} else {
		return errors.New("key:" + key + " exists")
	}
}

//redis GETSET
func GetSet(key string, value interface{}) (string, error) {
	if RedisClient == nil {
		return "", errors.New("未创建redis连接")
	}
	valueStr, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	var resStr string = ""
	resStr, err = RedisClient.Do("GETSET", key, string(valueStr)).String()
	if err != nil {
		return "", err
	}

	return resStr, nil
}
