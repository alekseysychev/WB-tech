/*
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

type linksStruct struct {
	target     string              // куда сохранять всё
	scheme     string              // http/https
	host       string              // хост разрешенный для скачивания
	downloaded map[string]struct{} // все скачанные файлы
	sync.RWMutex
}

func main() {
	// var link string
	// fmt.Scan(&link)
	// link = "https://webzlodimir.ru/works"
	link := "https://example.com"

	// обрабатываем ссылку
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	links := &linksStruct{
		target:     "./site/",
		scheme:     u.Scheme,
		host:       u.Hostname(),              // задаём хост в рамках которого мы можем парсить ссылки
		downloaded: make(map[string]struct{}), // создаём пустой список уже скачанных
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go links.download(&wg, link)
	wg.Wait()
	fmt.Println("done")
}

func (l *linksStruct) download(wg *sync.WaitGroup, link string) {
	fmt.Println("download -> ", link)

	l.Lock()
	_, ok := l.downloaded[link]
	if ok {
		l.Unlock()
		return
	}
	// добавляем ссылку в обработанные
	l.downloaded[link] = struct{}{}
	l.Unlock()

	// получаем по ссылке
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	// читаем из потока в переменную
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// закрываем ридер
	resp.Body.Close()

	// обрабатываем ссылку
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println(err)
	}
	// путь относительно хоста
	filePath := u.Path
	// убираем слеши
	filePathTrim := strings.Trim(filePath, "/")
	// разделяем на куски по слешам
	filePathTrimSplit := strings.Split(filePathTrim, "/")
	// если у последней части ссылки нет точки, делаем предположение что это директория
	if !strings.Contains(filePathTrimSplit[len(filePathTrimSplit)-1], ".") {
		filePathTrimSplit = append(filePathTrimSplit, "index.html")
	}
	// кусок пути без последнего элемента для создания папки
	filePathJoinDir := l.target + strings.Join(filePathTrimSplit[:len(filePathTrimSplit)-1], "/")
	// кусок пути до файла
	filePathJoinFile := l.target + strings.Join(filePathTrimSplit, "/")
	// создаём папки
	os.MkdirAll(filePathJoinDir, 0777)
	// создаём файл
	out, err := os.Create(filePathJoinFile)
	if err != nil {
		fmt.Println(err)
	}
	// копируем содержимое файла
	_, err = out.Write(body)
	// _, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// закрываем файл
	out.Close()
	// если страничка html, парсим её на наличие ссылок
	contentType := resp.Header.Get("content-type")
	if len(contentType) > 10 && contentType[:10] == "text/html;" {
		links := l.parseHtmlForLinks(body)
		for _, link := range links {
			wg.Add(1)
			go l.download(wg, link)
		}
	}
	wg.Done()
}

func (l *linksStruct) parseHtmlForLinks(body []byte) []string {
	var result []string
	hrefLinksRegexp := regexp.MustCompile(`(?mis)href="(?P<link>.*?)"`)
	// ищем все ссылки
	matches := hrefLinksRegexp.FindAllSubmatch(body, -1)
	// проходим по ссылкам
	for _, s := range matches {
		// преобразуем из байтовой группы в строку
		link := string(s[1])
		// обрабатываем ссылку
		u, err := url.Parse(link)
		if err != nil {
			fmt.Println(err)
		}
		// если хост не задан, скорее всего ссылка относительная
		if u.Hostname() == "" {
			// если ссылка не относительная, а например для перехода или почта
			if len(link) > 0 && link[0] != '/' {
				continue
			}
			link = l.scheme + "://" + l.host + link
		} else
		// если хост задан, но не соответствует оригиналу (ведет на другой сайт)
		if u.Hostname() != l.host {
			continue
		} else
		// хост задан, но ссылка относительная относительно протокола
		if len(link) > 2 && link[:2] == "//" {
			link = l.scheme + ":" + link
		}

		// проверяем не спарсена ли ещё ссылка
		l.RLock()
		_, ok := l.downloaded[link]
		// если спарсена, пропускаем
		if ok {
			l.RUnlock()
			continue
		}
		l.RUnlock()

		// добавляем в список, который надо ещё спарсить
		result = append(result, link)
	}
	return result
}
