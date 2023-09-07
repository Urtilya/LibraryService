package main

import (
	"LibraryService/config"
	"LibraryService/run"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "config.yaml", "path to config file")
	flag.Parse()
	//Инициализация конфигурацииы
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cfg := &config.Config{}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		panic(err)
	}
	//Инициализация приложения
	app := run.NewApp(cfg)
	app.Boot()
	//Запуск приложения
	app.Run()
}
