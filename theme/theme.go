package theme

import (
  "github.com/charmbracelet/lipgloss"
)

type ThemeBasics struct {
  TextColor                string
  BackgroundColor          string
  BorderColor              string
}

type Theme struct {
  Colors                   struct {
    Primary                string
    Secondary              string
    Success                string
    Danger                 string
    Warning                string
    Info                   string
    Light                  string
    Dark                   string
    Muted                  string
    White                  string
    Black                  string
  }
  Defaults                 struct {
    Base                   ThemeBasics
    H1                     ThemeBasics
    H2                     ThemeBasics
    H3                     ThemeBasics
    H4                     ThemeBasics
    H5                     ThemeBasics
    P                      ThemeBasics
    Label                  ThemeBasics
    Button                 struct {
                           ThemeBasics
      Hover                ThemeBasics
    }
    Module                 struct {
      TextColor            string
      BackgroundColor      string
      BorderColor          string
      Hover                ThemeBasics
      Active               ThemeBasics
    }
  }
}



func (t *Theme) DefaultModuleViewStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Margin(0, 0, 0, 0).
    Padding(1, 1).
    Foreground(lipgloss.Color(
      t.Defaults.Module.TextColor))

  if t.Defaults.Module.BackgroundColor != "" {
    s = s.Background(lipgloss.Color(
      t.Defaults.Module.BackgroundColor))
    if t.Defaults.Module.BorderColor != "" {
      s = s.BorderBackground(lipgloss.Color(
        t.Defaults.Base.BackgroundColor))
    }
  }
  if t.Defaults.Module.BorderColor != "" {
    s = s.Border(lipgloss.RoundedBorder()).
      BorderForeground(lipgloss.Color(
        t.Defaults.Module.BorderColor)).
      BorderTop(true).
      BorderLeft(true).
      BorderRight(true).
      BorderBottom(true)
  }

  return s
}

func (t *Theme) DefaultLabelStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Defaults.Label.TextColor))

  if t.Defaults.Label.BackgroundColor != "" {
    s = s.Background(lipgloss.Color(
      t.Defaults.Label.BackgroundColor))
    if t.Defaults.Label.BorderColor != "" {
      s = s.BorderBackground(lipgloss.Color(
        t.Defaults.Base.BackgroundColor))
    }
  }
  if t.Defaults.Label.BorderColor != "" {
    s = s.Border(lipgloss.RoundedBorder()).
      BorderForeground(lipgloss.Color(
        t.Defaults.Label.BorderColor)).
      BorderTop(true).
      BorderLeft(true).
      BorderRight(true).
      BorderBottom(true)
  }

  return s
}

func (t *Theme) DefaultTextPrimaryStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Primary))

  return s
}

func (t *Theme) DefaultTextSecondaryStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Secondary))

  return s
}

func (t *Theme) DefaultTextSuccessStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Success))

  return s
}

func (t *Theme) DefaultTextDangerStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Danger))

  return s
}

func (t *Theme) DefaultTextWarningStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Warning))

  return s
}

func (t *Theme) DefaultTextInfoStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Info))

  return s
}

func (t *Theme) DefaultTextLightStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Light))

  return s
}

func (t *Theme) DefaultTextDarkStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Dark))

  return s
}

func (t *Theme) DefaultTextMutedStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Muted))

  return s
}

func (t *Theme) DefaultTextWhiteStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.White))

  return s
}

func (t *Theme) DefaultTextBlackStyle() (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(
      t.Colors.Black))

  return s
}
