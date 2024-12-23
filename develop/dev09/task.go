package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура с опциями программы
type WgetOptions struct {
	// Файл для вывода ответа запроса
	OutputFile string
	// Вывод деталей ответа
	ServerResponse bool
	// Скачать весь сайт
	Mirror bool
}

// Функция для парсинга аргументов программы
func parseArgs() (*WgetOptions, string) {
	var options WgetOptions
	var url string

	flag.StringVar(&options.OutputFile, "O", "", "Output file")
	flag.BoolVar(&options.ServerResponse, "S", false, "Server response")
	flag.BoolVar(&options.Mirror, "m", false, "Mirror mode")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Пожалуйста, укажите URL")
	}
	url = args[0]

	if options.OutputFile == "" {
		filename := path.Base(url)
		matched, err := regexp.MatchString(".+\\..+", filename)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			options.OutputFile = filename
		} else if options.Mirror {
			options.OutputFile = "index.html"
		} else {
			options.OutputFile = strconv.Itoa(rand.Intn(100000))
		}
	}

	return &options, url
}

// Функция для скачивания файла
func downloadFile(url string, options *WgetOptions) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка при загрузке файла: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Ошибка в ответе URL: %s\n", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении тела ответа: ", err)
	}

	file, err := os.Create(options.OutputFile)
	if err != nil {
		log.Fatal("Ошибка при создании файла: ", err)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		log.Fatal("Ошибка при записи в файл: ", err)
	}
	return resp
}

// Функция для вывода информации о заголовках
func printServerHeaders(resp *http.Response) {
	for key, value := range resp.Header {
		fmt.Printf("%s=%s\n", key, strings.Join(value, ","))
	}
}

// Функиця для скачивания всего сайта
func downloadMirror(baseURL string) {
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в отправке запроса URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Ошибка в ответе URL: %s\n", resp.Status)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в парсинге HTML: %v\n", err)
		os.Exit(1)
	}

	downloadPage(baseURL, "index.html", doc)
}

func downloadPage(baseURL, filename string, doc *html.Node) {
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в обработки url: %v\n", err)
		os.Exit(1)
	}

	// Сайт скачивается в директорию с именем хоста в URL
	host := parsedUrl.Hostname()
	err = os.Mkdir(host, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в создании директории: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(path.Join(host, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в создании файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var buf strings.Builder
	var resources []string
	var walk func(*html.Node)

	// Функция для рекурсивного обхода DOM дерева для получения ссылок на ресурсы
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "a" || n.Data == "img" || n.Data == "link" || n.Data == "script" {
				for _, attr := range n.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						resources = append(resources, resolveURL(baseURL, attr.Val))
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	err = html.Render(&buf, doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в рендеринге HTML: %v\n", err)
		os.Exit(1)
	}

	_, err = file.WriteString(buf.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в записи файла: %v\n", err)
		os.Exit(1)
	}

	for _, res := range resources {
		downloadResource(host, res)
	}
}

// Функция для скачивания ресурса
func downloadResource(dir, resourceURL string) {
	resp, err := http.Get(resourceURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в скачивании ресурса %s: %v\n", resourceURL, err)
		return
	}
	defer resp.Body.Close()

	filename := getFileNameFromURL(resourceURL)
	file, err := os.Create(path.Join(dir, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в создании файла %s: %v\n", resourceURL, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в сохранении ресурса %s: %v\n", resourceURL, err)
	}
}

// Функиця для получения
func getFileNameFromURL(url string) string {
	parsed := strings.TrimSuffix(filepath.Base(url), "/")
	if parsed == "" || strings.Contains(parsed, "?") {
		return "index.html"
	}
	return parsed
}

// Функиця для получения ссылка на ресурс
func resolveURL(base, relative string) string {
	// Если ресурс имеет полную ссылку то возвращаем ее, иначе на основе базового URL
	if strings.HasPrefix(relative, "http") || strings.HasPrefix(relative, "https") {
		return relative
	}
	baseDir := path.Dir(base)
	return fmt.Sprintf("%s/%s", strings.TrimSuffix(baseDir, "/"), strings.TrimPrefix(relative, "/"))
}

func main() {
	options, url := parseArgs()

	if options.Mirror {
		downloadMirror(url)
	} else {
		resp := downloadFile(url, options)

		if options.ServerResponse {
			printServerHeaders(resp)
		}
	}
}
