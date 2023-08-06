COMPILER = g++
LIB_FLAGS = -lsfml-graphics -lsfml-window -lsfml-system
BUILD_FLAGS = -o build/codecrafty src/main.cpp
OPTIMIZE_FLAGS = -O2
CPP_FLAGS = -std=c++20

.PHONY: run build clean

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

clean:
	@echo " ------------------------------"
	@echo "|  Cleaning the build folder   |"
	@echo " ------------------------------"
	rm -rf build/*

