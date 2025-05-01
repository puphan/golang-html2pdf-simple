package wkhtmltopdfHelper

import (
	"fmt"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func CreatePDF(inputHtmlPath string, outputPdfPath string) {
	// create PDF new generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("error creating PDF generator:", err)
		return
	}

	// open HTML file in local system
	fileReader, err := os.Open(inputHtmlPath)
	if err != nil {
		fmt.Println("error opening HTML file:", err)
		return
	}
	defer fileReader.Close()

	// create new page from HTML file
	page := wkhtmltopdf.NewPageReader(fileReader)
	pdfg.AddPage(page)

	// setting PDF options
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(80)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// convert HTML to PDF
	err = pdfg.Create()
	if err != nil {
		fmt.Println("error generating PDF:", err)
		return
	}

	// write PDF to file
	err = pdfg.WriteFile(outputPdfPath)
	if err != nil {
		fmt.Println("error writing PDF file:", err)
		return
	}

	fmt.Println("wkhtmltopdf created PDF successfully")
}
