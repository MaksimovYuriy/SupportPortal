package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrBadRequest = errors.New("bad request")
)

func parsePathID(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return 0, ErrBadRequest
	}
	return id, nil
}

func decodeJSONBody(r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return ErrBadRequest
	}
	return nil
}
