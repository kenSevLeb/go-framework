package cache

import (
	"fmt"
	"golang.org/x/sync/singleflight"
)

var dsf singleflight.Group

// 缓存不存在
var NotExist = fmt.Errorf("cache not exist")

type GetterFunc func() (interface{}, error)

// SingleGet 执行singleFlight模式
func SingleGet(key string, cacheGetter, dbGetter GetterFunc) (interface{}, error) {
	result, err := cacheGetter()
	if err == NotExist { // 缓存不存在
		// 从数据库里获取
		result, err, _ = dsf.Do(key, dbGetter)
		if err != nil {
			return nil, fmt.Errorf("get from db: %v", err)
		}

	} else if err != nil { // 获取数据失败
		return nil, fmt.Errorf("get from cache: %v", err)
	}
	return result, nil
}
