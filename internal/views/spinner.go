package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}
)

type SpinnerView struct {
	spinner      spinner.Model
	spinnerIndex int // which spinnerIndex is active
}

func (m SpinnerView) Init() tea.Cmd {

	return m.spinner.Tick
}

func (m SpinnerView) View() string {
	var gap string
	switch m.spinnerIndex {
	case 1:
		gap = ""
	default:
		gap = " "
	}

	s := fmt.Sprintf("\n %s%s%s\n\n", m.spinner.View(), gap, "Spinning...")
	s += "h/l, ←/→: change spinner • q: exit\n"
	return s
}

func (m SpinnerView) Update(msg tea.Msg) (SpinnerView, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "h", "left":
			m.spinnerIndex--
			if m.spinnerIndex < 0 {
				m.spinnerIndex = len(spinners) - 1
			}
			m.ResetSpinner()
			return m, m.spinner.Tick
		case "l", "right":
			m.spinnerIndex++
			if m.spinnerIndex >= len(spinners) {
				m.spinnerIndex = 0
			}
			m.ResetSpinner()
			return m, m.spinner.Tick
		case "enter":
			return m, func() tea.Msg {
				return UpdateViewMessage(2)
			}
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}
func NewSpinnerView() SpinnerView {
	return SpinnerView{}

}

func (m *SpinnerView) ResetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Spinner = spinners[m.spinnerIndex]
}
