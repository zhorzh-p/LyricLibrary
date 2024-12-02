package repositories

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

// AutoMigrate выполняет миграцию моделей
// Этот метод используется для быстрого запуска проекта в рамках MVP.
// Для продакшн-среды рекомендуется использовать миграции с версионированием.
func (m *Migrator) AutoMigrate() error {
	err := m.db.AutoMigrate(
		&music.SongEntity{},
		&music.GroupEntity{},
		&music.VerseEntity{},
	)
	if err != nil {
		logrus.Errorf("Failed to apply migrations: %v", err)
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	logrus.Info("Database migration completed successfully.")
	return nil
}
