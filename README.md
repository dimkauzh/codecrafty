<h1 align="center">🚀 CodeCrafty</h1>
<p align="center">
🎉📝 Welcome to codecrafty! 🎉📝
</p>


### Description
🚀 codecrafty is a super fast and extensible code editor written in Go and powered by FLTK. It's designed to be an enjoyable and educational project, where you can dive into the world of code and unleash your creativity! 🤩

### Disclaimer
⚠️ Important: codecrafty is purely created for fun and educational purposes. 🤓 It's packed with weird design issues that will keep you on your toes! Embrace the quirks and let your coding adventures begin! 🕵️‍♂️💻

## Coming Features
🔥 Experience the thrill of coding in a unique environment, where the unexpected becomes the norm! 🌀🎢

🧩 Super Extensible: codecrafty provides a framework that allows you to add your own twists, extensions, and surprises! 🧠💡

🚀 Blazing Fast: Harness the power of Go and FLTK for a seamless coding experience. ⚡️💻

🌈 Aesthetically Quirky: Embrace the weirdness of codecrafty's design issues for a one-of-a-kind coding journey! 🎨😄

## Getting Started
As of now, codecrafty is only setup to run on unix, but you still can run it on windows if you're feeling adventurous! 🤠

### Setting up
You need go for codecrafty. In addition to Go, you will also need a C++11 compiler. The FLTK libraries are bundled with the repo for x86_64 Linux, MacOS and Windows (mingw64).
You also need some system libs which are normally available on operating systems with a graphical user interfaces:
- Windows: No external dependencies (for mingw64)
- MacOS: No external dependencies
- Linux: You need:
    - x11
    - Xrender
    - Xcursor
    - Xfixes
    - Xext
    - Xft
    - Xinerama
    - OpenGL

Clone the repo:
```bash
git clone https://github.com/dimkauzh/codecrafty.git
cd codecrafty
```

Then run the setup command:
```bash
make setup
```

### Build/Run
To follow these instructions you need to have the repo cloned

To run codecrafty, simply run the following command:
```bash
make run
```

And to build codecrafty, just run:
```bash
make build
```

Let the fun begin! 🎉🚀

## Contributing
🙌 We welcome contributions to codecrafty! Whether it's fixing a bug or adding new quirky features, your creativity is valued! 🤝🎭

Please adhere to the following guidelines:

Create a new branch for your contribution:

```bash
git checkout -b your-feature
```

Make your changes and commit:

```bash
git commit -m "Add some feature"
```

Push to the branch:

```bash
git push origin your-feature
```

Open a pull request and let's discuss the awesomeness you've brought to codecrafty! 🚀📩

## Support
🤝 We appreciate your interest in codecrafty! For any questions or feedback, reach out to us using github issues! 💌

## License
📜 codecrafty is licensed under the MIT License - see the LICENSE file for details. 📄

Happy coding! 😄🚀🎉
