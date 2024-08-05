package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/isantoshgyawali/music-tui/app"
)

func main() {
    p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error running program: \n %+v", err)
    }
}
