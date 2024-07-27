package tui

import "github.com/charmbracelet/lipgloss"

/* STYLES */
var (
    titleStyle = lipgloss.NewStyle(). 
    Foreground(lipgloss.Color("#FAFAFA")). 
    Background(lipgloss.Color("#7D56F4")). 
    Padding(0, 1)

    selectedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FAFAFA")).
        Background(lipgloss.Color("#2D7D9A"))

    fileStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FAFAFA"))

    errorStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FF0000"))
)
