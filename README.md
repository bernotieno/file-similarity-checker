# File Similarity Checker

File Similarity Checker is a Go-based tool designed to compare files and determine their similarity. It uses token-based and structural comparison techniques to identify duplicate or similar files efficiently across various programming languages and formats.

---

## Features

- Compare code and text files for similarity using custom token-based algorithms.
- Detect similar or duplicate files in a directory.
- Supports multiple formats for output reports: `text`, `html`, and `pdf`.
- Lightweight and fast, built with Go.

---

##  Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/bernotieno/file-similarity-checker.git
    ```
2. Navigate to the project directory:
    ```bash
    cd file-similarity-checker
    ```
3. Build the project:
    ```bash
    go build
    ```

---

## 🚀 Usage

### Basic Syntax

```bash
./similaritychecker <format> [directory]
```

* `<format>` – Required. Specifies the format of the output report.

* `[directory]` – Optional. Path to the directory to analyze. Defaults to the current directory if omitted.

### Supported Output Formats

| Format    | Description                        | Output File             |
|-----------|------------------------------------|--------------------------|
| `text`    | Plain text report (default)        | `report_output.txt`      |
| `html`    | HTML table format                  | `report_output.html`     |
| `pdf`     | Clean, printable PDF report        | `report_output.pdf`      |


### Examples
1. PDF output in the current directory:

```bash
./similaritychecker pdf
```

2. Plain text output (default format):

```bash
./similaritychecker text ~/projects/
```
3. Html output
```bash
./similaritychecker html ~/projects/
```
 * `N.B`: Open the file in the browser to properly view for the html format

### 📄 Output
The generated report will include:

* File pairs compared

* Similarity percentage

* Category (e.g. "High", "Moderate", "Low")


