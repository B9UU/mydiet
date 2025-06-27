package searchbox

import (
	"context"
	"encoding/json"
	"fmt"
	"mydiet/internal/logger"
	"mydiet/internal/types"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {
		logger.Log.Info("in")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		url := "https://dummyjson.com/products/search?q="
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return types.FailedRequest{Err: err}
		}
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return types.FailedRequest{Err: err}
		}

		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return types.FailedRequest{
				Err: fmt.Errorf("Failed Request with status: %d",
					response.StatusCode),
			}
		}

		logger.Log.Info("succesfull Req")
		var data types.DummyJson
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return types.FailedRequest{Err: err}
		}
		logger.Log.Info("extracting titles")
		return types.SuccessDummy{Data: data}
	}
}
