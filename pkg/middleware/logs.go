package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Loggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.SetFormatter(&log.JSONFormatter{})
		log.WithFields(log.Fields{
			"method":     r.Header,
			"URL":        r.URL.Path,
			"StatusCode": wrapper.StatusCode,
		}).Info("Information about request and resposnse")
	})
}
