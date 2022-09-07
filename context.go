package libwth

import (
  "errors"
  "fmt"
)

type Ctx struct {
  config       *Cfg
  moduleID     string
  Module       *CfgModule
}

func NewCtx(config *Cfg, id string) (*Ctx, error) {
  ctx := new(Ctx)
  ctx.config = config
  ctx.moduleID = id
  ctx.Module = nil
  for i := 0; i < len(ctx.config.Modules); i++ {
    if ctx.config.Modules[i].ID == id {
      ctx.Module = &ctx.config.Modules[i]
      break
    }
  }
  if ctx.Module == nil {
    return nil, errors.New(
      fmt.Sprint(
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

func (ctx *Ctx) Theme() (*CfgTheme) {
  return &ctx.config.Theme
}

