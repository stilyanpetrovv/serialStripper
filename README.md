# SerialStripper

## Features

- Reads serial numbers from a text file (one serial number per line).
- Converts the list into a single line of comma-separated values.
- Saves the result into a new text file.
- Converts the text file into a PDF document.

## Usage

### 1. Install the Tool

You can install SerialStripper using `go install`:

```bash
go install github.com/stilyanpetrovv/serialStripper@latest
```

### Precompiled Binaries

For convenience, precompiled binaries are available for download. If you are using Windows, you can download the Windows executable directly from the [Releases](https://github.com/stilyanpetrovv/serialStripper/releases) page on GitHub.

- **Windows**: [serialStripper.exe](https://github.com/stilyanpetrovv/serialStripper/releases/latest/download/serialStripper.exe)

## Installing from Source

To build and install SerialStripper from source, follow these steps:

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher) must be installed on your system. Verify installation with `go version`.

### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/myacc/serialStripper.git
   cd serialStripper
   ```
   
3. **Build the project**

  ```bash   
  go build -o serialStripper
  ```

3. **Install the Binary**
  ```bash
  go install
  ```
  This installs the binary to your $GOPATH/bin directory.

## Packages

This project uses the following external packages:

- **[gofpdf](https://github.com/jung-kurt/gofpdf)**: A PDF generation library for Go. Used for converting the formatted serial numbers into a PDF document.


To install the required packages, run:

```bash
go mod tidy
```

## Running the Tool
#### You must name the file containing the serial numbers "serial_numbers.txt" otherwise it will throw an error and exit.
Considering you installed via `go install` you can use 
```bash
./serialStripper
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.
