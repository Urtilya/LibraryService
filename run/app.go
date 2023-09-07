package run

import (
	"LibraryService/config"
	"LibraryService/internal/api"
	"LibraryService/internal/db"
	"LibraryService/internal/handles"
	"LibraryService/internal/service"
	"LibraryService/internal/storage"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type App struct {
	cfg    *config.Config
	server *grpc.Server
	lis    net.Listener
	sig    chan os.Signal
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
		sig: make(chan os.Signal, 1),
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		//Ожидание сигнала на kill
		<-a.sig
		cancel()
	}()
	//Запуск сервера
	a.server.Serve(a.lis)
	//Плавныая остановка сервера
	<-ctx.Done()
	a.server.GracefulStop()
}

func (a *App) Boot() {
	//Подключение к БД
	database, err := db.NewSqlDB(a.cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	//Создание оббертки над базой данных
	storage := storage.NewStorage(database)
	//Создание обработчика данных
	service := service.NewService(storage)
	//Создание обработчика запросов
	handles := handles.NewLibraryRPC(service)
	//Создание grpc сервера
	a.server = grpc.NewServer()
	//Регистрация обработчика
	api.RegisterLibraryServer(a.server, handles)
	//Создание слушателя
	a.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.GrpcPort))
	if err != nil {
		log.Fatal(err)
	}
}
