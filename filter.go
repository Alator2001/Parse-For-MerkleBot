package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Функция для фильтрации репозиториев
func filterCompanyProfiles(urls map[string][]string) map[string]string {
	companyProfiles := make(map[string]string)
	// Регулярное выражение для поиска ссылок на профили компаний (не issues и репозиториев)
	profileRegex := regexp.MustCompile(`^https://github\.com/[^/]+$`)

	for company, companyUrls := range urls {
		for _, url := range companyUrls {
			// Если найден профиль компании, сохраняем его
			if profileRegex.MatchString(url) {
				companyProfiles[company] = url
				break
			}
		}
	}
	return companyProfiles
}

func main() {
	// Открытие файла для чтения
	file, err := os.Open("C:/Users/Ghulqul/Downloads/Список Git.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Чтение файла и группировка ссылок по компаниям
	urlsByCompany := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Разделение строки по табуляции или пробелам
		parts := strings.Fields(line)
		if len(parts) == 2 {
			company := parts[0]
			url := parts[1]
			urlsByCompany[company] = append(urlsByCompany[company], url)
		}
	}

	// Проверка на ошибки при чтении файла
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	// Фильтрация профилей компаний
	companyProfiles := filterCompanyProfiles(urlsByCompany)

	// Сортировка по названиям компаний
	var companyNames []string
	for company := range companyProfiles {
		companyNames = append(companyNames, company)
	}
	sort.Strings(companyNames)

	// Запись результатов в файл
	outputFile, err := os.Create("company_profiles.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	fmt.Fprintln(writer, "Основные профили компаний:")
	for _, company := range companyNames {
		fmt.Fprintf(writer, "%s\t%s\n", company, companyProfiles[company])
	}
	writer.Flush()

	fmt.Println("Результаты записаны в файл company_profiles.txt")
}
