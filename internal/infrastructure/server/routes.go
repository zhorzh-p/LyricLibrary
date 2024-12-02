package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zhorzh-p/LyricLibrary/cmd/server/api"
)

// RegisterRoutes регистрирует маршруты API
func RegisterRoutes(router *gin.Engine, songHandler api.SongHandler) {
	api := router.Group("/api")
	{
		// Группа маршрутов для песен
		songs := api.Group("/songs")
		{
			songs.POST("", songHandler.CreateSong)
			songs.DELETE("/:id", songHandler.DeleteSong)
			songs.PUT("/:id", songHandler.ChangeSong)
			songs.GET("/:id/verses", songHandler.GetSongVerses)
			songs.GET("", songHandler.GetSongsWithFilters)
		}
	}
}
