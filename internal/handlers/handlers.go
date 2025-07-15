package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// handleMain returns HTML from a file index.html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fileName := "index.html"

	data, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("parsing multipart form")
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Printf("Parse form error: %v\n", err)
		http.Error(w, fmt.Sprintf("failed to parse form: %v", err), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Form file error: %v\n", err)
		http.Error(w, "error when receiving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Read file error: %v\n", err)
		http.Error(w, fmt.Sprintf("failed to read file: %v", err), http.StatusInternalServerError)
		return
	}

	result := service.TextDifinition(string(data))
	fmt.Printf("TextDifinition result: %s\n", result)

	fileName := fmt.Sprintf("%s.txt", time.Now().UTC().Format("02.01.06 15.4.5"))
	fmt.Printf("Saving to: %s\n", fileName)

	if err := os.WriteFile(fileName, []byte(result), 0644); err != nil {
		fmt.Printf("Write file error: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to write file: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
