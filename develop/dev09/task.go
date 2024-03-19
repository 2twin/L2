package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Использование: ./wget <URL>")
        return
    }

    url := os.Args[1]
    if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
        url = "http://" + url
    }

    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ошибка при получении страницы: %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Fprintf(os.Stderr, "ошибка: код состояния %d\n", resp.StatusCode)
        return
    }

    // Получаем имя файла из URL и сохраняем в текущей директории
    fileName := getFileName(url)
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ошибка при создании файла: %v\n", err)
        return
    }
    defer file.Close()

    // Копируем содержимое ответа в файл
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ошибка при копировании данных в файл: %v\n", err)
        return
    }

    fmt.Printf("Скачивание завершено. Содержимое сохранено в файл %s\n", fileName)
}

// Функция для получения имени файла из URL
func getFileName(url string) string {
    parts := strings.Split(url, "/")
    fileName := parts[len(parts)-1]
    if fileName == "" {
        fileName = "index.html"
    }
    return fileName
}