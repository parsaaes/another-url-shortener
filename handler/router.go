package handler

import (
	"encoding/json"
	"github.com/catinello/base62"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gitlab.com/parsaaes/another-url-shortener/model"
	"net/http"
	"strings"
	"time"
)

type ShortURLHandler struct {
	UrlRepo model.SQLUrlRepo
}

func NewShortURLHandler(db *gorm.DB) ShortURLHandler {
	return ShortURLHandler{
		UrlRepo: model.NewSQLUrlRepo(db),
	}
}

func (s ShortURLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Path
	if !s.validateShort(short) {
		s.jsonResponse(w, "invalid short url request", http.StatusBadRequest)
		logrus.Println("invalid short url request: ", short)
		return
	}

	short = strings.TrimPrefix(short, "/")

	// check cookies
	cookie, err := r.Cookie(short)
	if err != nil {
		logrus.Println(err)
	} else {
		logrus.Println("used cookies to redirect")
		http.Redirect(w, r, cookie.Value, http.StatusMovedPermanently)
		return
	}

	id, err := base62.Decode(short)
	if err != nil {
		s.jsonResponse(w, "can't decode short url", http.StatusInternalServerError)
		logrus.Println("can't decode short url: ", err)
		return
	}

	url, err := s.UrlRepo.Find(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.jsonResponse(w, "not found", http.StatusNotFound)
			logrus.Println("not found ", err)
			return
		}
		s.jsonResponse(w, "database error", http.StatusInternalServerError)
		logrus.Println("database error: ", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    short,
		Value:   url.Url,
		Expires: time.Now().AddDate(0, 1, 0),
	})
	http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
}

func (s ShortURLHandler) jsonResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]string{"message": message})
	if err != nil {
		logrus.Println(err)
	}
}

func (s ShortURLHandler) validateShort(short string) bool {
	return !(short == "")
}
