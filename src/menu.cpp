#include "engine/import.hpp"
#include "engine/engine.hpp"

class menu {
public:
    menu() {
        loading_screen();
    }
private:
    engine engine;

    void loading_screen() {
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