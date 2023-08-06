#include "import.hpp"

class engine {
    public:
        void set_background_color(auto window, sf::Color color) {
            window->clear(color);
        }
        void display(auto window) {
            window->display();
        }

};