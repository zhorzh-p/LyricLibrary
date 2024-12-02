package music

import (
	"fmt"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories"

	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
	"gorm.io/gorm"
)

type SongDatabaseRepositoryImpl struct {
	db *gorm.DB
}

// NewSongDatabaseRepository создает новый репозиторий для работы с песнями.
func NewSongDatabaseRepository(db *gorm.DB) *SongDatabaseRepositoryImpl {
	return &SongDatabaseRepositoryImpl{db: db}
}

func (repo *SongDatabaseRepositoryImpl) GetByFilter(filter music.SongFilter, offset int, limit int) ([]music.SongEntity, error) {
	var songs []music.SongEntity

	var tx *gorm.DB
	if filter.Group != "" {
		tx = repo.db.
			InnerJoins("Group", repo.db.Where(&music.GroupEntity{Name: filter.Group}))
	} else {
		tx = repo.db.
			InnerJoins("Group")
	}
	if filter.VerseMinNumber > 0 && filter.VerseMaxNumber >= filter.VerseMinNumber {
		tx = tx.Preload("Verses", "\"order\" >= ? and \"order\" <= ?", filter.VerseMinNumber, filter.VerseMaxNumber)
	} else if filter.VerseMinNumber > 0 {
		tx = tx.Preload("Verses", "\"order\" >= ?", filter.VerseMinNumber)
	} else if filter.VerseMaxNumber > 0 {
		tx = tx.Preload("Verses", "\"order\" <= ?", filter.VerseMaxNumber)
	} else {
		tx = tx.Preload("Verses")
	}
	if filter.Name != "" {
		tx.Where(&music.SongEntity{Name: filter.Name})
	}
	if !filter.ReleaseDate.IsZero() {
		tx.Where(&music.SongEntity{ReleaseDate: filter.ReleaseDate})
	}
	err := tx.Limit(limit).Offset(offset).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

// Create добавляет новую песню в базу данных.
func (repo *SongDatabaseRepositoryImpl) Create(song *music.SongEntity) error {
	if err := repo.db.Create(song).Error; err != nil {
		return repositories.NewErrRepositoryError("failed to save song: %w", err)
	}
	return nil
}

// Delete удаляет песню и ее куплеты по id.
func (repo *SongDatabaseRepositoryImpl) Delete(id uint) error {
	err := repo.db.Delete(&music.SongEntity{}, id).Error
	if err != nil {
		return repositories.NewErrRepositoryError(fmt.Sprintf("failed to delete song by id (%v)", id), err)
	}
	return nil
}

func (repo *SongDatabaseRepositoryImpl) Update(id uint, entity *music.SongEntity) error {
	err := repo.db.Model(&music.SongEntity{}).
		Where("id = ?", id).
		Updates(entity).Error
	if err != nil {
		return repositories.NewErrRepositoryError(fmt.Sprintf("failed to update song by id (%v)", id), err)
	}
	return nil
}
