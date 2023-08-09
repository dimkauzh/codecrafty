# CodeCrafty
🎉📝 Welcome to codecrafty! 🎉📝

### Description
🚀 codecrafty is a super fast and extensible code editor written in Python and powered by Toga + Briefcase. It's designed to be an enjoyable and educational project, where you can dive into the world of code and unleash your creativity! 🤩

### Disclaimer
⚠️ Important: codecrafty is purely created for fun and educational purposes. 🤓 It's packed with weird design issues that will keep you on your toes! Embrace the quirks and let your coding adventures begin! 🕵️‍♂️💻

## Coming Features
🔥 Experience the thrill of coding in a unique environment, where the unexpected becomes the norm! 🌀🎢

🧩 Super Extensible: codecrafty provides a framework that allows you to add your own twists, extensions, and surprises! 🧠💡

🚀 Blazing Fast: Harness the power of Python and Toga + Briefcase for a seamless coding experience. ⚡️💻

🌈 Aesthetically Quirky: Embrace the weirdness of codecrafty's design issues for a one-of-a-kind coding journey! 🎨😄

## Getting Started
As of now, codecrafty is only setup to run on unix, but you still can [run it on windows](#Windows) if you're feeling adventurous! 🤠

### Unix
Because were are using briefcase and stuff, we included a command in our Makefile to setup everything you need, this is all it does for you:
 - Create a virtual environment
 - Install Briefcase
 - Install Toga

To run the command, clone the repo:
```bash
git clone https://github.com/dimkauzh/codecrafty.git
cd codecrafty
```

Then run it:
```bash
make setup
```

### Windows
If you're on windows, you can still run codecrafty, but you'll have to do a bit more work. 🤓

First, clone the repo:
```bash
git clone https://github.com/dimkauzh/codecrafty.git
cd codecrafty
```

Then, create a virtual environment and activate it:
```bash
python -m venv venv
venv\Scripts\activate
```

Next, install the dependencies:
```bash
pip install -r requirements.txt
```

And your done! Now head over to [run](#Run) to start coding! 🚀

### Run
To run codecrafty, simply run the following command:
```bash
make run
```

This will run the dev release of our code editor (with briefcase).
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
