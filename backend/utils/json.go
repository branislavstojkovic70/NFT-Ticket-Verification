package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSON(writer http.ResponseWriter, status int, dto any) error {
    writer.Header().Add("Content-Type", "application/json")
    writer.WriteHeader(status)

    return json.NewEncoder(writer).Encode(dto)
}

func ReadJSONSimple(request *http.Request, dto any) error {
    body, err := io.ReadAll(request.Body)
    if err != nil {
        return err
    }
    defer request.Body.Close()

    return json.Unmarshal(body, dto)
}