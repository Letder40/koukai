package routing

import (
	"koukai/requests"
	"net/http"
	"strings"
)

func Routing(router *http.ServeMux) {

	router.HandleFunc("POST /api/login", handleLogin)

	router.HandleFunc("POST /api/singup", handleSingup)

	var appRuntime AppRuntime
	router.HandleFunc("GET /api/user", appRuntime.userHandler)

	fileServer := http.FileServer(http.Dir("./frontend/build"))
	router.Handle("GET /static/", fileServer)

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, "./frontend/build/index.html")
	})

	router.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, "./frontend/build/index.html")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := requests.GetUserJWT(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}

		if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, "./frontend/build/index.html")
	})
}
