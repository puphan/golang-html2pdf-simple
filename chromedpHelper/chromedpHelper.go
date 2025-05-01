package chromedpHelper

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func CreatePDF(inputHtmlPath string, outputPdfPath string) {

	pdfBytes, err := PrintToPDF(inputHtmlPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(outputPdfPath, pdfBytes, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("chromedp created PDF successfully")

}

func PrintToPDF(htmlPath string) ([]byte, error) {
	log.Println(LocalFileURI(htmlPath))
	// create context for chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pdfBuf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(LocalFileURI(htmlPath)),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return nil, err
	}

	return pdfBuf, nil
}

func LocalFileURI(path string) string {
	absPath, _ := filepath.Abs(path)
	return "file://" + absPath
}
