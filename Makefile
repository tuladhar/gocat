#
# Makefile - `gocat` (https://github.com/tuladhar/gocat)
# 
PROGRAM_NAME:=gocat

build:
	go build -o $(PROGRAM_NAME) src/main.go

cleanup:
	rm $(PROGRAM_NAME)
