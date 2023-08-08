#include "engine/import.hpp"
#include "menu.cpp"

class codecrafty {
public:
    codecrafty() {
        init();

    }
private:
    char title[11] = "codecrafty";
    int width = 1280;
    int height = 720;

    void init() {
        InitWindow(width, height, title);
        SetTargetFPS(60);

        menu menu;
    }
};

int main() {
    codecrafty app;
    return 0;
}

