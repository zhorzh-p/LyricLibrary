package music

import (
	"github.com/kr/pretty"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories"
	"gorm.io/gorm"
)

type VerseDatabaseRepositoryImpl struct {
	db *gorm.DB
}

func NewVerseDatabaseRepository(db *gorm.DB) *VerseDatabaseRepositoryImpl {
	return &VerseDatabaseRepositoryImpl{db: db}
}

func (repo *VerseDatabaseRepositoryImpl) Get(songId uint, offset uint, limit uint) ([]music.VerseEntity, error) {
	var verses []music.VerseEntity
	err := repo.db.Model(&music.VerseEntity{}).
		Where("song_id = ? and \"order\" >= ? and \"order\" < ?", songId, offset+1, offset+limit+1).
		Find(&verses).Error

	if err != nil {
		return nil, repositories.NewErrRepositoryError(
			pretty.Sprintf("failed to load verses by offset = %v and limit = %v", offset, limit),
			err,
		)
	}
	return verses, nil
}
