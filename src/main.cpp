#include "import.hpp"
#include "engine.cpp"
#include "fonts.cpp"

class codecrafty {
    public:
        codecrafty() {
            init();
            app();

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
        }

        void app() {

            Button button(engine.new_rect(300, 300, 200, 200), YELLOW, YELLOW, "HELLO");

            while (!WindowShouldClose())
            {

                BeginDrawing();

                    ClearBackground(RAYWHITE);
                    DrawText("Congrats! You created your first window!", 190, 200, 20, LIGHTGRAY);
                    button.draw();


                EndDrawing();

                if (button.is_pressed()) {
                    engine.print("clicked");
                }

            }
        }
};

int main() {
    codecrafty cc;
    return 0;
}

