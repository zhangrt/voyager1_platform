package global

import (
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/zhangrt/voyager1_platform/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GS_DB                  *gorm.DB
	GS_DBList              map[string]*gorm.DB
	GS_REDIS_STANDALONE    *redis.Client
	GS_REDIS_CLUSTER       *redis.ClusterClient
	GS_CONFIG              config.Server
	GS_VP                  *viper.Viper
	GS_LOG                 *zap.Logger
	GS_Concurrency_Control = &singleflight.Group{}
	BlackCache             local_cache.Cache
	lock                   sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GS_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GS_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
