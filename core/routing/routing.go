package routing

import (
	"koukai/requests"
	"net/http"
)

func Routing(router *http.ServeMux) {
	var (
		appRuntime   AppRuntime
		publicServer Server
	)

    appRuntime.initRuntime()
	publicServer.InitGlobalServer()

	router.HandleFunc("POST /api/login", handleLogin)

	router.HandleFunc("POST /api/singup", handleSignup)

	router.HandleFunc("GET /api/user", appRuntime.userHandler)

	router.HandleFunc("GET /api/listen/server/public", publicServer.ListenChannel)

	router.HandleFunc("POST /api/write/server/public", publicServer.HandlePostMessage)

	fileServer := http.FileServer(http.Dir("./frontend/build"))
	router.Handle("GET /static/", fileServer)

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/build/index.html")
	})

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/build/index.html")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := requests.GetUserJWT(r)
		if err != nil {
			println(err.Error())
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		http.ServeFile(w, r, "./frontend/build/index.html")
	})
}
