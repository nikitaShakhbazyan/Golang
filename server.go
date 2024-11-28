package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// Структура для хранения данных
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	db   *sql.DB
	mu   sync.Mutex 
	wg   sync.WaitGroup
)

// Инициализация базы данных
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы, если она не существует
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Обработчик для получения всех записей
func getItems(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	rows, err := db.Query("SELECT id, name FROM items")
	mu.Unlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Обработчик для добавления новой записи
func addItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	query := "INSERT INTO items (name) VALUES (?)"
	result, err := db.Exec(query, item.Name)
	mu.Unlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func handleRequests() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getItems(w, r)
		} else if r.Method == http.MethodPost {
			addItem(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Инициализация базы данных
	initDB()
	defer db.Close()

	// Запуск HTTP-сервера в отдельной горутине
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleRequests()
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}
