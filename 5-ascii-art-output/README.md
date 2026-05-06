# 💻 ASCII Art Output Generator

## Description
`ascii-art-output` is a program written in Go that converts text strings via command line into stylized ASCII art. This project extends the basic ASCII art generator by adding a dedicated output redirection feature, allowing users to save their art directly to a text file using a specialized command-line flag.

The program supports multiple banner styles and handles complex input including newline characters, ensuring the output maintains its structural integrity whether viewed in a terminal or a text editor.

## Features
- **File Redirection:** Save ASCII art directly to a `.txt` file using the `--output=<fileName.txt>` flag.
- **Multiple Banner Support:** Supports `standard`, `shadow`, and `thinkertoy` banner formats.
- **Newline Processing:** Correctly interprets and renders both literal and escaped newline characters (`\\n`).
- **Robust Argument Parsing:** Flexibly handles different argument counts and positions.
- **Cross-Platform Compatibility:** Normalizes line endings (`\\r\\n` to `\\n`) to ensure banner files load correctly on all operating systems.

## Project Structure
```text
.
├── banners/          # Directory containing .txt banner files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── main.go           # Primary source code
├── main_test.go      # Source code test file
├── README.md         # Project documentation
└── [generated_file].txt # Created when using the --output flag
```

## Installation
1. Ensure you have [Go](https://go.dev/doc/install) installed (version 1.16 or higher recommended).
2. Clone this repository or copy the source code to your local machine.
3. Ensure your banner files are located in a subdirectory named `banners/`.

## Usage
The program follows a specific syntax for its arguments:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Options
- `--output=<fileName.txt>`: Specifies the name of the file where the ASCII art will be saved.

### Examples

**1. Output to Terminal (Default)**
```bash
go run . "Hello" standard
```

**2. Save Output to a File**
```bash
go run . --output=result.txt "Hello World" shadow
```

**3. Multi-line Input to a File**
```bash
go run . --output=test.txt "Graphic\\nDesign" thinkertoy
```

## Technical Implementation
### The Storage Engine (`loadBanner`)
The program reads banner files from the `banners/` directory. Each character in a banner file is represented by 8 lines of ASCII art, separated by an empty line. The engine maps these blocks to their corresponding characters (starting from ASCII Space - 32).

### The Processing Engine (`Render`)
ASCII art is rendered horizontally. The engine iterates through the 8 rows of each character in a string segment, building the final output string line-by-line using a `strings.Builder` for optimal performance.

### Argument Management
The `main` function performs "argument filtering." It scans for the `--output` prefix to extract the target filename and separates it from the actual text and banner style inputs.

## Verification
To ensure the output file contains no trailing spaces and that newlines are correctly placed, you can verify the file using the following command:

```bash
cat -e [fileName].txt
```
*The `$` character at the end of each line in the `cat -e` output indicates the line ending, helping you confirm that the formatting matches the requirements precisely.*
