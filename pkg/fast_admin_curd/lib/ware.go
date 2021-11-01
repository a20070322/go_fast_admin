package lib

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"github.com/pkg/errors"
)

type Ware struct {
	Middlewares []types.MiddlewareFn
}

func (w *Ware) Use(middleware types.MiddlewareFn) *Ware {
	w.Middlewares = append(w.Middlewares, middleware)
	return w
}

func (w *Ware) Run(config *types.ServerConfig) (*types.ServerConfig, error) {
	var err error
	conf := config
	for _, fn := range w.Middlewares {
		conf, err = fn(conf)
		if err != nil {
			return config, errors.New(fmt.Sprintf("%s: %s", utils.GetFunctionName(fn), err))
		}
	}
	return config, nil
}
