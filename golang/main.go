package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// var result []User
	result:=make([]User, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		result[i] = User {
			ID:      i,
			Name:    fmt.Sprintf("User%d", i),
			Content: "<h1>Hello, World!</h1>",
		}
	}

	// Фильтрация нечетных ID
	filteredResult := make([]User, 5_000_000)

	for i := 0; i < 10_000_000; i++ {
		if i%2 != 0 {
			filteredResult[i/2] = result[i]
		}
	}

	// Выбор элемента из середины списка
	midIndex := len(filteredResult) / 2

	selectedUser := filteredResult[midIndex]

	// Сериализация в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selectedUser)

	duration := time.Since(start)
	fmt.Println("Total duration:", duration)
}

func main() {
	http.HandleFunc("/test", testHandler)
	fmt.Println("Server is running on http://localhost:3001")
	http.ListenAndServe(":3001", nil)
}
