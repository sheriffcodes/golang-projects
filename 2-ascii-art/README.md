# 💻 ASCII Art Generator

ASCII Art Generator is a command-line tool written in Go that transforms standard text strings into ASCII art banners. The program parses a template file and reconstructs characters horizontally to form large-scale graphic representations.

---

## ## Features

* **Custom Font Support:** Loads character templates from a `standard.txt` banner file.
* **Horizontal Rendering:** Correctly aligns multi-line ASCII characters side-by-side.
* **Newline Handling:** Supports literal `\n` inputs to create multi-block ASCII art.
* **Error Handling:** Validates command-line arguments and character availability.
* **Performance Optimized:** Uses `strings.Builder` for efficient string concatenation.

---

## ## Getting Started

### ### Prerequisites
* Go (version 1.18 or higher recommended)
* A `standard.txt` file in the root directory.

### ### Installation
1.  Clone the repository:
    ```bash
    git clone https://github.com/sheriffcodes/golang-projects/
    cd golang-projects/ascii-art
    ```
2.  Ensure your `standard.txt` file is present.

### ### Usage
Run the program by passing a string as an argument (in the terminal):

```bash
go run . "Hello"
```

To include newlines in your art:
```bash
go run . "Hello\nWorld"
```

---

## ## Testing

The project includes table-driven tests to verify rendering accuracy, horizontal alignment, and newline processing.

### ### Running Tests
To run the automated tests, execute:

```bash
go test -v
```

### ### Test Logic
The tests use a **Mock Banner Map** to isolate the rendering logic from the physical `standard.txt` file. This ensures that the code's logic is verified independently of the data source. Key test cases include:
* **Single Character:** Validates 1:1 mapping of runes.
* **Horizontal Concatenation:** Ensures letters are glued together row-by-row.
* **Newline Processing:** Confirms that `\n` triggers a vertical break in the output.

---

## ## File Structure

* `main.go`: Contains the core logic and engine functions.
* `main_test.go`: Contains the unit tests and mock data.
* `standard.txt`: The font template file (required for execution).