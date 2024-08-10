package app

import "github.com/charmbracelet/lipgloss"

/* COLORS */
const (
    PRIMARY = lipgloss.Color("#ffffff")
    SECONDARY = lipgloss.Color("")
    TERTIARY = lipgloss.Color("#f5f5f5")

    HIGHLIGHTER = lipgloss.Color("#0ac123")
    SUCCESS = lipgloss.Color("")
    DANGER = lipgloss.Color("")
)

// LIST STYLES 
var (
    GeneralListStyle = lipgloss.NewStyle().
        MarginBottom(1)

    SelectedListStyle = lipgloss.NewStyle().
        Background(lipgloss.Color("#808080")).
        Foreground(lipgloss.Color("#ffffff")).
        MarginBottom(1)
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
func MainContentStyle(height, width int) lipgloss.Style {
    return lipgloss.NewStyle(). 
        Background(lipgloss.Color("")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PRIMARY).
        Padding(1,2).
        Width(width - 5).
        MaxHeight(height - 2)

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
        Width(width - 5)
}

/* UI COMPONENTS */
var (
    spinnerStyle = lipgloss.NewStyle().Foreground(PRIMARY)

    helpStyle = lipgloss.NewStyle(). 
        Foreground(lipgloss.Color(TERTIARY)). 
        Italic(true)
)

