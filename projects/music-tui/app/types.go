package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/isantoshgyawali/music-tui/utils"
)

type Model struct {
	SearchInput string
        Player    *utils.Player

	Layout *Layout
	Songs  *Songs
	Help   *Help
}

type Layout struct {
	Height int
	Width  int
	Cursor int

	VisualizerHeight      int
	VisualizerWidth       int
	VisualizerColorScheme lipgloss.Color

	SongsListScrollOffset int
	SongsListViewHeight   int
}

type Help struct {
}

type Songs struct {
	SongsList []string
	Downloads []string

	Favorite []string
	Playlist []string
}
