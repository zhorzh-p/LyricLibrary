package music

import "time"

type SongEntity struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"not null"`
	GroupID     uint           `gorm:"not null"`
	Group       *GroupEntity   `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
	ReleaseDate time.Time      `gorm:"not null"`
	Link        string         `gorm:"not null"`
	Verses      []*VerseEntity `gorm:"foreignKey:SongID;constraint:OnDelete:CASCADE;"`
}

func (SongEntity) TableName() string {
	return "songs"
}

type GroupEntity struct {
	ID    uint         `gorm:"primaryKey;autoIncrement"`
	Name  string       `gorm:"unique;not null"`
	Songs []SongEntity `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
}

func (GroupEntity) TableName() string {
	return "groups"
}

type VerseEntity struct {
	ID     uint       `gorm:"primaryKey;autoIncrement"`
	SongID uint       `gorm:"not null;index"`
	Song   SongEntity `gorm:"foreignKey:SongID;constraint:OnDelete:CASCADE;"`
	Text   string     `gorm:"not null"`
	Order  uint       `gorm:"not null"`
}

func (VerseEntity) TableName() string {
	return "verses"
}
