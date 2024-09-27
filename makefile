build: 
	docker build -t server:latest . 

run: 
	docker run -p 80:80 --name server server
