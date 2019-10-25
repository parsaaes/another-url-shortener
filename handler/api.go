package handler

import (
	"encoding/json"
	"github.com/catinello/base62"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gitlab.com/parsaaes/another-url-shortener/model"
	"net/http"
	"strings"
)

type FullURLHandler struct {
	UrlRepo model.SQLUrlRepo
}

func NewFullURLHandler(db *gorm.DB) FullURLHandler {
	return FullURLHandler{
		UrlRepo: model.NewSQLUrlRepo(db),
	}
}

func (f FullURLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["url"]

	if !f.validateQuery(query) {
		f.jsonResponse(w, "message", "invalid request syntax", http.StatusBadRequest)
		logrus.Println("bad query param:", query)
		return
	}

	url := query[0]
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	id, err := f.UrlRepo.Save(model.Url{
		Url: url,
	})
	if err != nil {
		f.jsonResponse(w, "message", "database error", http.StatusInternalServerError)
		logrus.Println("can't save to database: ", err)
		return
	}

	shortUrl := r.Host + "/" + base62.Encode(id)
	logrus.Println("short url generated: ", shortUrl)

	f.jsonResponse(w, "short", shortUrl, http.StatusOK)
}

func (f FullURLHandler) jsonResponse(w http.ResponseWriter, key string, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]string{key: message})
	if err != nil {
		logrus.Println(err)
	}
}

func (f FullURLHandler) validateQuery(query []string) bool {
	return !(query == nil || len(query) < 1 || len(query[0]) < 1)
}
