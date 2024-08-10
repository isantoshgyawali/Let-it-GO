package app

import (
    "fmt"
    "strings"
)

/*
func NavBarView returns
the Navbar for navigating to other features
in the programs like yt_downloader, maybe syncDevices
*/
func NavBarView(m *Model) string {
    return NavBarStyle(m.Layout.Height, m.Layout.Width).Render()
}

/* 
func SongsListContatinerView returns 
List of Songs adding into a View with 
MainContentStyles defined on: package tui > const.go
*/
func SongsListContatinerView(m *Model, height, width int) string {
    // Preventing nil dereferencing 
    // FACT_CHECK [ You Know CrowdStrike right??? ]
    if m == nil {
        return fmt.Sprintf("[Error: SongsListContatinerView.go]\n Model is nil")
    }
    s := strings.Builder{}

    songsList := make([]string, len(m.Songs.SongsList))
    for i, song := range m.Songs.SongsList  {
        songsList[i] = fmt.Sprintf("%d. ", i+1) + song
    }

    for i, song := range songsList{ 
        if m.Layout.Cursor == i {
            s.WriteString(fmt.Sprintf("> %s\n", SelectedListStyle.Render(song)))
        } else {
            s.WriteString(fmt.Sprintf("%s\n", GeneralListStyle.Render(song)))
        }
    }

    return MainContentStyle(height, width).Render(s.String())
}

/*
func SideBarView returns 
the SideBar which contains options for 
playlists, favourites ... 
*/
func SideBarView(m *Model) string {
    return SideBarStyle(m.Layout.Height, m.Layout.Width).Render()
}

/*
func HelpView returns
the helping footNotes to navigate and access
other components ie. Views and options like add, select, delete, find ....
*/
func HelpView(m *Model) string {
    return HelpContentStyle(m.Layout.Height, m.Layout.Width).Render()
}

/*
func VisualizerView returns
visualizer something like cava while 
music is being played or maybe details about the songs
*/
func VisualizerView(m *Model) string {
    return VisualizerStyle(m.Layout.Height, m.Layout.Width).Render()
}
