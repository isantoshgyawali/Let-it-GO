package tui

import (
	"fmt"
	"strings"
)

/*
   func NavBarView returns
   the Navbar for navigating to other features
   in the programs like yt_downloader, maybe syncDevices
*/
func NavBarView(height, width int) string {
    return NavBarStyle(height, width).Render()
}

/* 
    func SongsListContatinerView returns 
    List of Songs adding into a View with 
    MainContentStyles defined on: package tui > const.go
*/
func SongsListContatinerView(height int, width int, songs []string) string {
    s := strings.Builder{}
    for i, song := range songs{ 
        s.WriteString(fmt.Sprintf("%d. %s\n",i+1,song))
    }
    
    return MaiContentStyle(height, width).Render(s.String())
}

/*
    func SideBarView returns 
    the SideBar which contains options for 
    playlists, favourites ... 
*/
func SideBarView(height, width int) string {
    return SideBarStyle(height, width).Render()
}

/*
    func HelpView returns
    the helping footNotes to navigate and access
    other components ie. Views and options like add, select, delete, find ....
*/
func HelpView(height, width int) string {
    return HelpContentStyle(height, width).Render()
}

/*
    func VisualizerView returns
    visualizer something like cava while 
    music is being played or maybe details about the songs
*/
func VisualizerView(height, width int) string {
    return VisualizerStyle(height, width).Render()
}
