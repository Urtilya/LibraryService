package main

import (
	"LibraryService/config"
	"LibraryService/run"
)

func main() {
	//Инициализация конфигурацииы
	cfg := &config.Config{}
	//Инициализация приложения
	app := run.NewApp(cfg)
	app.Boot()
	//Запуск приложения
	app.Run()
}
