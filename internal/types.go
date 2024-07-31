package internal

type Interval struct {
	Start, End int
	Name       string
}

type ResultInterval struct {
	Start, End int
	Names      []string
}
