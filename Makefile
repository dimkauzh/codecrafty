COMPILER = g++
LIB_FLAGS = -Ilibraries/raylib/src/ -Llibraries/raylib/src/ -lraylib
BUILD_FLAGS = -o build/codecrafty src/main.cpp
OPTIMIZE_FLAGS = -O2
CPP_FLAGS = -std=gnu++2b

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S), Darwin)
	LIB_FLAGS += -framework OpenGL -framework OpenAL -framework Cocoa -framework IOKit -framework CoreVideo
endif
ifeq ($(UNAME_S), Linux)
	LIB_FLAGS += -lm -lpthread -ldl -lrt -lX11
endif
ifeq ($(findstring MINGW,$(UNAME_S)),MINGW)
	LIB_FLAGS += -lopengl32 -lgdi32
endif

.PHONY: run build clean setup

run:
	@echo " ------------------------------"
	@echo "|  Building the application    |"
	@echo " ------------------------------"
	$(COMPILER) $(CPP_FLAGS) $(OPTIMIZE_FLAGS) $(BUILD_FLAGS) $(LIB_FLAGS)

	@echo " ------------------------------"
	@echo "|  Running the application     |"
	@echo " ------------------------------"
	./build/codecrafty

build:
	@echo " ------------------------------"
	@echo "|  Building the application    |"
	@echo " ------------------------------"
	$(COMPILER) $(CPP_FLAGS) $(OPTIMIZE_FLAGS) $(BUILD_FLAGS) $(LIB_FLAGS)

setup:
	git clone https://github.com/raysan5/raylib.git libraries/raylib
	cd libraries/raylib/src && make PLATFORM=PLATFORM_DESKTOP RAYLIB_LIBTYPE=STATIC

setup_github:
	sudo apt-get update
	sudo apt-get install libx11-dev
	sudo apt-get install libxcursor-dev
	sudo apt-get install libxinerama-dev
	sudo apt-get install libxcursor-dev
	sudo apt-get install libgl1-mesa-dev
	sudo apt-get install xorg-dev
	
	git clone https://github.com/raysan5/raylib.git libraries/raylib
	cd libraries/raylib/src && make PLATFORM=PLATFORM_DESKTOP RAYLIB_LIBTYPE=STATIC

clean:
	@echo " ------------------------------"
	@echo "|  Cleaning the build folder   |"
	@echo " ------------------------------"
	rm -rf build/*
