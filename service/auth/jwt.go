package auth

import (
	"time"

	luna "github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_platform/global"
	"go.uber.org/zap"
)

// JWT Service
type JwtService struct {
}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList luna.JwtBlacklist) (err error) {
	err = global.GS_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GS_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@function: GetCacheJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetCacheJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GS_CACHE.Get(decorateKey(userName))
	return redisJWT, err
}

//@function: SetCacheJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetCacheJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GS_CONFIG.JWT.ExpiresTime) * time.Second
	_ = global.GS_CACHE.SetX(decorateKey(userName), jwt, timer)
	return err
}

func (jwtService *JwtService) LoadAll() error {
	var data []string
	err := global.GS_DB.Model(&luna.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GS_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return err
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
	return nil
}

// 缓存前缀
func decorateKey(k string) string {
	return constant.CACHE_TOKEN_PREFIX + k
}
