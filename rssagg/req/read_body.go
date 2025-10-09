package req

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadToJson(r *http.Request, p interface{}) error {
	body := r.Body

	defer func(c io.ReadCloser) {
		err := c.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(body)

	decoder := json.NewDecoder(body)

	return decoder.Decode(&p)
}
