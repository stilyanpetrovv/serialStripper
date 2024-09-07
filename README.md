# SerialStripper

SerialStripper is a simple CLI tool designed to process serial numbers from a text file, format them into a comma-separated list, and save the result in a new text file. It will also convert the formatted list into a PDF file.

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

## Running the Tool
#### You must name the file containing the serial numbers "serial_numbers.txt" otherwise it will throw an error and exit.
Considering you installed via `go install` you can use 
```bash
./serialStripper
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
