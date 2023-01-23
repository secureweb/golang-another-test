package main

import (
	"log"
	"net/http"

	structure "github.com/secureweb/golang-test/pkg"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tpl, err := structure.Test()
		if err != nil {
			log.Printf("page %s not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"userAgent": r.UserAgent(),
		}
		if err := tpl.Execute(w, data); err != nil {
			return
		}
	})
	server.ListenAndServe()
	log.Println("server started...")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
