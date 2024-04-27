package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"http-panic-handler/domain/web"
	"http-panic-handler/repository"
	"io"
	"log"
	"net/http"
)

func PanicRecovery(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {

				var errResponse map[string]string
				var errCodeStatus int
				var logStdOut web.Log

				switch {
				case errors.Is(err.(error), repository.ErrUserNotFound):
					errResponse = map[string]string{"error": err.(error).Error()}
					errCodeStatus = http.StatusNotFound

					if r.Method == http.MethodPost {
						body, err := io.ReadAll(r.Body)
						if err != nil {
							log.Println("Error reading request body:", err)
						}
						logStdOut.Body = string(body)

						// reset request body
						r.Body = io.NopCloser(bytes.NewBuffer(body))
					}

				default:
					errResponse = map[string]string{"error": "Internal Server Error"}
					errCodeStatus = http.StatusInternalServerError
				}

				logStdOut.Method = r.Method
				logStdOut.Url = r.Host
				logStdOut.Path = r.URL.Path
				jsonLogBody, _ := json.Marshal(&logStdOut)
				log.Printf("%v\n", string(jsonLogBody))

				jsonBody, _ := json.Marshal(errResponse)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(errCodeStatus)
				w.Write(jsonBody)
			}
		}()

		n.ServeHTTP(w, r)
	})
}
