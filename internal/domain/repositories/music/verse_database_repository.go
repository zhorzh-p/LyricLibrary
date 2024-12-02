package music

type VerseDatabaseRepository interface {
	Get(songId uint, offset uint, limit uint) ([]VerseEntity, error)
}
