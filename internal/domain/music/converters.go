package music

import (
	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
)

func entityToSong(e *music.SongEntity) Song {
	verses := make([]Verse, len(e.Verses))
	for i, entity := range e.Verses {
		verses[i] = entityToVerse(entity)
	}
	return Song{
		ID:          e.ID,
		Name:        e.Name,
		Group:       entityToGroup(e.Group),
		ReleaseDate: e.ReleaseDate,
		Link:        e.Link,
		Verses:      verses,
	}
}

func entityToGroup(e *music.GroupEntity) Group {
	return Group{
		Name: e.Name,
	}
}

func entityToVerse(e *music.VerseEntity) Verse {
	return Verse{
		Text:  e.Text,
		Order: e.Order,
	}
}
