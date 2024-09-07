package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf" // Import should be at the top, not inside the function
)

func main() {
	// Read serial numbers from input file
	serialNumbers, err := readSerialNumbers("serial_numbers.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Format serial numbers into a single line
	formattedText := formatSerialNumbers(serialNumbers)

	// Save the formatted text to a new file
	err = saveToFile("formatted_serial_numbers.txt", formattedText)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	// Convert the text file to a PDF
	err = convertTextToPDF("formatted_serial_numbers.txt", "formatted_serial_numbers.pdf")
	if err != nil {
		fmt.Println("Error converting to PDF:", err)
		return
	}

	fmt.Println("Process completed successfully.")
}

func readSerialNumbers(filename string) ([]string, error) {
	var serialNumbers []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		serialNumbers = append(serialNumbers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return serialNumbers, nil
}

func formatSerialNumbers(numbers []string) string {
	return strings.Join(numbers, ", ")
}

func saveToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func convertTextToPDF(inputFile, outputFile string) error {
	// Read the content of the text file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, string(content), "", "", false)

	// Save the PDF file
	return pdf.OutputFileAndClose(outputFile)
}

