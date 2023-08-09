"""
A super fast and extensible code editor written in Python
"""

import tkinter as tk
import codecrafty.menu as menu

class CodeCrafty:
    def __init__(self) -> None:
        self.run()

    def run(self):
        self.root = tk.Tk()
        menu.Menu(self.root)
        self.root.mainloop()
