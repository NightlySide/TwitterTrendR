.PHONY: server

.ONESHELL:
server:
	@echo "-- Starting server"
	go run backend/*
