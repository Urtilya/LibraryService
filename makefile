all: build database

database:
	docker run --name mysql_test_library -d -v $(PWD)/data_db:/var/lib/mysql -p 3606:3606  mysql
	
build:
	go build cmd/main.go
	