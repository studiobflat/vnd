package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func Get[RES any](url string, heades map[string]string) (*RES, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range heades {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result RES
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
