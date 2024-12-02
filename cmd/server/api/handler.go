package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhorzh-p/LyricLibrary/internal/domain/music"
	"net/http"
	"strconv"
)

type SongHandler struct {
	service *music.SongService
}

// @title Lyric Library API
// @version 1.0
// @description Сервер для реализации библиотеки музыки. Создан для выполнения тестового задания.
// @BasePath /api/songs

func NewSongHandler(service *music.SongService) *SongHandler {
	return &SongHandler{service: service}
}

// CreateSong
//
// @Summary Добавление песни
// @Description Добавляет новую песню в базу данных
// @Accept  json
// @Produce  json
// @Param request body CreateSongRequest true "Тело запроса"
// @Success 201 {object} CreateSongResponse
// @Failure 400 {object} ApiError "Неверные данные"
// @Failure 500 {object} ApiError "Системная ошибка"
// @Router / [POST]
func (h *SongHandler) CreateSong(c *gin.Context) {
	var request CreateSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid request: " + err.Error()})
		return
	}

	song, err := h.service.CreateSong(request.Group, request.Song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{"Failed to create song: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateSongResponse{ID: song.ID})
}

// DeleteSong
//
// @Summary Удалить песню по ID
// @Description Удаляет песню из базы данных по её уникальному идентификатору (ID)
// @Produce  json
// @Param id path int true "ID песни, которую необходимо удалить"
// @Success 200 {object} EmptyResponse "Песня успешно удалена"
// @Failure 400 {object} ApiError "Неверные данные запроса"
// @Failure 500 {object} ApiError "Системная ошибка"
// @Router /{id} [DELETE]
func (h *SongHandler) DeleteSong(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid id: " + err.Error()})
		return
	}

	err = h.service.DeleteSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{"Failed to delete song: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, EmptyResponse{})
}

// ChangeSong
//
// @Summary изменить песню по ID
// @Description Обновляет данные песни по её уникальному идентификатору (ID). Можно изменить имя, группу и другие данные.
// @Produce  json
// @Param id path int true "id песни, которую необходимо обновить"
// @Param request body ChangeSongRequest true "Тело запроса"
// @Success 200 {object} EmptyResponse "Песня успешно обновлена"
// @Failure 400 {object} ApiError "Неверный формат данных"
// @Failure 404 {object} ApiError "Песня с указанным ID не найдена"
// @Failure 500 {object} ApiError "Внутренняя ошибка сервера"
// @Router /{id} [PUT]
func (h *SongHandler) ChangeSong(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid id: " + err.Error()})
		return
	}
	var request ChangeSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid request: " + err.Error()})
		return
	}
	data := music.ChangeSong{
		Name:        request.Name,
		ReleaseDate: request.ReleaseDate,
		Link:        request.Link,
	}

	err = h.service.ChangeSong(uint(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{"Failed to change song: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// GetSongVerses
//
// @Summary получить текста песни с пагинацией по куплетам
// @Description получает текста песни с пагинацией по куплетам
// @Produce  json
// @Param id path int true "id песни, куплеты которой необходимо получить"
// @Param offset query int true "offset свиг позиций куплетов от начальной"
// @Param limit query int true "limit максимальное кол-во выводящихся куплетов"
// @Success 200 {object} GetVersesResponse "выводит куплеты, соответствующие условиям"
// @Failure 400 {object} ApiError "Неверный формат данных"
// @Failure 404 {object} ApiError "Песня с указанным ID не найдена"
// @Failure 500 {object} ApiError "Внутренняя ошибка сервера"
// @Router /{id}/verses [GET]
func (h *SongHandler) GetSongVerses(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid id: " + err.Error()})
		return
	}
	offset, err := strconv.ParseUint(c.Query("offset"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid offset: " + err.Error()})
		return
	}
	limit, err := strconv.ParseUint(c.Query("limit"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid limit: " + err.Error()})
		return
	}

	verses, err := h.service.GetSongVerses(uint(id), uint(offset), uint(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{"Failed to get song verses: " + err.Error()})
		return
	}
	versesData := make([]VerseData, len(verses))
	for i, vers := range verses {
		versesData[i] = VerseData{
			Text:  vers.Text,
			Order: vers.Order,
		}
	}
	pageData := PageData{
		Page:  uint(offset/limit) + uint(1),
		Limit: len(versesData),
	}

	c.JSON(http.StatusOK, GetVersesResponse{
		Page:   pageData,
		Verses: versesData,
	})
}

// GetSongsWithFilters
//
// @Summary получить данные библиотеки с фильтрацией по всем полям и пагинацией
// @Description получает данные библиотеки с фильтрацией по всем полям и пагинацией
// @Produce  json
// @Param name query string false "параметр фильтра по имени"
// @Param group query string false "параметр фильтра по группе"
// @Param release_date query string false "параметр фильтра по дате публикации" Format(dd.mm.YYYY) example(01.06.2006)
// @Param verse_min_number query int false "параметр фильтра по минимальному значению куплета"
// @Param verse_max_number query int false "параметр фильтра по максимальному значению куплета"
// @Param offset query int true "offset свиг позиций куплетов от начальной"
// @Param limit query int true "limit максимальное кол-во выводящихся куплетов"
// @Success 200 {object} GetVersesResponse "выводит данные, соответствующие фильтрам и условиям"
// @Failure 400 {object} ApiError "Неверный запрос"
// @Failure 404 {object} ApiError "Данные в библиотеке не найдены"
// @Failure 500 {object} ApiError "Внутренняя ошибка сервера"
// @Router / [GET]
func (h *SongHandler) GetSongsWithFilters(c *gin.Context) {
	name := c.Query("name")
	group := c.Query("group")
	releaseDate := c.Query("release_date")
	verseMinNumber, err := strconv.ParseUint(c.Query("verse_min_number"), 10, 32)
	if err != nil {
		verseMinNumber = 0
	}
	verseMaxNumber, err := strconv.ParseUint(c.Query("verse_max_number"), 10, 32)
	if err != nil {
		verseMaxNumber = 0
	}
	if verseMinNumber > 0 && verseMaxNumber > 0 && verseMinNumber > verseMaxNumber {
		c.JSON(http.StatusBadRequest, ApiError{"verse_max_number should be greater or equal verse_min_number"})
		return
	}
	offset, err := strconv.ParseUint(c.Query("offset"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid offset: " + err.Error()})
		return
	}
	limit, err := strconv.ParseUint(c.Query("limit"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiError{"Invalid limit: " + err.Error()})
		return
	}

	songs, err := h.service.GetSongByFilter(name,
		group, releaseDate, uint(verseMinNumber), uint(verseMaxNumber), uint(offset), uint(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{"Failed to get song verses: " + err.Error()})
		return
	}

	songDataList := make([]SongData, len(songs))
	for i, song := range songs {
		songDataList[i] = domainToSongData(song)
	}
	data := GetSongByFilterResponse{
		Page: PageData{
			Page:  uint(offset/limit) + uint(1),
			Limit: len(songDataList),
		},
		Songs: songDataList,
	}

	c.JSON(http.StatusOK, data)
}
