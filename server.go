package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "encoding/json"
)

func main() {
  port := os.Getenv("PORT")
  secret := os.Getenv("SECRET")

  if port == "" {
    port = "8080"
  }

  log.Printf("Booting on port %s", port)

  if secret == "" {
    secret = "secret"
    log.Printf("SECRET not given, using \"%s\"", secret)
  }

  http.HandleFunc("/", http.NotFound)
  http.HandleFunc("/temperatures", Temperatures)
  http.HandleFunc("/" + secret + "/temperature", StoreTemperature)

  log.Fatal(http.ListenAndServe(":" + port, nil))
}

func WriteJSON(content interface{}, writer http.ResponseWriter) {
  writer.Header().Set("Content-Type", "application/json")
  string, error := json.Marshal(content)

  if error != nil {
		fmt.Println("error:", error)
	}

  writer.Write(string)
}

func StoreTemperature(writer http.ResponseWriter, request *http.Request) {
  if(request.Method != "POST") {
    http.NotFound(writer, request)
    return
  }

  WriteJSON("OK", writer)
}

func Temperatures(writer http.ResponseWriter, request *http.Request) {
  values := make([]string, 0)
  WriteJSON(values, writer)
}
