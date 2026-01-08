package main

import (
	"fmt"
	"os"
	"github.com/jung-kurt/gofpdf"
)

// pdfSaver structure
type PdfSaver struct {
	OutputDir string  
}

// Constructor pdf
func NewPdf(outputDir string) (*PdfSaver, error) {
	if outputDir == "" {
		return nil,fmt.Errorf("Repertoire de sortie non precisee !")

	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return nil,fmt.Errorf("Le repertoire %s est inexistant !", outputDir)
	}
	// pdf creation
	return &PdfSaver{OutputDir: outputDir}, nil
}

// Function to save ours pdf
func (p *PdfSaver) Save(pir Pirate) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("WANTED", false)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 36)
	pdf.WriteAligned(0, 30, "WANTED", "C")
	pdf.Ln(20)

	pdf.SetFont("Arial", "B", 24)
	pdf.WriteAligned(0, 20, pir.Name, "C")
	pdf.Ln(10)

	if _, err := os.Stat(pir.Img); err == nil {
		pdf.ImageOptions(
			pir.Img,
			55, 90,
			100, 0,
			false,
			gofpdf.ImageOptions{ReadDpi: true},
			0,
			"",
		)
	}

	pdf.Ln(110)
	pdf.SetFont("Arial", "B", 20)
	pdf.WriteAligned(0, 20, "PRIME : "+pir.Prime+" BERRYS", "C")

	filename := fmt.Sprintf("%s/%s_WANTED.pdf", p.OutputDir, pir.Name)
	fmt.Println("PDF généré :", filename)
	return pdf.OutputFileAndClose(filename)
}