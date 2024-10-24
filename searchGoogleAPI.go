package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const googleAPIURL = "https://www.googleapis.com/customsearch/v1"

type SearchResults struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

func main() {
	apiKey := ""
	cx := ""

	// Открываем файл для чтения
	file, err := os.Open("companies.txt")
	if err != nil {
		log.Printf("Ошибка открытия")
		return
	}

	defer file.Close()

	// Открываем файл для записи результатов в режиме добавления
	outputFile, err := os.OpenFile("Список Git.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Ошибка открытия файла для записи")
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Split(line, "\t")
		if len(fields) > 1 {
			companyName := fields[0]
			query := companyName + " site:github.com"
			escapedQuery := url.QueryEscape(query)

			url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s", googleAPIURL, apiKey, cx, escapedQuery)
			log.Printf("Запрос к Google API: %s", url)

			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("Ошибка при запросе к Google API: %v", err)
			}
			defer resp.Body.Close()

			log.Println("Ответ получен, статус:", resp.Status)

			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Неуспешный HTTP статус: %d %s", resp.StatusCode, resp.Status)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Ошибка при чтении тела ответа: %v", err)
			}

			// log.Println("Ответ от API:", string(body)) // Выводим тело ответа для отладки

			log.Println("Декодирование JSON...")

			var results SearchResults
			if err := json.Unmarshal(body, &results); err != nil {
				log.Fatalf("Ошибка декодирования JSON: %v", err)
			}

			if len(results.Items) > 0 {
				for _, item := range results.Items {
					fmt.Println("Найдено ссылка на GitHub:", item.Link)
					writer.WriteString(companyName + "\t" + item.Link + "\n")
				}
			} else {
				fmt.Println("GitHub профили не найдены с этой ссылкой.")
				writer.WriteString(companyName + "\tGitHub профили не найдены с этой ссылкой.\n")
			}
		}
	}

	writer.Flush()
}
