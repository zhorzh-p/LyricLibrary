package music

type GroupDatabaseRepository interface {
	GetByName(name string, loadSongs bool, loadSongVerses bool) (*GroupEntity, error)
}
