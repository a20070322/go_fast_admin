package cache_service

import (
	"context"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/go-redis/redis/v8"
)

func Init(ctx context.Context) *Cache {
	art := &Cache{}
	art.ctx = ctx
	art.AdminUserPrefix = "admin_user_"
	art.AdminRolePrefix = "admin_role_"
	art.Rdb = global.Rdb
	return art
}

type Cache struct {
	ctx             context.Context
	AdminUserPrefix string
	AdminRolePrefix string
	Rdb             *redis.Client
}




