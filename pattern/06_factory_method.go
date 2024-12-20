package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
// Тип источника данных
type SourceType string

// Интерфейс (Product) для описания источника данных с методом для получения данных
type Source interface {
	GetData() string
}

// Источник для получения данных из MySQL (ConcreteProduct)
type MySqlSource struct{}

func (s *MySqlSource) GetData() string {
	return fmt.Sprintln("Получение данных из базы данных MySQL")
}

// Источник для получения данных из PostgreSQL (ConcreteProduct)
type PostgresSource struct{}

func (s *PostgresSource) GetData() string {
	return fmt.Sprintln("Получение данных из базы данных PostgreSQL")
}

// Источник для получения данных из Excel (ConcreteProduct)
type ExcelSource struct{}

func (s *ExcelSource) GetData() string {
	return fmt.Sprintln("Получение данных из Excel файла")
}

// Источник для получения данных из CSV (ConcreteProduct)
type CsvSource struct{}

func (s *CsvSource) GetData() string {
	return fmt.Sprintln("Получение данных из CSV файла")
}

// Создатель (Creator) с фабричным методом для создания источников данных
type SourceFactory interface {
	CreateSource(sourceType SourceType) (Source, error)
}

// Конкретный создатель (ConcreteCreator) для создания источников данных из файлов
type FileSourceFactory struct{}

func (factory *FileSourceFactory) CreateSource(sourceType SourceType) (Source, error) {
	switch sourceType {
	case "CSV":
		return &CsvSource{}, nil
	case "EXCEL":
		return &ExcelSource{}, nil
	}
	return nil, errors.New("Non supported type")
}

// Конкретный создатель (ConcreteCreator) для создания источников данных из баз данных
type DBSourceFactory struct{}

func (factory *DBSourceFactory) CreateSource(sourceType SourceType) (Source, error) {
	switch sourceType {
	case "MYSQL":
		return &MySqlSource{}, nil
	case "POSTGRES":
		return &PostgresSource{}, nil
	}
	return nil, errors.New("Non supported type")
}

func GetData(factory SourceFactory, sourceType SourceType) {
	source, err := factory.CreateSource(sourceType)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(source.GetData())
}

func TestFactoryMethod() {
	fileFactory := &FileSourceFactory{}
	GetData(fileFactory, "EXCEL")
	GetData(fileFactory, "CSV")

	dbFactory := &DBSourceFactory{}
	GetData(dbFactory, "MYSQL")
	GetData(dbFactory, "POSTGRES")
}
