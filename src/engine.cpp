#include "import.hpp"

class Button {
private:
    Rectangle bounds;
    Color normalColor;
    Color hoverColor;
    bool hovered;
    bool clicked;
    const char* text;

public:
    Button(Rectangle _bounds, Color _normalColor, Color _hoverColor, const char* _text)
        : bounds(_bounds), normalColor(_normalColor), hoverColor(_hoverColor),
        hovered(false), clicked(false), text(_text) {}

    bool is_pressed() {
        hovered = CheckCollisionPointRec(GetMousePosition(), bounds);

        if (hovered && IsMouseButtonReleased(MOUSE_LEFT_BUTTON)) {
            clicked = true;
        } else {
            clicked = false;
        }

        return clicked;
    }

    void draw() {
        DrawRectangleRec(bounds, hovered ? hoverColor : normalColor);
        DrawText(text, bounds.x + 10, bounds.y + 10, 20, RAYWHITE);
    }
};

class engine {
public:
    Rectangle new_rect(float x, float y, float width, float height) {
        Rectangle rect = {x, y, width, height};
        return rect;
    }

    void print(std::string str) {
        std::cout << str << std::endl;
    }

};