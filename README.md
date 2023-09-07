# LibraryServise

## База данных

Для создания docker image базы данных использовать команду

```bash
docker build -t library_db .
```

Запуск контейнера производится командой

```bash
docker run --name mysql_library -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=library library_db
```

## Запуск сервера

Для того чтобы скомпилировать исполняющий файл выполнить команду

```bash
go build cmd/main.go
```

Затем просто выполните исполняемый файл

```bash
./main
```

## Конфигурации

По дефолту конфигурация берется из файла config.yaml в той же папке что и исполняемый файл. 

Для того чтобы использовать свой конфигурационный файл при запуске сервера использовать флаг config

### Пример

```bash
./main --config /home/config.yaml
```