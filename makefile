all: database build 

database:
	docker build -t library_db .
	docker run --name mysql_library -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=library library_db
	
build:
	go build cmd/main.go
	
rm:
	docker stop mysql_library
	docker rm mysql_library
	docker rmi library_db