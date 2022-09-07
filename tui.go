package libwth

import (
  "github.com/charmbracelet/lipgloss"
)

func DefaultModuleViewStyle(t *CfgTheme) (lipgloss.Style) {
  return lipgloss.NewStyle().
    Margin(0, 0, 0, 0).
    Padding(1, 1).
    Foreground(lipgloss.Color(
      t.Defaults.Module.TextColor)).
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color(
      t.Defaults.Module.BorderColor)).
    BorderTop(true).
    BorderLeft(true).
    BorderRight(true).
    BorderBottom(true)
}

