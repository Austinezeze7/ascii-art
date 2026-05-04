# ASCII Art

## 📌 Overview

**Ascii-art** is a Go program that takes a string as input and outputs a graphical representation of that string using ASCII characters.

The program reads from predefined banner files and prints characters in a stylized multi-line format (each character has a height of 8 lines).

It supports:

* Letters (uppercase & lowercase)
* Numbers
* Spaces
* Special characters
* Line breaks (`\n`)

---

## 🚀 Features

* Convert text into ASCII art
* Support for multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Handles multi-line input (`\n`)
* Clean and modular code structure
* Unit tests included

---

## 🗂️ Project Structure

```id="s9k2jd"
.
├── main.go
├── main_test.go
├── go.mod
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
└── src/
    ├── printer.go
    ├── reader.go
    ├── utils.go
    ├── logic_test.go
    ├── reader_test.go
```

---

## 🧩 File Responsibilities

### `main.go`

* Entry point of the application
* Handles command-line arguments
* Calls functions from `src/` to process and print ASCII art

### `main_test.go`

* Tests the overall behavior of the program
* Validates integration between components

### `src/reader.go`

* Reads banner files from the `banners/` directory
* Parses ASCII representations into usable data structures (e.g., maps)

### `src/printer.go`

* Responsible for rendering the ASCII art to the terminal
* Combines character lines correctly across multiple rows

### `src/utils.go`

* Contains helper functions
* Handles string manipulation, validation, and formatting logic

### `src/logic_test.go`

* Unit tests for core transformation logic
* Ensures characters are correctly mapped and processed

### `src/reader_test.go`

* Tests banner file reading and parsing
* Ensures correct handling of file input and structure

### `banners/*.txt`

* Contain ASCII templates for characters
* Each character is represented in 8 lines
* Used as the source for rendering output

---

## 🧠 Why the `src/` Folder?

The `src/` folder is used to separate the **core logic** of the application from the entry point (`main.go`).

### Benefits:

* **Separation of concerns**:
  `main.go` handles input/output, while `src/` handles processing logic.

* **Better maintainability**:
  Code is easier to navigate and modify.

* **Improved testing**:
  Logic inside `src/` can be tested independently without running the whole program.

* **Reusability**:
  Functions in `src/` can be reused or extended without affecting the main program.

---

## ⚙️ How It Works

1. The program reads the input string from command-line arguments.
2. It loads the selected banner file.
3. Each character is mapped to its ASCII representation.
4. The result is printed line by line to form the final output.

---

## 🧪 Running the Program

### Basic usage:

```bash id="h8z1mf"
go run . "Hello"
```

### With new lines:

```bash id="4vq9iy"
go run . "Hello\nWorld"
```

### Using output visualization:

```bash id="e5j3nt"
go run . "Hello" | cat -e
```

---

## 🧪 Running Tests

```bash id="u7k9cw"
go test ./...
```

---

## 📄 Banner Format

* Each character is represented using **8 lines**
* Characters are separated by a newline
* ASCII values range from **32 (space) to 126 (~)**

---

## 📦 Allowed Packages

This project strictly uses only Go's standard library. The following packages are used:

* `fmt` → for formatted I/O operations
* `os` → for handling command-line arguments and file operations
* `bufio` → for efficient file reading
* `strings` → for string manipulation
* `testing` → for writing and running unit tests

---

## 🎯 Learning Objectives

This project helps you understand:

* File handling in Go (`fs` API)
* String manipulation
* Structuring Go applications
* Writing unit tests
* Working with ASCII encoding

---

## 👨‍💻 Author

* amuyonga

---

## 🔗 Repository

https://learn.zone01kisumu.ke/git/amuyonga/ascii-art

---
