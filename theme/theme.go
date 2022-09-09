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

func (t *Theme) defaultStyle(
  fg string,
  bg string,
  bd string,
) (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Margin(0, 0, 0, 0).
    Padding(1, 1).
    Foreground(lipgloss.Color(fg))

  if bg != "" {
    s = s.Background(lipgloss.Color(bg))
    if bd != "" {
      s = s.BorderBackground(lipgloss.Color(bg))
    }
  }
  if bd != "" {
    s = s.Border(lipgloss.RoundedBorder()).
      BorderForeground(lipgloss.Color(bd)).
      BorderTop(true).
      BorderLeft(true).
      BorderRight(true).
      BorderBottom(true)
  }

  return s
}

func (t *Theme) DefaultBaseStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Base.TextColor,
    t.Defaults.Base.BackgroundColor,
    t.Defaults.Base.BorderColor,
  )
}

func (t *Theme) DefaultH1Style() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.H1.TextColor,
    t.Defaults.H1.BackgroundColor,
    t.Defaults.H1.BorderColor,
  )
}

func (t *Theme) DefaultH2Style() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.H2.TextColor,
    t.Defaults.H2.BackgroundColor,
    t.Defaults.H2.BorderColor,
  )
}

func (t *Theme) DefaultH3Style() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.H3.TextColor,
    t.Defaults.H3.BackgroundColor,
    t.Defaults.H3.BorderColor,
  )
}

func (t *Theme) DefaultH4Style() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.H4.TextColor,
    t.Defaults.H4.BackgroundColor,
    t.Defaults.H4.BorderColor,
  )
}

func (t *Theme) DefaultH5Style() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.H5.TextColor,
    t.Defaults.H5.BackgroundColor,
    t.Defaults.H5.BorderColor,
  )
}

func (t *Theme) DefaultPStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.P.TextColor,
    t.Defaults.P.BackgroundColor,
    t.Defaults.P.BorderColor,
  )
}

func (t *Theme) DefaultLabelStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Label.TextColor,
    t.Defaults.Label.BackgroundColor,
    t.Defaults.Label.BorderColor,
  )
}

func (t *Theme) DefaultButtonStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Button.TextColor,
    t.Defaults.Button.BackgroundColor,
    t.Defaults.Button.BorderColor,
  )
}
func (t *Theme) DefaultButtonHoverStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Button.Hover.TextColor,
    t.Defaults.Button.Hover.BackgroundColor,
    t.Defaults.Button.Hover.BorderColor,
  )
}

func (t *Theme) DefaultModuleViewStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Module.TextColor,
    t.Defaults.Module.BackgroundColor,
    t.Defaults.Module.BorderColor,
  )
}
func (t *Theme) DefaultModuleHoverViewStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Module.Hover.TextColor,
    t.Defaults.Module.Hover.BackgroundColor,
    t.Defaults.Module.Hover.BorderColor,
  )
}
func (t *Theme) DefaultModuleActiveViewStyle() (lipgloss.Style) {
  return t.defaultStyle(
    t.Defaults.Module.Active.TextColor,
    t.Defaults.Module.Active.BackgroundColor,
    t.Defaults.Module.Active.BorderColor,
  )
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

