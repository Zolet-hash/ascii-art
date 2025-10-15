# ASCII Art in Go 🖋️

## 📘 Project Overview

**ascii-art** is a Go program that takes a string as an argument and outputs it as a **graphic representation using ASCII characters** — effectively *writing text in big ASCII letters*.  

The project supports rendering using different banner templates (`standard`, `shadow`, `thinkertoy`), handles spaces, numbers, special characters, and even multi-line inputs using the literal `\n`.

This project is inspired by the **FIGlet** concept but implemented from scratch using only the **Go standard library**.

---

## 🎯 Objectives

- Learn to manipulate data and files in Go.
- Understand how to use the Go **filesystem (fs)** API.
- Gain experience in **string handling**, **command-line arguments**, and **error checking**.
- Practice **clean code** and **unit testing** principles.

---

## 🏗️ Project Structure

ascii-art/
├── main.go
├── standard
├── shadow
├── thinkertoy
└── README.md


- **main.go** → Main Go source file for the ASCII renderer.  
- **standard**, **shadow**, **thinkertoy** → Banner files that define how each character should look in ASCII form.  
- **README.md** → Project documentation (this file).  

---

## ⚙️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/ascii-art.git
   cd ascii-art

    Make sure Go is installed

    go version

    You should have Go 1.20+.

    Place the banner files
    Ensure the banner files standard, shadow, or thinkertoy are in the same directory as main.go.

🚀 Usage
Basic usage

go run . "Hello"

Using braces (matches example format)

go run . "{Hello There}"

Using special characters and spaces

go run . "1Hello 2There"

Multi-line input (using \n)

go run . "Hello\nThere"

Selecting a banner

go run . "Hello" shadow
go run . "Hello" thinkertoy

    🔹 If no banner is provided, the default banner standard will be used.

🧠 Example Outputs
Example 1

go run . "{Hello There}" | cat -e

Output:
```bash
   __  _    _          _   _                 _______   _                           __    $
  / / | |  | |        | | | |               |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                              /_/   $
                                                                                         $
```
Example 2 (multi-line)

go run . "Hello\nThere" | cat -e

Output:
```console
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```
🧩 Banner File Format

Each banner file contains 95 printable ASCII characters (from space " " to "~")
Each character representation has:

    A height of 8 lines

    A blank line separating each character

Example of the first few characters in a banner file:

```console
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......
```

🧪 Testing

You are encouraged to create test files (*_test.go) for unit testing.
Here’s what you can test:
Function	Test Focus
loadBanner()	Ensure correct parsing of the banner file (95 glyphs).
sanitizeInput()	Handles {}, quotes, and \n correctly.
renderLine()	Builds ASCII lines as expected for simple text.

Example test command:

go test -v

🧱 Code Highlights

    Uses only standard library packages (fmt, os, strings, bufio, errors, path/filepath).

    No external dependencies.

    Handles invalid characters by rendering them as spaces.

    Gracefully fails with clear error messages if banner file is missing or malformed.

    Clean separation of logic into helper functions:

        loadBanner()

        sanitizeInput()

        glyphForRune()

        renderLine()

💡 Learning Outcomes

By completing this project, you’ll understand how to:

    Manipulate files and strings in Go.

    Parse and map ASCII glyph data.

    Work with the Go command-line interface (CLI).

    Build modular, testable Go programs.

    Follow good code formatting and documentation practices.

🧰 Troubleshooting
Problem	Solution
failed to load banner file	Make sure standard, shadow, or thinkertoy file is in the same directory.
Output looks broken	Check for extra spaces or tabs in your banner file. Each line must be consistent width.
Input not showing	Try quoting it: go run . "{Hello There}"
\n not working	Ensure you used double quotes: "Hello\nThere" not 'Hello\nThere'.
🧑‍💻 Author

Sir Emmanuel Otieno
Founder of Kenyan Brain Foundation
Ambassador of AI in Africa 🌍
📍 Kisumu, Kenya
📧 [Contact: +2547 5971 9674]
📜 License

This project is open-source and distributed for educational purposes.
You may freely use, modify, and share it with proper attribution.

✨ “Time to write big — with ASCII!” ✨


---

Would you like me to include **unit test templates (`ascii_art_test.go`)** for this program next — covering banner loading and rendering?
