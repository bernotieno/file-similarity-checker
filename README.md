# File Similarity Checker

File Similarity Checker is a Go-based tool designed to compare files and determine their similarity. It uses hashing and other comparison techniques to identify duplicate or similar files efficiently.

## Features

- Compare files for similarity using hashing algorithms.
- Identify duplicate files in a directory.
- Lightweight and fast, built with Go.

## Installation

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
    go build -o checker .
    ```

## Usage

Run the tool with the following command:
```bash
./checker  <directory>
```
* Replace <directory> with the path of the directory you want to analyze. If no directory is provided, the tool will default to the current working directory.

* The tool will generate a similarity report in the file_similarity_report.txt file.