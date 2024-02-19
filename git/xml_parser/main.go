package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)


// Структура Item JSON можно добавить еще поля, структура соответствует атрибутам XML
type Item struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
	Price           string `json:"price"`
	Author          string `json:"author"`
	URL             string `json:"url"`
	Image           string `json:"image"`
	PubDate         string `json:"pub_date"`
	Media           string `json:"media"`
}

func main() {
	// Здесь URL берем из консоли в процессе можно изменить
	var link string
	fmt.Scan(&link)

	doc, err := xmlquery.LoadURL(link)
	if err != nil {
		log.Fatalf("Error parsing XML: %q", err)
	}

	// Создаем map для хранения данных для JSON
	result := make(map[string]interface{})
	channelNode := xmlquery.FindOne(doc, "//channel")
	result["title"] = channelNode.SelectElement("title").InnerText()
	result["description"] = channelNode.SelectElement("description").InnerText()
	result["image"] = "" // Здесь нужно вставить URL изображения, если есть
	result["url"] = channelNode.SelectElement("link").InnerText()

	// Тут наполняем уже структуру Item, обработку доп.полей в процессе дописать
	items := make([]Item, 0)
	itemNodes := xmlquery.Find(doc, "//item")
	for _, itemNode := range itemNodes {
		item := Item{}
		item.Title = itemNode.SelectElement("title").InnerText()
		item.Description = strings.TrimSpace(itemNode.SelectElement("description").InnerText())
		item.DescriptionHTML = strings.TrimSpace(itemNode.SelectElement("content:encoded").InnerText())
		item.URL = itemNode.SelectElement("link").InnerText()
		// Здесь можно добавить обработку других полей, если они присутствуют в XML
		items = append(items, item)
	}

	result["items"] = items

	// Преобразовывем result в JSON
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %q", err)
	}

	// Здесь данные мы записываем в файл, в процессе изменить, тут это для проверки.
	file, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %q", err)
	}
	
	defer file.Close()

	// Запись JSON данных в файл
	_, err = file.Write(jsonResult)
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %q", err)
	}

	fmt.Println("Данные успешно записаны в файл output.json")
}
