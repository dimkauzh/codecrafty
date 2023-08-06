#include "import.hpp"

class codecrafty {
    public:
        codecrafty() {
            sf::Window window(sf::VideoMode(800, 600), "My window");
            loop(&window);

        }
    private:
        void loop(sf::Window *window) {
            while (window->isOpen())
            {
                sf::Event event;
                while (window->pollEvent(event))
                {
                    if (event.type == sf::Event::Closed)
                        window->close();
                }
            }
        }
};

int main() {
    codecrafty cc;
    return 0;
}

