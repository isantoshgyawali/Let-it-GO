package pkg

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}