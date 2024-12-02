package musicinfo

// SongDetailsClient интерфейс для взаимодействия с API музыкальных данных
type SongDetailsClient interface {
	GetSongInfo(group string, name string) (*SongInfoResponse, error)
}
