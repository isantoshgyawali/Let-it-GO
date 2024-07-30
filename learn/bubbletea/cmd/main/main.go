/*
   1. Define Model [ for the state of the program ]
   2. Initialize Model
   3. Update Model [ updates the state ]
   4. View [ returns String representing the UI ]
*/

package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	topInput      string
	mainContent   string
	rightTop      string
	rightBottom   string
	bottomContent string
	width         int
	height        int
}

func initialModel() model {
	return model{
		topInput:      "> ",
		mainContent:   "Main Content Area",
		rightTop:      "Right Top",
		rightBottom:   "Right Bottom",
		bottomContent: "Bottom Content",
		width:         100,
		height:        30,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "backspace":
			if len(m.topInput) > 2 {
				m.topInput = m.topInput[:len(m.topInput)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.topInput += msg.String()
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	// Define styles
	topStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(m.width - 4).
		Height(1)

	/* mainStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(m.width*2/3 - 4).
		Height(m.height - 8)

	rightTopStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(m.width/3 - 4).
		Height((m.height-8)/2) */

	/* rightBottomStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(m.width/3 - 4).
		Height((m.height-8)/2) */

	/* bottomStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(m.width - 4).
		Height(3) */

	// Render components
	top := topStyle.Render(m.topInput)
	/* main := mainStyle.Render(m.mainContent) */
	/* rightTop := rightTopStyle.Render(m.rightTop) */
	/* rightBottom := rightBottomStyle.Render(m.rightBottom) */
	/* bottom := bottomStyle.Render(m.bottomContent) */

	// Combine right column
        
	/* right := lipgloss.JoinVertical(lipgloss.Left, rightTop, rightBottom) */

	// Join main and right
	/* middle := lipgloss.JoinHorizontal(lipgloss.Top, main, right) */

	// Combine all parts
	return lipgloss.JoinVertical(lipgloss.Left,
		top,
		/* middle, */
		/* bottom, */
	)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
