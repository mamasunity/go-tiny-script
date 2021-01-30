package main

import (
	"github.com/signintech/gopdf"
)

// PdfColor has information about color and name.
type PdfColor struct {
	red   uint8
	green uint8
	blue  uint8
	name  string
}

// PdfColors is array of PdfColor
type PdfColors []PdfColor

func main() {
	var pdfColors PdfColors
	pdfColors = append(pdfColors, PdfColor{red: 174, green: 37, blue: 137, name: "neon_purple"})
	pdfColors = append(pdfColors, PdfColor{red: 225, green: 227, blue: 0, name: "yellow_submarine"})
	pdfColors = append(pdfColors, PdfColor{red: 47, green: 14, blue: 25, name: "darkness"})
	pdfColors = append(pdfColors, PdfColor{red: 0, green: 175, blue: 132, name: "illuminant_green"})
	pdfColors = append(pdfColors, PdfColor{red: 231, green: 66, blue: 143, name: "fluorescence_pink"})
	pdfColors = append(pdfColors, PdfColor{red: 91, green: 62, blue: 149, name: "universe"})
	pdfColors = append(pdfColors, PdfColor{red: 209, green: 86, blue: 36, name: "surprise_orange"})
	pdfColors = append(pdfColors, PdfColor{red: 167, green: 180, blue: 48, name: "electric_green"})
	pdfColors = append(pdfColors, PdfColor{red: 0, green: 126, blue: 170, name: "magical_turquoise"})

	for _, pdfColor := range pdfColors {
		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
		pdf.AddPage()

		pdf.SetFillColor(pdfColor.red, pdfColor.green, pdfColor.blue)
		pdf.Polygon([]gopdf.Point{{X: 0, Y: 0}, {X: 595, Y: 0}, {X: 595, Y: 842}, {X: 0, Y: 842}}, "DF")

		pdf.WritePdf(pdfColor.name + ".pdf")
	}
}
