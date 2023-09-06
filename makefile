database:
	- docker run -d -v $(PWD)/data_db:/var/lib/mysql -p 3606:3606 -d mysql