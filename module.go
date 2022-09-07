package libwth

import (
  tea "github.com/charmbracelet/bubbletea"
)

type Module interface {
  View() (string)
  Update(msg tea.Msg) (tea.Model, tea.Cmd)
}

type ModuleResizeEvent struct {
  Width  int
  Height int
}

