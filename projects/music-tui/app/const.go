package app

import "github.com/charmbracelet/lipgloss"

/* COLORS */
const (
    PRIMARY = lipgloss.Color("#ffffff")
    SECONDARY = lipgloss.Color("")
    TERTIARY = lipgloss.Color("#f5f5f5")

    HIGHLIGHTER = lipgloss.Color("")
    SUCCESS = lipgloss.Color("")
    DANGER = lipgloss.Color("")
)

/* TEXT STYLES */
var (
    titleTextStyles = lipgloss.NewStyle(). 
        Background(lipgloss.Color(HIGHLIGHTER)).
        Foreground(lipgloss.Color(PRIMARY)). 
        Bold(true)

    highlightedTextStyles = lipgloss.NewStyle(). 
        Background(lipgloss.Color(TERTIARY)). 
        Foreground(lipgloss.Color(PRIMARY))

    footNoteTextStyles = lipgloss.NewStyle()
)

/* BACKGROUND STYLES */
var (
    primaryBackgroundStyles = lipgloss.NewStyle().Background(PRIMARY)
    secondaryBackgroundStyles = lipgloss.NewStyle().Background(SECONDARY)

    highlighedBackgroundStyles = lipgloss.NewStyle().Background(HIGHLIGHTER)
)

/* BORDER STYLES */
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
    inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
    activeTabBorder   = tabBorderWithBottom("┘", " ", "└")

    normalBorder = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
    roundedBorder = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
)

/* LAYOUT STYLES */

// SETTING THE DYNAMIC HEIGHT AND WIDTH BASED
// ON THE SCREEN SIZE OF THE TERMINAL WINDOW TO PARENT_CONTAINER
func ScreenContainerStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle().
        Padding(1).
        Width(width).
        Height(height)
}

// FOR HELP_SECTION
func NavBarStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("#abc123")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Width(width - 5)
}

// SIMILARY, FOR MAIN_CONTENT ie. SongsContainer
func MaiContentStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("#0ac123")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Padding(1,3,1,3).
        Width(width - 5).
        Height(height - 5)
}

// FOR SIDEBARS
func SideBarStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("#fcc123")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Padding(1).
        Width(int(0.7/3.0 * float64(width)))
}

// FOR CAVA_LIKE_VISUALIZER
func VisualizerStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("#eec123")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Padding(1).
        Width(int(0.7/3.0 * float64(width)))
}

// FOR HELP_SECTION
func HelpContentStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("#cbc123")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Padding(1).
        Width(width - 5)
}

/* UI COMPONENTS */
var (
    spinnerStyle = lipgloss.NewStyle().Foreground(PRIMARY)

    helpStyle = lipgloss.NewStyle(). 
        Foreground(lipgloss.Color(TERTIARY)). 
        Italic(true)
)

