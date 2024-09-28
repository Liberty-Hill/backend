build: 
	docker build -t server:latest . 

run: 
	docker run -p 3001:3001 server
