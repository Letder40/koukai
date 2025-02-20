package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Urls struct {
    Id   int
    Urls string
    Published_at string
    Created_at string
    Updated_at string
}

func err_handle(err error) {
   if err != nil {
      fmt.Printf("Error: %s", err);
      os.Exit(1);
   } 
}

func main() {
   request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:1337/interesting-urls", nil);
   err_handle(err);
   response, err := http.DefaultClient.Do(request);
   err_handle(err);
   responseBody, err := io.ReadAll(response.Body);
   err_handle(err);

   fmt.Printf("status: %d\n", response.StatusCode);

   var urls []Urls;
   err = json.Unmarshal(responseBody, &urls);
   err_handle(err);

   for i:=0; i<len(urls); i++ {
      fmt.Printf("id: %d\nName: %s\n", urls[i].Id, urls[i].Urls);
   } 
}
