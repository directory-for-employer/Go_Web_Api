package req

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func DecodeParam(r *http.Request) (uint64, error) {
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DecodeIntQuery(r *http.Request, titleQuery string) (*int, error) {
	query, err := strconv.Atoi(r.URL.Query().Get(titleQuery))
	if err != nil {
		return nil, err
	}
	return &query, nil
}

func DecodeStringQuery(r *http.Request, titleQuery string) *string {
	query := r.URL.Query().Get(titleQuery)
	return &query
}

func DecodeTimeQuery(r *http.Request, titleQuery string) (*time.Time, error) {
	layout := "01-02-2006" // 15:04:05
	query, err := time.Parse(layout, r.URL.Query().Get(titleQuery))
	if err != nil {
		return nil, err
	}
	return &query, nil
}
