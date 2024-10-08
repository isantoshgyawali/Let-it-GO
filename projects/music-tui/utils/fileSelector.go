package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func FindDirFromHome() ([]string, error) {
	root, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error finding home dir\n %+v", err)
		return nil, err
	}

	var dirs []string
	var mu sync.Mutex

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		baseDepth := strings.Count(root, string(filepath.Separator))
		currentDepth := strings.Count(path, string(filepath.Separator)) - baseDepth
		if currentDepth > 2 {
			return filepath.SkipDir
		}

		if d.IsDir() {
			// Skip dot directories and node_modules
			if strings.HasPrefix(d.Name(), ".") || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			mu.Lock()
			dirs = append(dirs, path)
			mu.Unlock()
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return dirs, nil
}

func SelectDIR() {
	dir, err := FindDirFromHome()
	if err != nil {
		fmt.Println(err)
		return
	}

	idx, err := fuzzyfinder.FindMulti(
		dir,
		func(i int) string {
			return dir[i]
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			var files []string
			err := filepath.WalkDir(dir[i], func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					files = append(files, d.Name())
				}
				return nil
			})

			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return strings.Join(files, "\n")
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Selected directories:")
	for _, i := range idx {
		fmt.Println(dir[i])
	}
}

func GetSongs() ([]string, error) {
	/* root, err := os.UserHomeDir()
	   if err != nil {
	       log.Fatal("Error finding the user Home dir : \n", err)
	   } */

	var songs []string
	err := filepath.WalkDir("/home/cosnate/songs", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		// If the DirEntry "d" is not Dir ie. it's a file:
		// checks if the extension matches to .mp3
		// if true appends the base path ie. Just the fileName [SongName in this case] to the songs slice
		if !d.IsDir() {
			if strings.ToLower(filepath.Ext(path)) == ".mp3" {
				fileName := filepath.Base(path)
				songs = append(songs, fileName)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error walking to the file path [utils.go, GetSongs]\n %+v", err)
	}

	return songs, nil
}
