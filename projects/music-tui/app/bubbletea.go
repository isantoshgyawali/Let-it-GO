package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/isantoshgyawali/music-tui/utils"
)

func InitialModel() *Model {
    ListOfSongs, err := utils.GetSongs()
    if err != nil {
        fmt.Println(err.Error())
    }

    return &Model{
        Songs: ListOfSongs, 
        Cursor: 0,
    }
}

func (m *Model) Init() tea.Cmd {
    return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg: 
        switch msg.Type {
        // SELECTION KEYS
        case tea.KeyEnter: 
            return m, nil
        case tea.KeyTab:
            return m, nil

        // EXIT KEYS
        case tea.KeyCtrlC, tea.KeyEsc: 
            return m, tea.Quit

        // NAVIGATION KEYS
        case tea.KeyRunes: 
            switch string(msg.Runes) {
            case "j", "down": 
                if m.Cursor < len(m.Songs)-1 {
                    m.Cursor++
                }  
            case "k", "up": 
                if m.Cursor > 0 {
                    m.Cursor--
                }  
        }
    }

    // updates initially with the window dimensions then updates on each window resizes 
    // for "WINDOWS", this does nothing on resize as it doesn't support SIGWINCH signal [ copied from LSP ]
    case tea.WindowSizeMsg:
        m.height = msg.Height
        m.width = msg.Width

}
    return m, nil
}

func (m *Model) View() string {
    songsList := make([]string, len(m.Songs))
    for i, song := range m.Songs {
        if i == m.Cursor {
            songsList[i] = "> " + song
        } else {
            songsList[i] = " " + song
        }
    }

    ParentContainer := ScreenContainerStyle(m.height,m.width)

    topContent := NavBarView(m.height, m.width)
    bottomContent := HelpView(m.height, m.width)

    centerHeight := m.height - (lipgloss.Height(topContent) + lipgloss.Height(bottomContent))
    SideBarView := SideBarView(centerHeight, m.width)
    visualizerContent := VisualizerView(centerHeight, m.width)
    leftSideContent := lipgloss.JoinVertical(
            lipgloss.Center,
            SideBarView,
            visualizerContent,
        )

    songsContainerWidth := m.width - lipgloss.Width(SideBarView)
    centerContent := lipgloss.JoinHorizontal(
        lipgloss.Top,
        SongsListContatinerView(centerHeight, songsContainerWidth, m.Songs),
        leftSideContent,
        )

    return ParentContainer.Render(
            lipgloss.JoinVertical(
                lipgloss.Top,
                topContent,
                centerContent,
                bottomContent,
            ),
        )
}
