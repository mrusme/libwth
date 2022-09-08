package cfg

import (
  "fmt"
  "reflect"
  "strings"

  "github.com/mrusme/libwth/theme"
  "github.com/ryankurte/go-structparse"
  "github.com/spf13/viper"
)

type CfgModule struct {
  ID                     string
  Name                   string
  Path                   string
  Enabled                bool
  Config                 map[string]string
  RefreshInterval        string
}

type Cfg struct {
  Modules                []CfgModule
  Layout                 struct {
    Rows                 []struct {
      Cells              []struct {
        ModuleID         string
        RatioX           int
        RatioY           int
      }
    }
  }
  Theme                  theme.Theme
  Log                    struct {
    File                 string
    Level                string
  }
}

func NewCfg() (Cfg, error) {
  viper.SetDefault("Theme.Colors.Primary", "#007BFF")
  viper.SetDefault("Theme.Colors.Secondary", "#6C757D")
  viper.SetDefault("Theme.Colors.Success", "#28A745")
  viper.SetDefault("Theme.Colors.Danger", "#DC3545")
  viper.SetDefault("Theme.Colors.Warning", "#FFC107")
  viper.SetDefault("Theme.Colors.Info", "#17A2B8")
  viper.SetDefault("Theme.Colors.Light", "#F8F9FA")
  viper.SetDefault("Theme.Colors.Dark", "#343A40")
  viper.SetDefault("Theme.Colors.Muted", "#6C757D")
  viper.SetDefault("Theme.Colors.White", "#FFFFFF")
  viper.SetDefault("Theme.Colors.Black", "#000000")
  viper.SetDefault("Theme.Defaults.Base.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Base.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.Base.BorderColor", "$primary")
  viper.SetDefault("Theme.Defaults.H1.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.H1.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.H1.BorderColor", "")
  viper.SetDefault("Theme.Defaults.H2.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.H2.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.H2.BorderColor", "")
  viper.SetDefault("Theme.Defaults.H3.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.H3.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.H3.BorderColor", "")
  viper.SetDefault("Theme.Defaults.H4.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.H4.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.H4.BorderColor", "")
  viper.SetDefault("Theme.Defaults.H5.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.H5.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.H5.BorderColor", "")
  viper.SetDefault("Theme.Defaults.P.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.P.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.P.BorderColor", "")
  viper.SetDefault("Theme.Defaults.Label.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Label.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.Label.BorderColor", "")
  viper.SetDefault("Theme.Defaults.Button.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Button.BackgroundColor", "$secondary")
  viper.SetDefault("Theme.Defaults.Button.BorderColor", "")
  viper.SetDefault("Theme.Defaults.Button.Hover.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Button.Hover.BackgroundColor", "$primary")
  viper.SetDefault("Theme.Defaults.Button.Hover.BorderColor", "")
  viper.SetDefault("Theme.Defaults.Module.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Module.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.Module.BorderColor", "$secondary")
  viper.SetDefault("Theme.Defaults.Module.Hover.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Module.Hover.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.Module.Hover.BorderColor", "$primary")
  viper.SetDefault("Theme.Defaults.Module.Active.TextColor", "$white")
  viper.SetDefault("Theme.Defaults.Module.Active.BackgroundColor", "")
  viper.SetDefault("Theme.Defaults.Module.Active.BorderColor", "$success")

  viper.SetDefault("Log.File", "wth.log")
  viper.SetDefault("Log.Level", "info")

  viper.SetConfigName("wth.yaml")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("/etc/")
  viper.AddConfigPath("$XDG_CONFIG_HOME/")
  viper.AddConfigPath("$HOME/.config/")
  viper.AddConfigPath("$HOME/")
  viper.AddConfigPath(".")

  /*
  viper.SetEnvPrefix("wth")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
  viper.AutomaticEnv()
  */

  if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
      return Cfg{}, err
    }
  }

  var config Cfg
  if err := viper.Unmarshal(&config); err != nil {
    return Cfg{}, err
  }

  structparse.Strings(&config, &config.Theme.Defaults)

  return config, nil
}

func (cfg *Cfg) ParseString(val string) interface{} {
  if strings.HasPrefix(val, "$") {
    valueOfColors := reflect.ValueOf(cfg.Theme.Colors)
    varName := strings.ToLower(strings.TrimLeft(val, "$"))
    varName = fmt.Sprintf(
      "%s%s",
      strings.ToUpper(string(varName[0])),
      varName[1:],
    )
    colorValue := reflect.Indirect(valueOfColors).FieldByName(varName)
    colorValueString := colorValue.String()
    return colorValueString
  }

  return val
}

