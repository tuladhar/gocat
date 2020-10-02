#
# Makefile - `gocat` (https://github.com/tuladhar/gocat)
# 
PROGRAM_NAME=gocat
INSTALL_PATH=/usr/local/bin

build:
	go build -o $(PROGRAM_NAME) src/main.go

install:
	go build -o $(PROGRAM_NAME) src/main.go
	@sudo install $(PROGRAM_NAME) $(INSTALL_PATH)/$(PROGRAM_NAME) 
	echo "$(PROGRAM_NAME): installed at $(INSTALL_PATH)" 

clean:
	rm $(PROGRAM_NAME)

uninstall:
	rm $(PROGRAM_NAME)
	@sudo rm -f $(INSTALL_PATH)/$(PROGRAM_NAME)
	@echo "$(PROGRAM_NAME): uninstalled"
	
