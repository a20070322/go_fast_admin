package global

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/types"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

//全局配置
var AppSetting *types.AppConfigure

//全局日志
var Logger *zap.SugaredLogger

//全局数据库
var Db *ent.Client

//全局redis
var Rdb *redis.Client

var Rbac *casbin.Enforcer