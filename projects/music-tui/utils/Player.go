package utils

import (
    "fmt"
    "os"
    "time"

    "github.com/ebitengine/oto/v3"
    "github.com/hajimehoshi/go-mp3"
)

type Player struct {
    otoCtx *oto.Context
    CurrentPlayer *oto.Player
    paused bool
    volume float32
    speed  float32
    ready chan struct{}
}

func NewPlayer() (*Player, error) {
    // Oto context Options (will use default audio device)
    // This will play all our sounds. It's configuration can't be changed later
    op := &oto.NewContextOptions{
        // recommended values: 44100 or 48000, other values might cause distortion
        SampleRate:   44100,

        // no. of channels to play sound from, 1: mono sound && 2: stereo sound
        // When settin the count to 2 and If only one speaker is available then the sound 
        // will be downmixed by OS or hardware itself no harm is caused but some effects 
        // like panning[sound movement from one speaker to another] will not function due to downmixin
        ChannelCount: 2,     
        Format: oto.FormatSignedInt16LE,
    }

    // Remember "not" to create more than one context
    // because :
    //  resouce management : reating multiple audio contexts can lead to excessive consumption of system resources
    //  audio device conflicts
    //  synchronization issues
    //  cross-platform compatibility
    //  But Main reason is: Oto is designed to support only one context at a time
    otoCtx, readyChan, err := oto.NewContext(op)
    if err != nil {
        return nil, fmt.Errorf("Error while creating a NewContext:[func NewPlayer() - player.go]\n %+v", err)
    }

    return &Player{
        otoCtx: otoCtx,
        paused: false,
        volume: 1.0,
        ready: readyChan,
    }, nil
}

func (p *Player) PlaySong(song *os.File) error {

    // p.mutex.Lock()
    // defer p.mutex.Unlock()

    // Close player if already running 
    // to avoid resource conflicts and playback error while switching songs
    if p.CurrentPlayer != nil {
        p.CurrentPlayer.Close()
    }

    // THIS IS SLOW FOR LONGER SONGS because loading long songs on memory and decoding complete file is 
    // inefficient losing performance 
    // SO JUST OPEN FILE AND DECODE basically "doing streaming"

    // fileBytes, err := io.ReadAll(song)
    // if err != nil {
    //     return fmt.Errorf("Error reading the fileBytes, [func PlaySong() - Player.go ]\n%+v", err)
    // }
    // fileBytesReader := bytes.NewReader(fileBytes) 

    // Decoding "OPENED SONG" File
    decodeMp3, err := mp3.NewDecoder(song)
    if err != nil {
        return fmt.Errorf("Error while Creating NewDecoder: [mp3.NewDecoder failed - Player.go]\n%+v\n", err)
    }

    // It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
    <- p.ready

    p.CurrentPlayer = p.otoCtx.NewPlayer(decodeMp3)
    p.CurrentPlayer.Play()

    // MONITOR PLAYBACK
    // for p.CurrentPlayer.IsPlaying() && !p.paused {
    //     time.Sleep(time.Millisecond)
    // }

    return nil
}

func (p *Player) TogglePlaySong() {
    p.paused = !p.paused
    if p.paused {
        p.CurrentPlayer.Play()
    } else {
        p.CurrentPlayer.Pause()
    }
}

func (p *Player) ChangePlaybackSpeed() {

}

func (p *Player) ChangeVolume() {

}
