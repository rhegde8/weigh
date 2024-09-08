# Weigh CLI

**Weigh** is a simple CLI tool written in Go that calculates the sizes of files and folders in a human-readable format for Unix-based systems (Linux/macOS). The size can be output in KB, MB, or GB, with GB as the default.

## Features

- Accepts multiple file or folder names as arguments.
- Outputs size in GB, MB, or KB.
- Handles nested directories.
- Verbose error messages for troubleshooting.

## Installation

Make sure you have Go installed on your machine. You can install Go by following the instructions on the official [Go website](https://golang.org/dl/).

Clone the repository:

```bash
git clone <repository_url>
cd weigh
go mod tidy

Usage
To use the tool, run the following command:

go run main.go --unit [KB|MB|GB] <file/folder_path>

Example:
go run main.go --unit MB ./myfolder ./myfile.txt

Flags
--unit [KB|MB|GB]: Specify the unit of size (default is GB).
--verbose: Enable verbose error messages.

License
This project is licensed under the MIT License. See the LICENSE file for more details.
