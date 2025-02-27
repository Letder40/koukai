package routing

import (
	"fmt"
	"koukai/requests"
	"net/http"
	"strings"
)

type UserData struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLoginData struct {
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}

func Routing(router *http.ServeMux) {
    router.HandleFunc("POST /api/login", HandleLogin)

    router.HandleFunc("POST /api/singup", HandleSingup)


    router.HandleFunc("GET /api/user", func(w http.ResponseWriter, r *http.Request) {
        jwt, err := requests.GetUserJWT(w, r);
        if err != nil {
            return;
        }

        resp, err := requests.UserStrapiRequest("users/me", jwt)
        if err != nil {
            fmt.Println("Error in strapi request, error:", err)
        }

        w.Header().Add("Content-Type", "application/json")
        fmt.Fprint(w, resp)
    })

    fileServer := http.FileServer(http.Dir("./frontend/build"))
    router.Handle("GET /static/", fileServer)

    router.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

        http.ServeFile(w, r, "./frontend/build/index.html")
    })

    router.HandleFunc("/singup", func (w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

        http.ServeFile(w, r, "./frontend/build/index.html")
    })

    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        _, err := requests.GetUserJWT(w, r);
        if err != nil {
            return;
        }

        if strings.Contains(r.URL.Path, ".") {
			fileServer.ServeHTTP(w, r)
			return
		}

        http.ServeFile(w, r, "./frontend/build/index.html")
    })
}
