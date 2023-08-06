#include "import.hpp"
#include "engine.cpp"

class codecrafty {
    public:
        codecrafty() {
            sf::RenderWindow window;
            window.create(sf::VideoMode(width, height), title);
            app(&window);
            loop(&window);

        }
    private:
        engine engine;
        std::string title = "codecrafty";
        int width = 1280;
        int height = 720;

        void app(auto *window) {
            window->setFramerateLimit(60);
        }

        void loop(auto *window) {
            while (window->isOpen())
            {
                sf::Event event;
                while (window->pollEvent(event))
                {
                    if (event.type == sf::Event::Closed)
                        window->close();
                }

                engine.set_background_color(window, sf::Color::Black);

                engine.display(window);
            }
        }
};

int main() {
    codecrafty cc;
    return 0;
}

