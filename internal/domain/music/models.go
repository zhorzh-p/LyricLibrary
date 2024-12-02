package music

import "time"

type Song struct {
	ID          uint
	Name        string
	Group       Group
	ReleaseDate time.Time
	Link        string
	Verses      []Verse
}

type Group struct {
	Name string
}

type Verse struct {
	Text  string
	Order uint
}

type ChangeSong struct {
	Name        *string
	ReleaseDate *string
	Link        *string
}
