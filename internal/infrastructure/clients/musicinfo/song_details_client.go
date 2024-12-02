package musicinfo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	domain "github.com/zhorzh-p/LyricLibrary/internal/domain/clients/musicinfo"
)

// RestSongDetailsClient реализация SongDetailsClient
type RestSongDetailsClient struct {
	client *resty.Client
}

// NewRestSongDetailsClient создает новый RestSongDetailsClient
func NewRestSongDetailsClient(baseURL string) *RestSongDetailsClient {
	client := resty.New().
		EnableTrace().
		SetBaseURL(baseURL)

	return &RestSongDetailsClient{client: client}
}

// GetSongInfo получает информацию о песне из API
func (c *RestSongDetailsClient) GetSongInfo(group string, song string) (*domain.SongInfoResponse, error) {
	result := &domain.SongInfoResponse{}

	response, err := c.client.R().
		SetQueryParam("group", group).
		SetQueryParam("song", song).
		SetResult(result).
		ForceContentType("application/json").
		Get("/info")

	if err != nil {
		logrus.WithError(err).Errorln("Request failed")
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	if response.IsError() {
		logrus.WithError(err).Errorln("Received error response: Status=%d, Body=%s", response.StatusCode(), response.String())
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode(), response.String())
	}

	logrus.Debugf("Successful response: Status=%d, Result=%+v", response.StatusCode(), result)

	return result, nil
}
