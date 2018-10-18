package connection

import (
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	cachedDataSource *DataSource
	dataSourceMutex  sync.Once
)

type DataSource struct {
	Timeout    time.Duration
	BaseURL    string
	HTTPClient *http.Client
}

func NewDataSource() *DataSource {
	baseURL := os.Getenv("DATA_SOURCE_HOST")

	timeoutDuration := time.Duration(10) * time.Second

	httpClient := http.DefaultClient
	httpClient.Timeout = timeoutDuration

	return &DataSource{
		Timeout:    timeoutDuration,
		BaseURL:    baseURL,
		HTTPClient: httpClient,
	}
}

func DataSourceInstance() *DataSource {
	dataSourceMutex.Do(func() {
		cachedDataSource = NewDataSource()
	})

	return cachedDataSource
}
