package libwth

import (
  "strings"

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

type CfgTheme struct {
  Base                   struct {
    TextColor            string
    BackgroundColor      string
    BorderColor          string
  }
  Button                 struct {
    TextColor            string
    BackgroundColor      string
    BorderColor          string
  }
  Module                 struct {
    TextColor            string
    BackgroundColor      string
    BorderColor          string
  }
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
  Theme                  CfgTheme
}

func NewCfg() (Cfg, error) {
  viper.SetDefault("Theme.Base.TextColor", "#FFFFFF")
  viper.SetDefault("Theme.Base.BackgroundColor", "")
  viper.SetDefault("Theme.Base.BorderColor", "#FFFFFF")
  viper.SetDefault("Theme.Button.TextColor", "#FFFFFF")
  viper.SetDefault("Theme.Button.BackgroundColor", "")
  viper.SetDefault("Theme.Button.BorderColor", "#FFFFFF")
  viper.SetDefault("Theme.Module.TextColor", "#FFFFFF")
  viper.SetDefault("Theme.Module.BackgroundColor", "")
  viper.SetDefault("Theme.Module.BorderColor", "#FFFFFF")

  viper.SetConfigName("wth.yaml")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("/etc/")
  viper.AddConfigPath("$XDG_CONFIG_HOME/")
  viper.AddConfigPath("$HOME/.config/")
  viper.AddConfigPath("$HOME/")
  viper.AddConfigPath(".")

  viper.SetEnvPrefix("wth")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
      return Cfg{}, err
    }
  }

  var config Cfg
  if err := viper.Unmarshal(&config); err != nil {
    return Cfg{}, err
  }

  return config, nil
}

