#!make
include .env

build: 
	docker build -t server:latest . 

run: 
	docker run -p ${PORT}:${PORT} server
