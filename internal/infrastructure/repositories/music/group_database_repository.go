package music

import (
	"errors"
	"fmt"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories"
	"gorm.io/gorm"
)

type GroupDatabaseRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupDatabaseRepository(db *gorm.DB) *GroupDatabaseRepositoryImpl {
	return &GroupDatabaseRepositoryImpl{db: db}
}

func (repo *GroupDatabaseRepositoryImpl) GetByName(name string, loadSongs bool, loadSongVerses bool) (*music.GroupEntity, error) {
	group := &music.GroupEntity{}
	tx := repo.db.
		Model(&music.GroupEntity{})
	if loadSongs {
		tx = tx.Preload("Songs")
		if loadSongVerses {
			tx = tx.Preload("Songs.Verses")
		}
	}
	err := tx.
		Where("name = ?", name).
		First(group).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrEntityNotFound
		}
		return nil, repositories.NewErrRepositoryError(fmt.Sprintf("failed to get group by name (%v)", name), err)
	}

	return group, nil
}
