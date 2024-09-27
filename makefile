
build: 
	docker build -t server:latest . 

run: 
	docker run -p 3000:3000 server
