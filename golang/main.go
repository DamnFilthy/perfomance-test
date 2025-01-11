package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "strings"
  "time"
)

const USERS = 100_000_000

type User struct {
  ID      int
  Content string
}

func (u User) MarshalJSON() ([]byte, error) {
  var s strings.Builder
    s.WriteString("User")
    s.WriteString(strconv.FormatInt(int64(u.ID), 10))

    return json.Marshal(struct {
        ID      int    `json:"id"`
      Name    string `json:"name"`
      Content string `json:"content"`
    }{
        ID: u.ID,
        Name: s.String(),
        Content:  u.Content,
    })
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  start := time.Now()

  result := make([]User, USERS)

  for i := 0; i < USERS; i++ {
    result[i] = User{
      ID:      i,
      Content: "<h1>Hello, World!</h1>",
    }
  }

  // Фильтрация нечетных ID
  filteredResult := make([]User, USERS/2)

  for i := 0; i < USERS; i++ {
    if i%2 != 0 {
      filteredResult[i/2] = result[i]
    }
  }

  // Выбор элемента из середины списка
  midIndex := len(filteredResult) / 2

  selectedUser := filteredResult[midIndex]

  b, _ := json.Marshal(selectedUser)

  // Сериализация в JSON
  w.Header().Set("Content-Type", "application/json")
  w.Write(b)

  duration := time.Since(start)
  fmt.Println("Total duration:", duration)
}

func main() {
  http.HandleFunc("/test", testHandler)
  fmt.Println("Server is running on http://localhost:3001")
  http.ListenAndServe(":3001", nil)
}