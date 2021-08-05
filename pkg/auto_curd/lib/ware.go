package lib

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"github.com/pkg/errors"
)

type Ware struct {
	Middlewares []config_type.MiddlewareFn
}

func (w *Ware) Use(middleware config_type.MiddlewareFn) *Ware {
	w.Middlewares = append(w.Middlewares, middleware)
	return w
}

func (w *Ware) Run(config *config_type.Config) (*config_type.Config, error) {
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
