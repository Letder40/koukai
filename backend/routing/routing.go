package routing

import (
	"fmt"
	"koukai/requests"
	"net/http"
)

func Routing(router *http.ServeMux) {
    router.Handle("/", http.FileServer(http.Dir("./webfiles")))

    router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        resp, err := requests.RequestStrapi("users")
        if err != nil {
            fmt.Println("Error in strapi request, error:", err)
        }
        w.Header().Add("Content-Type", "application/json")
        fmt.Fprint(w, resp)
    })
}
