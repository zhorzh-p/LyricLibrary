package music

import (
	"errors"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/clients/musicinfo"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories"
	"strings"
	"time"
)

type SongService struct {
	client          musicinfo.SongDetailsClient
	songRepository  music.SongDatabaseRepository
	groupRepository music.GroupDatabaseRepository
	verseRepository music.VerseDatabaseRepository
}

func NewSongService(
	client musicinfo.SongDetailsClient,
	songRepository music.SongDatabaseRepository,
	groupRepository music.GroupDatabaseRepository,
	verseRepository music.VerseDatabaseRepository,
) *SongService {
	return &SongService{
		client:          client,
		songRepository:  songRepository,
		groupRepository: groupRepository,
		verseRepository: verseRepository,
	}
}

func (s *SongService) CreateSong(group string, name string) (*Song, error) {
	var songEntity *music.SongEntity = nil
	var groupEntity *music.GroupEntity = nil
	var err error = nil

	groupEntity, err = s.groupRepository.GetByName(group, true, false)
	if err != nil && !errors.Is(err, repositories.ErrEntityNotFound) {
		return nil, err
	}

	if groupEntity != nil {
		for _, song := range groupEntity.Songs {
			if song.Name == name {
				return nil, ErrSongAlreadyExists
			}
		}
	}

	info, err := s.client.GetSongInfo(group, name)
	if err != nil {
		return nil, err
	}

	// Разделение текста на куплеты
	splitVerses := splitToVerses(info.Text)
	verses := make([]*music.VerseEntity, 0, len(splitVerses))
	for i, verse := range splitVerses {
		verses = append(verses, &music.VerseEntity{
			Text:  verse,
			Order: uint(i + 1),
		})
	}

	parsedDate, err := s.parseReleaseDate(info.ReleaseDate)
	if err != nil {
		return nil, err
	}

	if groupEntity == nil {
		groupEntity = &music.GroupEntity{
			Name: group,
		}
	}
	songEntity = &music.SongEntity{
		Name:        name,
		Group:       groupEntity,
		ReleaseDate: *parsedDate,
		Link:        info.Link,
		Verses:      verses,
	}

	if err := s.songRepository.Create(songEntity); err != nil {
		return nil, err
	}
	song := entityToSong(songEntity)
	return &song, nil
}

func (s *SongService) DeleteSong(id uint) error {
	return s.songRepository.Delete(id)
}

func (s *SongService) ChangeSong(id uint, request ChangeSong) error {
	var name = ""
	if request.Name != nil {
		name = *request.Name
	}
	var releaseDate = time.Time{}
	if request.ReleaseDate != nil {
		parsedDate, err := s.parseReleaseDate(*request.ReleaseDate)
		if err != nil {
			return err
		}
		releaseDate = *parsedDate
	}
	var link = ""
	if request.Link != nil {
		link = *request.Link
	}
	err := s.songRepository.Update(id, &music.SongEntity{
		Name:        name,
		ReleaseDate: releaseDate,
		Link:        link,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *SongService) GetSongVerses(id uint, offset uint, limit uint) ([]Verse, error) {
	verseEntities, err := s.verseRepository.Get(id, offset, limit)
	if err != nil {
		return nil, err
	}
	verses := make([]Verse, len(verseEntities))
	for i, verse := range verseEntities {
		verses[i] = entityToVerse(&verse)
	}
	return verses, nil
}

func (s *SongService) GetSongByFilter(
	name string,
	group string,
	releaseDateString string,
	verseMinNumber uint,
	verseMaxNumber uint,
	offset uint,
	limit uint,
) ([]Song, error) {
	var releaseDate = time.Time{}
	if releaseDateString != "" {
		rd, err := s.parseReleaseDate(releaseDateString)
		if err == nil {
			releaseDate = *rd
		}
	}
	filter := music.SongFilter{
		Name:           name,
		Group:          group,
		ReleaseDate:    releaseDate,
		VerseMinNumber: verseMinNumber,
		VerseMaxNumber: verseMaxNumber,
	}
	songEntities, err := s.songRepository.GetByFilter(filter, int(offset), int(limit))
	if err != nil {
		return nil, err
	}
	songs := make([]Song, len(songEntities))
	for i, entity := range songEntities {
		songs[i] = entityToSong(&entity)
	}
	return songs, err
}

func (s *SongService) parseReleaseDate(releaseDate string) (*time.Time, error) {
	const layout = "02.01.2006"
	parsedDate, err := time.Parse(layout, releaseDate)
	if err != nil {
		return nil, NewErrWrongDateFormat(releaseDate, err)
	}
	return &parsedDate, nil
}

func splitToVerses(text string) []string {
	verses := strings.Split(text, "\n\n")
	for i, verse := range verses {
		verses[i] = strings.TrimSpace(verse)
	}
	return verses
}
