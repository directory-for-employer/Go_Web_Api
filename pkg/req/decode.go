package req

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
