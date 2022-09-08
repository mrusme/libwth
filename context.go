package libwth

import (
  "errors"
  "fmt"

  "github.com/mrusme/libwth/cfg"
  "github.com/mrusme/libwth/theme"
  "go.uber.org/zap"
)

type Ctx struct {
  config       *cfg.Cfg
  moduleID     string
  Module       *cfg.CfgModule
  Log          *zap.SugaredLogger
}

func NewCtx(
  config *cfg.Cfg,
  id string,
  log *zap.SugaredLogger,
) (*Ctx, error) {
  ctx := new(Ctx)
  ctx.config = config
  ctx.moduleID = id
  ctx.Module = nil
  ctx.Log = log
  for i := 0; i < len(ctx.config.Modules); i++ {
    if ctx.config.Modules[i].ID == id {
      ctx.Module = &ctx.config.Modules[i]
      break
    }
  }
  if ctx.Module == nil {
    return nil, errors.New(
      fmt.Sprintf(
        "No module configuration with ID %s available!",
        id,
      ),
    )
  }
  ctx.moduleID = id

  return ctx, nil
}

func (ctx *Ctx) ConfigValue(key string) (string) {
  if val, ok := ctx.Module.Config[key]; ok {
    return val
  }

  return ""
}

func (ctx *Ctx) Theme() (*theme.Theme) {
  return &ctx.config.Theme
}

