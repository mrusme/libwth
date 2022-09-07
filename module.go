package libwth

import (
  "time"
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

type HeartbeatMsg struct {
  Now    time.Time
}

func (hb *HeartbeatMsg) IsDueNow(nextUpdate time.Time) (bool) {
  if hb.Now.Unix() >= nextUpdate.Unix() {
    return true
  }
  return false
}

