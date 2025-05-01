package main

import (
	"pdf-chromedp/chromedpHelper"
	"pdf-chromedp/wkhtmltopdfHelper"
	"strings"
)

func main() {
	inputFileNameList := []string{"engFile.html", "thaiFile.html"}

	for _, inputFileName := range inputFileNameList {
		chromedpHelper.CreatePDF("./input/"+inputFileName, "./output/chromedp-"+strings.ReplaceAll(inputFileName, ".html", ".pdf"))
		wkhtmltopdfHelper.CreatePDF("./input/"+inputFileName, "./output/wkhtmltopdf-"+strings.ReplaceAll(inputFileName, ".html", ".pdf"))
	}

}
