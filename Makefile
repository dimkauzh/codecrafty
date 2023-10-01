GO = go
MAIN_FILE = src/main.go
NAME = codecrafty
FYNE = fyne
WEB_TARGET = wasm

.PHONY: run build build_run

run:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running $(NAME)...                 |"
	@echo " ----------------------------------------------------"
	@echo

	@$(GO) run $(MAIN_FILE)

build:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                |"
	@echo " ----------------------------------------------------"
	@echo

	@$(GO) build -o bin/$(NAME) $(MAIN_FILE)

setup:
	@echo
	@echo
	@echo " ----------------------------------------------------"
	@echo "      Installing Packages for Dev version...      |"
	@echo " ----------------------------------------------------"
	@echo
	$(GO) mod tidy

build_run:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                |"
	@echo " ----------------------------------------------------"
	@echo

	@$(GO) build -o bin/$(NAME) $(MAIN_FILE)

	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running $(NAME)...                 |"
	@echo " ----------------------------------------------------"
	@echo

	@./bin/$(NAME)

web:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|            Running Web version...              |"
	@echo " ----------------------------------------------------"
	@echo

	@$(FYNE) serve --sourceDir src --icon ../assets/icon/Icon.png --target $(WEB_TARGET)
