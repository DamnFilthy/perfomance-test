package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "strings"
  "time"
)

const USERS = 10_000_000

type User struct {
  ID      int    `json:"id"`
  Name    string `json:"name"`
  Content string `json:"content"`
}

func (u *User) getUser() *User {
  var s strings.Builder
  s.WriteString("User")
  s.WriteString(strconv.FormatInt(int64(u.ID), 10))
  u.Name = s.String()
  return u
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

  // Сериализация в JSON
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(selectedUser.getUser())

  duration := time.Since(start)
  fmt.Println("Total duration:", duration)
}

func main() {
  http.HandleFunc("/test", testHandler)
  fmt.Println("Server is running on http://localhost:3001")
  http.ListenAndServe(":3001", nil)
}