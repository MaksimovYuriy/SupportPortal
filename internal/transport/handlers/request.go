package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func parsePathID(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func decodeJSONBody(r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return err
	}
	return nil
}
