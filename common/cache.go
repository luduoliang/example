package common

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

const (
	CLEAR_CACHE_INTERVAL = 0 //自动缓存清理时间间隔，单位（秒）。为0表示不自动清理，等待缓存过期。
)

//初始化go-cache缓存
func InitCache() {
	Cache = cache.New(time.Minute*30, time.Minute*30)
	if Cache != nil {
		go ClearExprisedCache()
	}
}

//清理过期缓存
func ClearExprisedCache() {
	if CLEAR_CACHE_INTERVAL > 0 {
		clearTick := time.NewTicker(time.Second * time.Duration(CLEAR_CACHE_INTERVAL))
		for {
			select {
			case <-clearTick.C:
				{
					Cache.DeleteExpired()
				}
			}
		}
	}
}

//设置缓存
func SetCache(cacheKey string, cacheData interface{}, expriseTime time.Duration) {
	Cache.Set(cacheKey, cacheData, expriseTime)
}

//获取缓存
func GetCache(cacheKey string) (cacheData interface{}, isExist bool) {
	if cacheData, isExist = Cache.Get(cacheKey); isExist {
		return cacheData, isExist
	}
	return nil, false
}

//删除缓存
func DelCache(cacheKey string) {
	Cache.Delete(cacheKey)
}
