package routing

import (
	"fmt"
	"koukai/requests"
	"net/http"
	"strings"
)

func Routing(router *http.ServeMux) {
    fileServer := http.FileServer(http.Dir("./frontend/build"))
    router.Handle("GET /static/", fileServer)

    router.HandleFunc("GET /api/user", func(w http.ResponseWriter, r *http.Request) {
        jwt, err := requests.GetUserJWT(r);
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusMovedPermanently)
            fmt.Fprint(w, "")
            return;
        }
        resp, err := requests.UserStrapiRequest("users/me", jwt)
        if err != nil {
            fmt.Println("Error in strapi request, error:", err)
        }
        w.Header().Add("Content-Type", "application/json")
        fmt.Fprint(w, resp)
    })

    router.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {   

    })

    router.HandleFunc("POST /api/singup", func(w http.ResponseWriter, r *http.Request) {

    })

    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

        http.ServeFile(w, r, "./frontend/build/index.html")
    })

}
