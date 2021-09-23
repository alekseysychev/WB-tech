package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import "errors"

var ErrNameNotFound = errors.New("name not found")

var db = map[string]string{
	"a@a.com": "a",
	"b@b.com": "b",
}

type database struct{}

func (self *database) getNameByMail(mail string) (string, error) {
	name, ok := db[mail]
	if !ok {
		return "", ErrNameNotFound
	}
	return name, nil
}

type mdWriter struct{}

func (self *mdWriter) title(title string) string {
	return "# Welcome to " + title + "'s page!"
}

type PageMaker struct{}

func (self *PageMaker) MakeWelcomePage(mail string) (string, error) {
	database := database{}
	writer := mdWriter{}

	name, err := database.getNameByMail(mail)
	if err != nil {
		return "", err
	}
	page := writer.title(name)

	return page, nil
}
