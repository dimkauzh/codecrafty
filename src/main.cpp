#include "import.hpp"
#include "engine.cpp"
#include "fonts.cpp"

class codecrafty {
    public:
        codecrafty() {
            init();

        }
    private:
        engine engine;
        fonts all_fonts;

        char title[11] = "codecrafty";
        int width = 1280;
        int height = 720;

        void init() {
            InitWindow(width, height, title);
            SetTargetFPS(60);

            Menu();
        }

        void Menu() {

            Button new_project(engine.new_rect(50, 300, 150, 45), GRAY, LIGHTGRAY, "New Project");
            Button open_project(engine.new_rect(50, 350, 150, 45), GRAY, LIGHTGRAY, "Open Project");

            while (!WindowShouldClose())
            {

                BeginDrawing();

                    ClearBackground(RAYWHITE);
                    DrawText("Welcome to codecrafty!", 400, 10, 30, BLACK);
                    new_project.draw();
                    open_project.draw();


                EndDrawing();

                if (new_project.is_pressed()) {
                    engine.print("New Project");
                }
                else if (open_project.is_pressed()) {
                    engine.print("Open Project");
                }

            }
        }
};

int main() {
    codecrafty app;
    return 0;
}

