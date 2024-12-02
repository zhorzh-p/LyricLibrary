package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toorop/gin-logrus"
	"github.com/zhorzh-p/LyricLibrary/cmd/server/api"
	domainMusic "github.com/zhorzh-p/LyricLibrary/internal/domain/music"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/clients/musicinfo"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/repositories/music"
	"os"
)

func NewServer() (*gin.Engine, error) {
	db, err := repositories.NewGormDB()
	if err != nil {
		return nil, err
	}

	songRepository := music.NewSongDatabaseRepository(db)
	groupRepository := music.NewGroupDatabaseRepository(db)
	verseRepository := music.NewVerseDatabaseRepository(db)

	// Выполнение миграций
	// Эти миграции предназначены для быстрого запуска проекта (MVP).
	// Для продакшена следует использовать версионированную миграцию.
	migrator := repositories.NewMigrator(db)
	err = migrator.AutoMigrate()
	if err != nil {
		return nil, err
	}

	client := musicinfo.NewRestSongDetailsClient(os.Getenv("MUSICINFO_URL"))

	service := domainMusic.NewSongService(client, songRepository, groupRepository, verseRepository)
	songHandler := api.NewSongHandler(service)

	router := gin.Default()
	router.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())

	RegisterRoutes(router, *songHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logrus.Info("Server initialized successfully.")
	return router, nil
}
