package main

import (
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.SetFillColor(230, 180, 34)
	pdf.Polygon([]gopdf.Point{{X: 0, Y: 0}, {X: 595, Y: 0}, {X: 595, Y: 842}, {X: 0, Y: 842}}, "DF")

	pdf.WritePdf("kogane.pdf")

}
