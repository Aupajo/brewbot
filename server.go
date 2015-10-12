package main

import (
  "fmt"
  "html"
  "log"
  "os"
  "net/http"
)

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

  log.Printf("Booting on port %s", port)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":" + port, nil))
}
