package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

var redisClient cache.Cache
var enableRedis, _ = beego.AppConfig.Bool("enableRedis")
var ExperTime, _ = beego.AppConfig.Int("experTime")

func init() {
	if enableRedis {
		config := map[string]string{
			"key":      beego.AppConfig.String("redisKey"),
			"conn":     beego.AppConfig.String("redisConn"),
			"dbNum":    beego.AppConfig.String("redisDbNum"),
			"password": beego.AppConfig.String("redisPwd"),
		}
		bytes, _ := json.Marshal(config)//返回[]byte类型
		redisClient, err = cache.NewCache("redis", string(bytes))
		if err != nil {
			beego.Info("redis数据库连接失败")
		} else {
			beego.Info("redis数据库连接成功")
		}
	}
}

//定义缓存结构体，内部私有
type cacheDb struct {}

//封装redis操作接口
func (c cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		redisClient.Put(key, string(bytes), time.Second*time.Duration(ExperTime))
	}
}

func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if cacheData := redisClient.Get(key); cacheData != nil {
			cacheValue, ok := cacheData.([]uint8) //类型断言
			if !ok {
				beego.Info("缓存数据获取失败")
				return false
			}
			json.Unmarshal([]byte(cacheValue), obj)
			return true
		}
		return false
	}
	return false
}

func (c cacheDb) ClearAll() {
	if enableRedis {
		redisClient.ClearAll()
	}
}

//实例化结构体，并对外暴露方法，但隐藏内部方法实现细节
var CacheDb = &cacheDb{}