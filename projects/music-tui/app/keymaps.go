package app

import (
	"log"
	"os"
	"path/filepath"
)

// Execute on : J / Down
func OnDownKeysPress(m *Model) {
	if m.Layout.Cursor < len(m.Songs.SongsList)-1 {
		m.Layout.Cursor++
	} else {
		m.Layout.Cursor = 0
	}
}

// Execute on : K / Up
func OnUpKeysPress(m *Model) {
	if m.Layout.Cursor > 0 {
		m.Layout.Cursor--
	} else {
		m.Layout.Cursor = len(m.Songs.SongsList)
	}
}

// Execute on : Return / Enter
func OnEnterPress(m *Model) {
	cursor := m.Layout.Cursor
	selected_song := m.Songs.SongsList[cursor]

	song_file, err := os.Open(filepath.Join("/home/cosnate/songs", selected_song))
	if err != nil {
		log.Printf("Error Opening file [func OnEnterPress() - keymaps.go]\n %+v", err)
	}
	defer song_file.Close()

	if err := m.Player.PlaySong(song_file); err != nil {
	    log.Printf("Error Playing song [func OnEnterPress() - keymaps.go]\n%+v", err)
	}
}

func OnSpaceBarPress(m *Model) {
    m.Player.TogglePlaySong()
}
