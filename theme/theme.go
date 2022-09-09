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
      TextColor            string
      BackgroundColor      string
      BorderColor          string
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

func (t *Theme) DefaultStyle(
  fg string,
  bg string,
  bd string,
  padding [4]int,
  margin [4]int,
) (lipgloss.Style) {
  s := lipgloss.NewStyle().
    Foreground(lipgloss.Color(fg)).
    Padding(padding[0], padding[1], padding[2], padding[3]).
    Margin(margin[0], margin[1], margin[2], margin[3])

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
      BorderBottom(true).
  }

  return s
}

func (t *Theme) DefaultBaseStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Base.TextColor,
    t.Defaults.Base.BackgroundColor,
    t.Defaults.Base.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultH1Style() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.H1.TextColor,
    t.Defaults.H1.BackgroundColor,
    t.Defaults.H1.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{1, 0, 1, 0},
  )
}

func (t *Theme) DefaultH2Style() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.H2.TextColor,
    t.Defaults.H2.BackgroundColor,
    t.Defaults.H2.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{1, 0, 1, 0},
  )
}

func (t *Theme) DefaultH3Style() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.H3.TextColor,
    t.Defaults.H3.BackgroundColor,
    t.Defaults.H3.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{1, 0, 1, 0},
  )
}

func (t *Theme) DefaultH4Style() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.H4.TextColor,
    t.Defaults.H4.BackgroundColor,
    t.Defaults.H4.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{1, 0, 1, 0},
  )
}

func (t *Theme) DefaultH5Style() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.H5.TextColor,
    t.Defaults.H5.BackgroundColor,
    t.Defaults.H5.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{1, 0, 1, 0},
  )
}

func (t *Theme) DefaultPStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.P.TextColor,
    t.Defaults.P.BackgroundColor,
    t.Defaults.P.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultLabelStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Label.TextColor,
    t.Defaults.Label.BackgroundColor,
    t.Defaults.Label.BorderColor,
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultButtonStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Button.TextColor,
    t.Defaults.Button.BackgroundColor,
    t.Defaults.Button.BorderColor,
    [4]int{1, 4, 1, 4},
    [4]int{1, 2, 1, 2},
  )
}
func (t *Theme) DefaultButtonHoverStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Button.Hover.TextColor,
    t.Defaults.Button.Hover.BackgroundColor,
    t.Defaults.Button.Hover.BorderColor,
    [4]int{1, 4, 1, 4},
    [4]int{1, 2, 1, 2},
  )
}

func (t *Theme) DefaultModuleViewStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Module.TextColor,
    t.Defaults.Module.BackgroundColor,
    t.Defaults.Module.BorderColor,
    [4]int{1, 1, 1, 1},
    [4]int{0, 0, 0, 0},
  )
}
func (t *Theme) DefaultModuleHoverViewStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Module.Hover.TextColor,
    t.Defaults.Module.Hover.BackgroundColor,
    t.Defaults.Module.Hover.BorderColor,
    [4]int{1, 1, 1, 1},
    [4]int{0, 0, 0, 0},
  )
}
func (t *Theme) DefaultModuleActiveViewStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Defaults.Module.Active.TextColor,
    t.Defaults.Module.Active.BackgroundColor,
    t.Defaults.Module.Active.BorderColor,
    [4]int{1, 1, 1, 1},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextPrimaryStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Primary,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextSecondaryStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Secondary,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextSuccessStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Success,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextDangerStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Danger,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextWarningStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Warning,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextInfoStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Info,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextLightStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Light,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextDarkStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Dark,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextMutedStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Muted,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextWhiteStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.White,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

func (t *Theme) DefaultTextBlackStyle() (lipgloss.Style) {
  return t.DefaultStyle(
    t.Colors.Black,
    "",
    "",
    [4]int{0, 0, 0, 0},
    [4]int{0, 0, 0, 0},
  )
}

