package libwth

import (
	"testing"

	"github.com/mrusme/libwth/cfg"
)

// TestNewCtx tests initialization of new Context
func TestNewCtx(t *testing.T) {
	c, err := cfg.NewCfg()
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.Modules = append(c.Modules, cfg.CfgModule{ID: "test"})
	_, err = NewCtx(
		&c, "test", nil,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
