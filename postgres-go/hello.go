package main

import (
  "database/sql"
  "fmt"
  "net/http"


  _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
  db = getDatabase()
  defer db.Close()  

    http.HandleFunc("/", handler)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("ListenAndServe exited with error %v\n", err)
    }
}

func getDatabase() *sql.DB {
  host := "postgres-svc"
  port := 5432
  user := "postgres"
  dbname := "hello"
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    fmt.Printf("Failed to open database: %v\n", err)
    panic(err)
  }
  fmt.Println("Opened database")

  return db  
}

func incrementCount() int {
    var hits int
    row := db.QueryRow("UPDATE hits SET total = total + 1 RETURNING total;")
    err := row.Scan(&hits)
    if err != nil {
        fmt.Printf("Database increment error: %v\n", err)
        return -1
    }
    return hits
}

func handler(w http.ResponseWriter, r *http.Request) {
    hits := incrementCount()
    response := fmt.Sprintf("hits: %d\n", hits)
    w.Write([]byte(response))
}

