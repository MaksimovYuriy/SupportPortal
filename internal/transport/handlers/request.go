package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

var (
	ErrBadRequest = errors.New("bad request")
)

func parsePathID(r *http.Request) (int64, error) {
	return parsePositiveInt64(r.PathValue("id"))
}

func parsePositiveInt64(value string) (int64, error) {
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		return 0, ErrBadRequest
	}
	return id, nil
}

func decodeJSONBody(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return ErrBadRequest
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return ErrBadRequest
	}
	return nil
}
