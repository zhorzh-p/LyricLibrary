package api

import "time"

type CreateSongRequest struct {
	// Название песни
	Song string `json:"song" binding:"required" swaggertype:"string" example:"Supermassive Black Hole"`
	// Название группы-исполнителя песни
	Group string `json:"group" binding:"required" swaggertype:"string" example:"Muse"`
}

type CreateSongResponse struct {
	// Идентификатор песни
	ID uint `json:"id" binding:"required" swaggertype:"integer" example:"1"`
}

type ChangeSongRequest struct {
	// Новое название
	Name *string `json:"name" swaggertype:"string" example:"Aerials"`
	// Новая дата релиза
	ReleaseDate *string `json:"release_date" swaggertype:"string" example:"11.06.2002"`
	// Новая ссылка
	Link *string `json:"link" swaggertype:"string" example:"youtube.com"`
}

type GetVersesResponse struct {
	// Информация о странице
	Page PageData `json:"page"`
	// Список куплетов
	Verses []VerseData `json:"verses"`
}

type GetSongByFilterResponse struct {
	// Информация о странице
	Page PageData `json:"page"`
	// Список песен
	Songs []SongData `json:"songs"`
}

type PageData struct {
	// Текущая страница
	Page uint `json:"page"`
	// Количество элементов на странице
	Limit int `json:"limit"`
}

type VerseData struct {
	// Текст куплета
	Text string `json:"text" swaggertype:"string" example:"This is a verse text"`
	// ПОрядок куплета
	Order uint `json:"order" swaggertype:"integer" example:"1"`
}

type SongData struct {
	// Идентификатор песни
	ID uint `json:"id" swaggertype:"integer" example:"1"`
	// Название песни
	Name string `json:"name" swaggertye:"string" example:"Aerials"`
	// Группа-исполнитель
	Group GroupData `json:"group"`
	// Дата выпуска песни
	ReleaseDate time.Time `json:"release_date" swaggertype:"string" example:"11.06.2002"`
	// Ссылка на песню
	Link string `json:"link" swaggertype:"string" example:"youtube.com"`
	// Список куплетов
	Verses []VerseData `json:"verses"`
}

type GroupData struct {
	// Название группы
	Name string `json:"name" swaggertype:"string" example:"Muse"`
}

type EmptyResponse struct{}

type ApiError struct {
	// Описание ошибки
	Error string `json:"error"`
}
