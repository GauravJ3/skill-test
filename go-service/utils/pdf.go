package utils

import (
	"bytes"
	"fmt"

	"github.com/GauravJ3/go-service/models"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(student models.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Student Report - %s", student.FullName))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("ID: %d", student.ID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.FullName))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Class: %s", student.Class))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Roll No: %s", student.RollNo))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	return buf.Bytes(), err
}
