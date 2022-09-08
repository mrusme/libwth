package cfg

import (
  "testing"
)

// TestCfgParseString tests config variable filling
func TestCfgParseString(t *testing.T) {
  cfg, err := NewCfg()
  if err != nil {
    t.Fatalf(err.Error())
  }
  if cfg.Theme.Defaults.Module.BorderColor != cfg.Theme.Colors.Secondary {
    t.Fatalf(
      `BorderColor is not %s but instead %s`,
      cfg.Theme.Colors.Secondary,
      cfg.Theme.Defaults.Module.BorderColor,
    )
  }
}

