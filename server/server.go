package server

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/parsaaes/another-url-shortener/config"
	"gitlab.com/parsaaes/another-url-shortener/database"
	"gitlab.com/parsaaes/another-url-shortener/handler"
	"net/http"
	"strconv"
)

func StartServer() {
	config.Init(".")

	db, err := database.CreatePostgresDB()

	if err != nil {
		logrus.Fatal("cannot create db connection: ", err)
	}

	fullURLHandler := handler.NewFullURLHandler(db)
	shortURLHandler := handler.NewShortURLHandler(db)

	http.HandleFunc("/api", fullURLHandler.Shorten)
	http.HandleFunc("/", shortURLHandler.Redirect)

	logrus.Println("Server started...")
	logrus.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Cfg.Port), nil))
}

func SetupFront() {
	http.HandleFunc("/url-shortener", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "front/index.html")
	})
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("front/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("front/css"))))
}
