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
        Songs: &Songs{
            SongsList: ListOfSongs,
        }, 

        Layout: &Layout{
            Cursor: 0,
        },
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
                if m.Layout.Cursor < len(m.Songs.SongsList)-1 {
                    m.Layout.Cursor++
                }  
            case "k", "up": 
                if m.Layout.Cursor > 0 {
                    m.Layout.Cursor--
                }  
        }
    }

    // updates initially with the window dimensions then updates on each window resizes 
    // for "WINDOWS", this does nothing on resize as it doesn't support SIGWINCH signal [ copied from LSP ]
    case tea.WindowSizeMsg:
        m.Layout.Height = msg.Height
        m.Layout.Width = msg.Width

}
    return m, nil
}

func (m *Model) View() string {
    ParentContainer := ScreenContainerStyle(m.Layout.Height,m.Layout.Width) // ParentContainer: Inside which Every other containers are aligned
    topContent := NavBarView(m) // TopContent:    Where every other features are introduced like yt_downloader
    bottomContent := HelpView(m)// BottomContent: Where resides the help section

    SideBarView := SideBarView(m)
    visualizerContent := VisualizerView(m)
    leftSideContent := lipgloss.JoinVertical(
            lipgloss.Center,
            SideBarView,
            visualizerContent,
        )


    songsContainerHeight := m.Layout.Height - (lipgloss.Height(topContent) + lipgloss.Height(bottomContent))
    songsContainerWidth := m.Layout.Width - lipgloss.Width(SideBarView)
    centerContent := lipgloss.JoinHorizontal(
        lipgloss.Top,
        SongsListContatinerView(m, songsContainerHeight, songsContainerWidth),
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
