package app

type Model struct {
    SearchInput string

    Layout *Layout
    Songs *Songs
    Help *Help
}

type Layout struct {
    Height int
    Width int

    Cursor int

    VisualizerHeight int 
    VisualizerWidth int 
    VisualizerColorScheme string 
}

type Help struct {
     
}

type Songs struct {
    SongsList []string
    Downloads []string

    Favorite []string   
    Playlist []string

}
