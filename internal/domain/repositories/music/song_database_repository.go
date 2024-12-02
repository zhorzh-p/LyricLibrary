package music

import "time"

type SongDatabaseRepository interface {
	GetByFilter(filter SongFilter, offset int, limit int) ([]SongEntity, error)

	Create(song *SongEntity) error

	Delete(id uint) error

	Update(id uint, entity *SongEntity) error
}

type SongFilter struct {
	Name           string
	Group          string
	ReleaseDate    time.Time
	VerseMinNumber uint
	VerseMaxNumber uint
}
